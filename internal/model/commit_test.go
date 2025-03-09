package model

import (
	"testing"
)

func TestCommitMessageFormat(t *testing.T) {
	testCases := []struct {
		name     string
		message  CommitMessage
		expected string
	}{
		{
			name: "Simple message",
			message: CommitMessage{
				Type:    "feat",
				Subject: "add new feature",
			},
			expected: "feat: add new feature",
		},
		{
			name: "With emoji",
			message: CommitMessage{
				Type:    "fix",
				Subject: "resolve issue",
				Emoji:   "🐛",
			},
			expected: "🐛 fix: resolve issue",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.message.Format()
			if result != tc.expected {
				t.Errorf("Format() = %q, want %q", result, tc.expected)
			}
		})
	}
}

func TestValidateSubject(t *testing.T) {
	testCases := []struct {
		name      string
		message   CommitMessage
		maxLength int
		expected  bool
	}{
		{
			name: "Valid subject",
			message: CommitMessage{
				Type:    "feat",
				Subject: "add new feature",
			},
			maxLength: 100,
			expected:  true,
		},
		{
			name: "Empty subject",
			message: CommitMessage{
				Type:    "feat",
				Subject: "",
			},
			maxLength: 100,
			expected:  false,
		},
		{
			name: "Subject too long",
			message: CommitMessage{
				Type:    "feat",
				Subject: "this is an extremely long subject line that exceeds the maximum length allowed for a commit message subject line",
			},
			maxLength: 50,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.message.ValidateSubject(tc.maxLength)
			if result != tc.expected {
				t.Errorf("ValidateSubject(%d) = %v, want %v", tc.maxLength, result, tc.expected)
			}
		})
	}
}
