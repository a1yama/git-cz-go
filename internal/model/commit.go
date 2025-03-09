package model

// CommitMessage represents a conventional commit message structure
type CommitMessage struct {
	Type    string
	Subject string
	Emoji   string
}

// ValidateSubject checks if the subject is valid
func (c *CommitMessage) ValidateSubject(maxLength int) bool {
	if len(c.Subject) == 0 {
		return false
	}

	// Calculate the length of the complete subject line
	// type: subject
	prefix := c.Type + ": "
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

	// Add type and subject
	header += c.Type + ": " + c.Subject

	return header
}
