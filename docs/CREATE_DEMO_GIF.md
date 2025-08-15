# Creating Demo GIF for git-cz-go

This guide explains how to create a demo GIF showcasing git-cz-go functionality.

## Prerequisites

### macOS
```bash
# Install terminal recorder
brew install asciinema

# Install gif converter
brew install agg
```

### Alternative: Using Gifski
```bash
# Install screen recording tool
brew install --cask kap
```

## Method 1: Using asciinema + agg

### Step 1: Record Terminal Session
```bash
# Start recording
asciinema rec demo.cast

# Run git-cz-go and demonstrate all features:
./git-cz-go

# During the demo, showcase:
# 1. Select commit type (use arrow keys and number shortcuts)
# 2. Enter scope (e.g., "ui")
# 3. Enter subject (e.g., "add interactive commit message builder")
# 4. Enter body (multi-line description)
# 5. Select breaking change (No)
# 6. Skip footer (press Enter)
# 7. Confirm the commit

# Stop recording with Ctrl+D
```

### Step 2: Convert to GIF
```bash
# Convert asciinema recording to gif
agg demo.cast demo.gif

# With custom settings
agg --theme monokai --font-size 14 demo.cast demo.gif
```

## Method 2: Using Kap (Recommended for macOS)

1. Open Kap
2. Select the terminal window area
3. Start recording
4. Run through the demo:
   ```bash
   ./git-cz-go
   ```
5. Stop recording
6. Export as GIF with settings:
   - FPS: 10
   - Quality: High
   - Size: Optimize for web

## Method 3: Using ttygif

```bash
# Install ttygif
brew install ttygif

# Record session
ttyrec demo

# Run git-cz-go
./git-cz-go

# Exit recording (Ctrl+D)

# Convert to gif
ttygif demo

# Optimize the gif
gifsicle -O3 --colors 256 tty.gif > demo.gif
```

## Demo Script

Here's the recommended flow for the demo:

```bash
# 1. Start git-cz-go
./git-cz-go

# 2. Select commit type
# - Show arrow navigation
# - Show number shortcut (press "1" for feat)

# 3. Enter scope
# Type: "ui"

# 4. Enter subject  
# Type: "add complete conventional commit flow"

# 5. Enter body (optional)
# Type: "Implemented all conventional commit fields including:
# - Scope for categorizing changes
# - Body for detailed descriptions
# - Breaking change indicator
# - Footer for issue references"
# Press Ctrl+D

# 6. Breaking changes?
# Select: No (or press 'n')

# 7. Footer (optional)
# Press Enter to skip

# 8. Confirm
# Show the preview
# Select: Yes
```

## Optimizing the GIF

```bash
# Reduce file size
gifsicle -O3 --lossy=30 -k 256 demo.gif -o demo-optimized.gif

# Resize if needed
gifsicle --resize-width 800 demo.gif -o demo-resized.gif
```

## Adding to README

Once you have the GIF, place it in the repository:

```bash
# Create assets directory
mkdir -p docs/assets

# Move gif to assets
mv demo.gif docs/assets/demo.gif

# The README will reference it as:
# ![Demo](docs/assets/demo.gif)
```

## Tips for a Good Demo GIF

1. **Keep it short**: 30-60 seconds maximum
2. **Clear terminal**: Use a clean terminal with good contrast
3. **Consistent speed**: Don't rush through inputs
4. **Show features**: Demonstrate key features like number shortcuts
5. **Terminal size**: Use a reasonable terminal size (80x24 or 100x30)
6. **Font size**: Ensure text is readable in the GIF

## Example Terminal Settings

For best results:
- Terminal: iTerm2 or Terminal.app
- Theme: Dark theme with good contrast (Dracula, One Dark, etc.)
- Font: SF Mono or JetBrains Mono, 14-16pt
- Window size: ~100 columns x 30 rows