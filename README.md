# minipage - Minimal Static Page Generator

[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/minipage.svg)](https://pkg.go.dev/github.com/hymkor/minipage)
[![Go Test](https://github.com/hymkor/minipage/actions/workflows/go.yml/badge.svg)](https://github.com/hymkor/minipage/actions/workflows/go.yml)

**minipage** is a simple static page generator designed to be easy to use without unnecessary complexity. If you'd like to render Markdown files into clean HTML with minimal effort, **minipage** is for you.

If you'd like to use **GitHub Pages** but find existing static site generators too complex or heavy, **minipage** is a lightweight alternative that gets the job done.

With **minipage**, you can render Markdown into clean, GitHub-like HTML without sidebars, banners, or other distractions. It generates one page at a time, giving you full control over your site's structure using tools like `make`.

## Features

- **Fast and Simple:** Converts one or more Markdown files into a single HTML page in one command.
- **Customizable Layout:** Optionally include a header, footer, and sidebar for layout consistency.
- **Lightweight by Design:** No configuration files — just straightforward options.
- **Flexible Workflow:** Designed to be called from `make` or similar build tools to construct full websites.

## Installation

### From GitHub Releases
Download the appropriate `.zip` file for your OS and architecture from the [releases page](https://github.com/hymkor/minipage/releases). Extract the archive and place the `minipage` executable in a directory listed in your `PATH`.

### Using `go install` (any OS)
To install **minipage** using Go, run the following command:

```
go install github.com/hymkor/minipage@latest
```

### Using Scoop (Windows)
For Windows users, you can install **minipage** via [Scoop](https://scoop.sh/):

```
scoop install https://raw.githubusercontent.com/hymkor/minipage/master/minipage.json
```

Alternatively, add the `hymkor` bucket and install:

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install hymkor/minipage
```

## Usage

```
minipage {options} FILE1.md [FILE2.md ...] > OUTPUT.html
```

### Options

- `-sidebar SIDEBAR.MD` — Include a Markdown file as the sidebar.
- `-css CSSURL` — Specify a custom CSS URL (default: GitHub-like CSS).
- `-title TITLE` — Specify the page title.
- `-anchor-text string` — Specify the anchor text (default `""`).
- `-title-from-file file` - Read the HTML title from the specified file
- `-readme-to-index` - Replace file names starting with 'README' with 'index' in relative anchor URLs
- `-outline-in-sidebar` - Output the outline in the sidebar

## Example

To build a simple page with multiple sections:

```
minipage header.md content.md footer.md > index.html
```

Combine it with `make` for efficient site generation:

```makefile
all:
	minipage header.md index.md footer.md > index.html
	minipage header.md about.md footer.md > about.html
```

## Example Use Case

The website for [nyagos](https://nyaos.org/nyagos), a command-line shell for Windows, is built using **minipage**.

You can also use minipage to generate GitHub Pages content for its own project site:

```Makefile
docs:
	"./minipage" -anchor-text "#" -readme-to-index -title "minipage - Minimal Static Page Generator" README.md > docs/index.html
	"./minipage" -anchor-text "#" -readme-to-index -title "Release Notes" release_note.md > docs/release_note.html
	"./minipage" -anchor-text "#" -readme-to-index -title "Release Notes(ja)" release_note_ja.md > docs/release_note_ja.html

.PHONY: all test dist _dist clean manifest release docs
```

This example generates clean HTML pages from Markdown source files for use with GitHub Pages.

## Technical Details

- **Markdown to HTML Conversion:** Powered by [goldmark](https://github.com/yuin/goldmark), a CommonMark-compliant Markdown parser written in Go that is easy to extend and well-structured.
- **GitHub-like CSS:** Uses [github-markdown-css](https://github.com/sindresorhus/github-markdown-css), a minimal CSS file that replicates the GitHub Markdown style.
- **Anchor URL Rewriting:** In relative anchor URLs, file extensions `.md` are automatically rewritten to `.html`. This allows links between Markdown files to remain functional after conversion.

## Why minipage?

If you value:

✅ Clean, distraction-free Markdown rendering  
✅ Straightforward usage without complex configurations  
✅ Simple workflows that integrate with existing tools  

**minipage** is designed for you.

## Release notes

- [English](release_note.md)
- [Japanese](release_note_ja.md)

## Author

[hymkor (HAYAMA Kaoru)](https://github.com/hymkor)

## License

MIT License
