package message

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/ChrisPJohnstone/git-conc/v2/internal/git"
)

type Message struct {
	Body       string
	Breaking   bool
	CommitType string
	Footer     string
	Scope      string
	Subject    string
	branch     string
	branchErr  error
	hasBranch  bool
}

func (msg *Message) getGitBranch() {
	slog.Debug("Attempting to get git branch")
	msg.branch, msg.branchErr = git.GitBranch()
	msg.hasBranch = true
	if msg.branchErr != nil {
		slog.Debug("Error getting git branch", "error", msg.branchErr)
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
	if msg.Breaking || strings.HasPrefix(msg.Footer, "BREAKING CHANGE:") {
		message = fmt.Sprintf("%s!", message)
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
