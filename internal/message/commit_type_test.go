package message

import (
	"errors"
	"testing"
)

func TestCommitTypeFromInput(t *testing.T) {
	tests := []struct {
		input     string
		expected  string
		expectErr bool
	}{
		// Simple Cases
		{input: "build", expected: commitTypeBuild},
		{input: "chore", expected: commitTypeChore},
		{input: "ci", expected: commitTypeCi},
		{input: "docs", expected: commitTypeDocs},
		{input: "feat", expected: commitTypeFeat},
		{input: "fix", expected: commitTypeFix},
		{input: "perf", expected: commitTypePerf},
		{input: "refactor", expected: commitTypeRefactor},
		{input: "revert", expected: commitTypeRevert},
		{input: "style", expected: commitTypeStyle},
		{input: "test", expected: commitTypeTest},

		// Single letter
		{input: "a", expectErr: true},
		{input: "b", expected: commitTypeBuild},
		{input: "c", expectErr: true},
		{input: "d", expected: commitTypeDocs},
		{input: "e", expectErr: true},
		{input: "f", expectErr: true},
		{input: "g", expectErr: true},
		{input: "h", expectErr: true},
		{input: "i", expectErr: true},
		{input: "j", expectErr: true},
		{input: "k", expectErr: true},
		{input: "l", expectErr: true},
		{input: "m", expectErr: true},
		{input: "n", expectErr: true},
		{input: "o", expectErr: true},
		{input: "p", expected: commitTypePerf},
		{input: "q", expectErr: true},
		{input: "r", expectErr: true},
		{input: "s", expected: commitTypeStyle},
		{input: "t", expected: commitTypeTest},
		{input: "u", expectErr: true},
		{input: "v", expectErr: true},
		{input: "w", expectErr: true},
		{input: "x", expectErr: true},
		{input: "y", expectErr: true},
		{input: "z", expectErr: true},

		// Short Cases
		{input: "ch", expected: commitTypeChore},
		{input: "fe", expected: commitTypeFeat},
		{input: "fi", expected: commitTypeFix},
		{input: "ref", expected: commitTypeRefactor},
		{input: "rev", expected: commitTypeRevert},

		// Bad Cases
		{input: "feature", expectErr: true},
		{input: "performance", expectErr: true},
		{input: "cicd", expectErr: true},
		{input: "cohre", expectErr: true},
		{input: "faet", expectErr: true},
		{input: "tst", expectErr: true},
	}

	commitTypes := GetCommitTypes()
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			msg := &Message{CommitType: test.input}
			err := msg.commitTypeFromInput(commitTypes)
			if test.expectErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if msg.CommitType != test.expected {
				t.Errorf(
					"Expected '%s'. Got '%s'",
					test.expected,
					msg.CommitType,
				)
			}
		})
	}
}

func TestCommiTTypeFromBranch(t *testing.T) {
	tests := []struct {
		hasBranch bool
		branch    string
		branchErr bool
		expected  string
		expectErr bool
	}{
		// Simple Cases
		{hasBranch: true, branch: "build", expected: commitTypeBuild},
		{hasBranch: true, branch: "chore", expected: commitTypeChore},
		{hasBranch: true, branch: "ci", expected: commitTypeCi},
		{hasBranch: true, branch: "docs", expected: commitTypeDocs},
		{hasBranch: true, branch: "feat", expected: commitTypeFeat},
		{hasBranch: true, branch: "fix", expected: commitTypeFix},
		{hasBranch: true, branch: "perf", expected: commitTypePerf},
		{hasBranch: true, branch: "refactor", expected: commitTypeRefactor},
		{hasBranch: true, branch: "revert", expected: commitTypeRevert},
		{hasBranch: true, branch: "style", expected: commitTypeStyle},
		{hasBranch: true, branch: "test", expected: commitTypeTest},

		// Fallback Cases
		{hasBranch: false, expected: CommitTypeFallback},
		{hasBranch: true, branch: "-", expected: CommitTypeFallback},
		{hasBranch: true, branch: "dev", expected: CommitTypeFallback},
		{hasBranch: true, branch: "a feature it is", expected: CommitTypeFallback},

		// Long Cases
		{
			hasBranch: true,
			branch:    "feature/fuzzy-flag-support",
			expected:  commitTypeFeat,
		},
		{
			hasBranch: true,
			branch:    "refactor/modules",
			expected:  commitTypeRefactor,
		},
		{
			hasBranch: true,
			branch:    "tests/commit-types",
			expected:  commitTypeTest,
		},
	}
	commitTypes := GetCommitTypes()
	for _, test := range tests {
		t.Run(test.branch, func(t *testing.T) {
			msg := &Message{hasBranch: test.hasBranch, branch: test.branch}
			if test.branchErr {
				msg.branchErr = errors.New("-")
			}
			err := msg.commitTypeFromBranch(commitTypes)
			if test.expectErr {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if msg.CommitType != test.expected {
				t.Errorf(
					"Expected '%s'. Got '%s'",
					test.expected,
					msg.CommitType,
				)
			}
		})
	}
}
