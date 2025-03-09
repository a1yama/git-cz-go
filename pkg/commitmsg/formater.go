package commitmsg

import (
	"fmt"
	"strings"
)

// Format formats the commit message according to the conventional commit spec
func Format(
	commitType string,
	scope string,
	isBreaking bool,
	subject string,
	body string,
	footerType string,
	footerValue string,
	emoji string,
) string {
	// Build the header
	header := ""

	// Add emoji if present
	if emoji != "" {
		header += emoji + " "
	}

	// Add type
	header += commitType

	// Add scope if present
	if scope != "" {
		header += "(" + scope + ")"
	}

	// Add breaking change marker if applicable
	if isBreaking {
		header += "!"
	}

	// Add subject
	header += ": " + subject

	// Build the full message
	message := header

	// Add body if present
	if body != "" {
		message += "\n\n" + body
	}

	// Add footer if present
	if footerType != "" && footerValue != "" {
		message += "\n\n" + footerType
		if footerType == "BREAKING CHANGE" {
			message += ": " + footerValue
		} else {
			message += ": " + footerValue
		}
	} else if isBreaking && (footerType == "" || footerType != "BREAKING CHANGE") {
		// Add BREAKING CHANGE footer if breaking is checked but no footer is specified
		message += "\n\nBREAKING CHANGE: Breaking changes were introduced in this commit."
	}

	return message
}

// ValidateSubject checks if the subject meets the requirements
func ValidateSubject(subject string, maxLength int) (bool, string) {
	if len(subject) == 0 {
		return false, "Subject cannot be empty"
	}

	if len(subject) > maxLength {
		return false, fmt.Sprintf("Subject exceeds maximum length of %d characters", maxLength)
	}

	if strings.HasPrefix(strings.ToUpper(subject[:1]), subject[:1]) {
		return false, "Subject should not start with a capital letter"
	}

	if strings.HasSuffix(subject, ".") {
		return false, "Subject should not end with a period"
	}

	return true, ""
}

// ParseCommitMessage parses a commit message into its components
func ParseCommitMessage(message string) (string, string, bool, string, string, string, string) {
	var commitType, scope, subject, body, footerType, footerValue string
	var isBreaking bool

	// Split into header, body, and footer
	parts := strings.SplitN(message, "\n\n", 3)
	header := parts[0]

	// Parse body and footer if present
	if len(parts) > 1 {
		body = parts[1]
	}
	if len(parts) > 2 {
		footer := parts[2]
		footerParts := strings.SplitN(footer, ": ", 2)
		if len(footerParts) == 2 {
			footerType = footerParts[0]
			footerValue = footerParts[1]
		}
	}

	// Parse header
	// Check for breaking change marker
	if strings.Contains(header, "!:") {
		isBreaking = true
		header = strings.Replace(header, "!", "", 1)
	}

	// Check for scope
	if strings.Contains(header, "(") && strings.Contains(header, ")") {
		scopeStart := strings.Index(header, "(")
		scopeEnd := strings.Index(header, ")")

		if scopeStart < scopeEnd {
			commitType = header[:scopeStart]
			scope = header[scopeStart+1 : scopeEnd]
			subject = header[scopeEnd+2:] // +2 to skip "): "
		}
	} else {
		// No scope
		parts := strings.SplitN(header, ": ", 2)
		if len(parts) == 2 {
			commitType = parts[0]
			subject = parts[1]
		}
	}

	return commitType, scope, isBreaking, subject, body, footerType, footerValue
}
