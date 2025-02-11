package domain

var commitEmojis = map[string]string{
	"feat":     "✨",
	"fix":      "🐛",
	"docs":     "📚",
	"style":    "💎",
	"refactor": "🔨",
	"perf":     "⚡",
	"test":     "✅",
	"chore":    "🛠",
}

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
	emoji := commitEmojis[c.Type]
	if c.Scope != "" {
		c.Scope = "(" + c.Scope + ")"
	}
	message := c.Type + c.Scope + ": " + emoji + " " + c.Summary
	if c.Body != "" {
		message += "\n\n" + c.Body
	}
	return message
}