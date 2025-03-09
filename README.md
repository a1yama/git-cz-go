# git-cz-go

A beautiful and interactive Conventional Commits CLI tool written in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea).

![git-cz-go demo](https://via.placeholder.com/800x400.png?text=git-cz-go+Demo)

## Features

- 💎 Beautiful TUI with keyboard navigation
- 🚀 Interactive prompts for all parts of the commit message
- 📋 Conventional Commits format support
- 🔍 Scope suggestions from your project structure
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

Simply run `git-cz` in a git repository to start the interactive commit process.

```bash
git-cz
```

You can also use it as a git subcommand if you add it to your PATH:

```bash
git cz
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
  "skipScope": false,
  "skipBody": false,
  "skipBreakingBody": false,
  "skipFooter": false,
  "maxSubjectLength": 100,
  "maxBodyLineLength": 100
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
go build -o git-cz ./cmd/git-cz-go

# Run
./git-cz
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