package message

import (
	"fmt"
	"log/slog"
	"regexp"

	"github.com/ChrisPJohnstone/conc/internal/git"
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
	if msg.Scope != "" && msg.Scope != ScopeFlagSentinel {
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
