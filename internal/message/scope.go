package message

import (
	"fmt"
	"log/slog"
	"regexp"
)

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
