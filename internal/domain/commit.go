package domain

type Commit struct {
	Type    string
	Scope   string
	Summary string
	Body    string
}

func NewCommit(commitType, scope, summary, body string) *Commit {
	return &Commit{
		Type:    commitType,
		Scope:   scope,
		Summary: summary,
		Body:    body,
	}
}

func (c *Commit) ConstructMessage() string {
	if c.Scope != "" {
		c.Scope = "(" + c.Scope + ")"
	}
	message := c.Type + c.Scope + ": " + c.Summary
	if c.Body != "" {
		message += "\n\n" + c.Body
	}
	return message
}