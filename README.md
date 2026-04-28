# coverlight

An alternative to `go tool cover -html` that supports light and dark mode coverage reports.

The standard `go tool cover` generates HTML coverage reports with a dark background. `coverlight` provides a light mode option where covered code is displayed in bold black and uncovered code appears as inverse text (white on black), making coverage gaps immediately visible.

This code is mostly based on the original `cover` tool included with the Go standard library.

## Install

```bash
go install github.com/jonbodner/coverlight@latest
```

## Usage

First, generate a coverage profile:

```bash
go test -coverprofile=coverage.out ./...
```

Then generate the HTML report:

```bash
# Open in browser (light mode, default)
coverlight -html coverage.out

# Write to file
coverlight -html coverage.out -o coverage.html

# Use dark mode
coverlight -html coverage.out -mode dark
```

## Flags

| Flag | Description |
|------|-------------|
| `-html` | Path to the coverage profile |
| `-o` | Output file (omit to open in browser) |
| `-mode` | `light` (default) or `dark` |