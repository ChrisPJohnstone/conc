package message

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/ChrisPJohnstone/conc/internal/git"
	"github.com/ChrisPJohnstone/conc/internal/globals"
)

type Message struct {
	Subject    string
	CommitType string
	Scope      string
	Body       string
	Footer     string
	hasBranch  bool
	branch     string
	branchErr  error
}

func (msg *Message) getGitBranch() {
	slog.Debug("Attempting to get git branch")
	msg.branch, msg.branchErr = git.GitBranch()
	msg.hasBranch = true
	if msg.branchErr != nil {
		slog.Debug("Error getting git branch", "error", msg.branchErr)
	}
}

func (msg *Message) parseCommitType() error {
	slog.Debug(fmt.Sprintf(
		"Attempting to parse commit type: '%s'",
		msg.CommitType,
	))
	if msg.CommitType != "" {
		slog.Debug(fmt.Sprintf("Type flag specified: '%s'", msg.CommitType))
		var matches []string
		for _, commitType := range globals.CommitTypes {
			if strings.HasPrefix(commitType, msg.CommitType) {
				slog.Debug(fmt.Sprintf("Flag matches '%s'", commitType))
				matches = append(matches, commitType)
			}
		}
		switch len(matches) {
		case 0:
			return fmt.Errorf(
				"'%s' is not a valid commit type",
				msg.CommitType,
			)
		case 1:
			msg.CommitType = matches[0]
			return nil
		default:
			return fmt.Errorf(
				"'%s' is ambiguous: %s. Use a longer prefix",
				msg.CommitType,
				strings.Join(matches, ", "),
			)
		}
	}
	if !msg.hasBranch {
		msg.getGitBranch()
	}
	if msg.branchErr != nil {
		slog.Debug(
			"Error getting branch. Using fallback",
			"error",
			msg.branchErr,
		)
		msg.CommitType = globals.CommitTypeFallback
		return nil
	}
	// TODO: Support prefix / fuzzy
	slog.Debug(fmt.Sprintf("Checking branch name '%s' for type", msg.branch))
	for _, commitType := range globals.CommitTypes {
		if strings.HasPrefix(msg.branch, commitType) {
			msg.CommitType = commitType
			return nil
		}
	}
	slog.Debug("No type matched. Using fallback")
	msg.CommitType = globals.CommitTypeFallback
	return nil
}

func (msg *Message) parseScope() {
	if msg.Scope != "" {
		slog.Debug(fmt.Sprintf(
			"Scope provided: '%s'. Parse not required",
			msg.Scope,
		))
		return
	}
	slog.Debug(fmt.Sprintf("Attempting to parse scope: '%s'", msg.Scope))
	if !msg.hasBranch {
		msg.getGitBranch()
	}
	if msg.branchErr != nil {
		slog.Debug(
			"Error getting branch. Using fallback",
			"error",
			msg.branchErr,
		)
		msg.Scope = scopeFallback
		return
	}
	pattern := regexp.MustCompile(scopeBranchPattern)
	matches := pattern.FindStringSubmatch(msg.branch)
	if len(matches) == 2 {
		msg.Scope = matches[1]
		slog.Debug(fmt.Sprintf(
			"Branch name matches pattern, found scope: '%s'",
			msg.Scope,
		))
		return
	}
}

func (msg *Message) Parse() error {
	err := msg.parseCommitType()
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Parsed commit type: '%s'", msg.CommitType))
	msg.parseScope()
	slog.Debug(fmt.Sprintf("Parsed scope: '%s'", msg.Scope))
	return nil
}

func (msg *Message) String() string {
	message := msg.CommitType
	if msg.Scope != "" && msg.Scope != globals.ScopeFlagSentinel {
		message = fmt.Sprintf("%s(%s)", message, msg.Scope)
	}
	message = fmt.Sprintf("%s: %s", message, msg.Subject)
	if msg.Body != "" {
		message = fmt.Sprintf("%s\n\n%s", message, msg.Body)
	}
	if msg.Footer != "" {
		message = fmt.Sprintf("%s\n\n%s", message, msg.Footer)
	}
	return message
}
