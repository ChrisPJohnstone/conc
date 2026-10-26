package message

import (
	"errors"
	"testing"
)

func TestParseScope(t *testing.T) {
	tests := []struct {
		input     string
		hasBranch bool
		branch    string
		branchErr bool
		expected  string
	}{
		// Input: Simple
		{input: "test", expected: "test"},

		// Input: Sentinel
		{input: "-", expected: "-"},

		// Branch: Regexp
		{hasBranch: true, branch: "feat/TEST-123", expected: "TEST-123"},
		{hasBranch: true, branch: "fix/ABC-456", expected: "ABC-456"},
		{hasBranch: true, branch: "long/LONGER-5555", expected: "LONGER-5555"},

		// Branch: Fallback
		{hasBranch: true, branch: "main", expected: scopeFallback},
		{hasBranch: true, branch: "dev", expected: scopeFallback},
		{hasBranch: true, branch: "feature", expected: scopeFallback},

		// Branch: Branch Error
		{hasBranch: true, branchErr: true, expected: scopeFallback},
	}
	for _, test := range tests {
		var test_name string
		if test.input != "" {
			test_name = test.input
		} else if test.branch != "" {
			test_name = test.branch
		}
		t.Run(test_name, func(t *testing.T) {
			msg := &Message{
				Scope:     test.input,
				hasBranch: test.hasBranch,
				branch:    test.branch,
			}
			if test.branchErr {
				msg.branchErr = errors.New("-")
			}
			msg.parseScope()
			if msg.Scope != test.expected {
				t.Errorf(
					"Expected '%s'. Got '%s'",
					test.expected,
					msg.Scope,
				)
			}
		})
	}
}
