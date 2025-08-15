package model

// CommitMessage represents a conventional commit message structure
type CommitMessage struct {
	Type        string
	Scope       string
	Subject     string
	Body        string
	Breaking    bool
	FooterType  string
	FooterValue string
	Emoji       string
}

// ValidateSubject checks if the subject is valid
func (c *CommitMessage) ValidateSubject(maxLength int) bool {
	if len(c.Subject) == 0 {
		return false
	}

	// Calculate the length of the complete subject line
	prefix := c.Type
	if c.Scope != "" {
		prefix += "(" + c.Scope + ")"
	}
	if c.Breaking {
		prefix += "!"
	}
	prefix += ": "

	totalLength := len(prefix) + len(c.Subject)
	if c.Emoji != "" {
		totalLength += len(c.Emoji) + 1 // +1 for the space
	}

	return totalLength <= maxLength
}

// Format returns the formatted commit message
func (c *CommitMessage) Format() string {
	// Build the header
	header := ""

	// Add emoji if present
	if c.Emoji != "" {
		header += c.Emoji + " "
	}

	// Add type
	header += c.Type

	// Add scope if present
	if c.Scope != "" {
		header += "(" + c.Scope + ")"
	}

	// Add breaking change marker if applicable
	if c.Breaking {
		header += "!"
	}

	// Add subject
	header += ": " + c.Subject

	// Build the full message
	message := header

	// Add body if present
	if c.Body != "" {
		message += "\n\n" + c.Body
	}

	// Add footer if present
	if c.FooterType != "" && c.FooterValue != "" {
		message += "\n\n" + c.FooterType + ": " + c.FooterValue
	} else if c.Breaking && c.FooterType != "BREAKING CHANGE" {
		// Add BREAKING CHANGE footer if breaking is checked but no footer is specified
		message += "\n\nBREAKING CHANGE: Breaking changes were introduced in this commit."
	}

	return message
}
