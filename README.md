# git-cz-go

A beautiful and interactive Conventional Commits CLI tool written in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea).

![Demo](docs/assets/demo.gif)

## Features

- 💎 Beautiful TUI with keyboard navigation
- 🚀 Interactive prompts for all parts of the commit message
- 📋 Full Conventional Commits format support
  - Type, scope, subject, body, breaking changes, and footer
- 🔍 Scope input for categorizing changes
- 📝 Multi-line body editor for detailed descriptions
- ⚠️ Breaking change indicator
- 🔗 Footer for issue references and metadata
- 🔢 Quick selection using number keys (1-9) and Y/N shortcuts
- 😀 Optional emoji support
- ⚙️ Customizable via configuration file
- 🌈 Color-coded interface

## Installation

### Using Go

```bash
go install github.com/a1yama/git-cz-go/cmd/git-cz-go@latest
```

### From Releases

Download the appropriate binary for your platform from the [GitHub Releases](https://github.com/a1yama/git-cz-go/releases) page.

## Usage

Simply run `git-cz-go` in a git repository to start the interactive commit process.

```bash
git-cz-go
```

### Navigation

- **Commit Type**: Use arrow keys or press 1-9 for quick selection
- **Scope**: Enter scope manually (optional, press Enter to skip)
- **Subject**: Type your commit message
- **Body**: Multi-line editor (Ctrl+D to continue, optional)
- **Breaking Change**: Use arrow keys or Y/N for quick selection
- **Footer**: Enter issue references (optional, press Enter to skip)
- **General**: 
  - `Esc` to go back to previous step
  - `Ctrl+C` to cancel
  - `Enter` to confirm/continue

To view the list of available commit types:

```bash
git-cz-go --types
```

You can also create an alias in your git config:

```bash
git config --global alias.cz "!git-cz-go"
```

After setting up this alias, you can simply use:

```bash
git cz
```

Or to view commit types:

```bash
git cz --types
```

## Configuration

git-cz-go can be configured using a JSON file. The configuration file is searched for in the following locations:

1. `./.git-cz.json` (current directory)
2. `~/.git-cz.json` (home directory)
3. `~/.config/git-cz/config.json` (XDG config directory)

Example configuration:

```json
{
  "types": [
    {
      "type": "feat",
      "description": "A new feature",
      "emoji": "✨"
    },
    {
      "type": "fix",
      "description": "A bug fix",
      "emoji": "🐛"
    }
  ],
  "useEmoji": true,
  "maxSubjectLength": 100
}
```

## Development

### Prerequisites

- Go 1.20 or higher

### Build from source

```bash
# Clone the repository
git clone https://github.com/a1yama/git-cz-go.git
cd git-cz-go

# Build
go build -o git-cz-go ./cmd/git-cz-go

# Run
./git-cz-go
```

### Creating a release

1. Create a tag following semantic versioning
   ```
   git tag -a v0.1.0 -m "First release"
   git push origin v0.1.0
   ```

2. GitHub Actions will automatically build and publish the release

## License

MIT
