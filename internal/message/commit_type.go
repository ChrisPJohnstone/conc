package message

import (
	"fmt"
	"log/slog"
	"strings"
)

func GetCommitTypes() []string {
	return []string{
		commitTypeBuild,
		commitTypeChore,
		commitTypeCi,
		commitTypeDocs,
		commitTypeFeat,
		commitTypeFix,
		commitTypePerf,
		commitTypeRefactor,
		commitTypeRevert,
		commitTypeStyle,
		commitTypeTest,
	}
}

func (msg *Message) commitTypeFromInput() error {
	slog.Debug(fmt.Sprintf("Type flag specified: '%s'", msg.CommitType))
	commitTypes := GetCommitTypes()
	var matches []string
	for _, commitType := range commitTypes {
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

func (msg *Message) commitTypeFromBranch() error {
	// TODO: Support prefix / fuzzy
	slog.Debug(fmt.Sprintf("Checking branch name '%s' for type", msg.branch))
	commitTypes := GetCommitTypes()
	for _, commitType := range commitTypes {
		if strings.HasPrefix(msg.branch, commitType) {
			msg.CommitType = commitType
			return nil
		}
	}
	slog.Debug("No type matched. Using fallback")
	msg.CommitType = CommitTypeFallback
	return nil
}

func (msg *Message) parseCommitType() error {
	slog.Debug("Attempting to parse commit type")
	if msg.CommitType != "" {
		return msg.commitTypeFromInput()
	}
	if !msg.hasBranch {
		msg.getGitBranch()
	}
	if msg.branchErr == nil {
		return msg.commitTypeFromBranch()
	} else {
		slog.Debug("Error getting git branch", "error", msg.branchErr)
	}
	slog.Debug("Failed to parse commit type. Using fallback")
	msg.CommitType = CommitTypeFallback
	return nil
}
