package message

import (
	"fmt"
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
	logger.Printf("Attempting to get git branch\n")
	msg.branch, msg.branchErr = git.GitBranch()
	msg.hasBranch = true
	if msg.branchErr != nil {
		logger.Printf("Error getting git branch: '%s'", msg.branchErr)
	}
}

func (msg *Message) parseCommitType() error {
	logger.Printf("Attempting to parse commit type: '%s'\n", msg.CommitType)
	if msg.CommitType != "" {
		logger.Printf("Type flag specified: '%s'\n", msg.CommitType)
		var matches []string
		for _, commitType := range globals.CommitTypes {
			if strings.HasPrefix(commitType, msg.CommitType) {
				logger.Printf("Flag matches '%s'\n", commitType)
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
		logger.Printf(
			"Error getting branch: '%s'. Using fallback\n",
			msg.branchErr,
		)
		msg.CommitType = globals.CommitTypeFallback
		return nil
	}
	// TODO: Support prefix / fuzzy
	logger.Printf("Checking branch name '%s' for type\n", msg.branch)
	for _, commitType := range globals.CommitTypes {
		if strings.HasPrefix(msg.branch, commitType) {
			msg.CommitType = commitType
			return nil
		}
	}
	logger.Printf("No type matched. Ussing fallback\n")
	msg.CommitType = globals.CommitTypeFallback
	return nil
}

func (msg *Message) parseScope() {
	if msg.Scope != "" || msg.Scope == globals.ScopeFlagSentinel {
		logger.Printf("Scope provided: '%s'. Parse not required\n", msg.Scope)
		return
	}
	logger.Printf("Attempting to parse scope: '%s'\n", msg.Scope)
	if !msg.hasBranch {
		msg.getGitBranch()
	}
	if msg.branchErr != nil {
		logger.Printf(
			"Error getting branch: '%s'. Falling back to '%s'\n",
			msg.branchErr,
			scopeFallback,
		)
		msg.Scope = scopeFallback
		return
	}
	pattern := regexp.MustCompile(scopeBranchPattern)
	matches := pattern.FindStringSubmatch(msg.branch)
	if len(matches) == 2 {
		msg.Scope = matches[1]
		logger.Printf(
			"Branch name matches pattern, found scope: '%s'\n",
			msg.Scope,
		)
		return
	}
}

func (msg *Message) Parse() error {
	err := msg.parseCommitType()
	if err != nil {
		return err
	}
	logger.Printf("Parsed commit type: '%s'\n", msg.CommitType)
	msg.parseScope()
	logger.Printf("Parsed scope: '%s'\n", msg.Scope)
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
