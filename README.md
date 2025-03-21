# minipage - Minimal Static Page Generator

**minipage** is a simple static page generator designed to be easy to use without unnecessary complexity. If you'd like to render Markdown files into clean HTML with minimal effort, **minipage** is for you.

If you'd like to use **GitHub Pages** but find existing static site generators too complex or heavy, **minipage** is a lightweight alternative that gets the job done.

With **minipage**, you can render Markdown into clean, GitHub-like HTML without sidebars, banners, or other distractions. It generates one page at a time, giving you full control over your site's structure using tools like `make`.

## Features

- **Fast and Simple:** Converts a Markdown file into a single HTML page in one command.
- **Customizable Layout:** Optionally include a header, footer, and sidebar for layout consistency.
- **Lightweight by Design:** No configuration files — just straightforward options.
- **Flexible Workflow:** Designed to be called from `make` or similar build tools to construct full websites.

## Installation

### From GitHub Releases
Download the appropriate `.zip` file for your OS and architecture from the [releases page](https://github.com/hymkor/minipage/releases). Extract the archive and place the `minipage` executable in a directory listed in your `PATH`.

### Using Scoop (Windows)
For Windows users, you can install **minipage** via [Scoop](https://scoop.sh/):

```
scoop install https://raw.githubusercontent.com/hymkor/minipage/master/minipage.json
```

Alternatively, add the `hymkor` bucket and install:

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install minipage
```

## Usage

```
minipage {options} MARKDOWN.md > OUTPUT.html
```

### Options

- `-header HEADER.MD` — Include a Markdown file as the header.
- `-footer FOOTER.MD` — Include a Markdown file as the footer.
- `-sidebar SIDEBAR.MD` — Include a Markdown file as the sidebar.
- `-css CSSURL` — Specify a custom CSS URL (default: GitHub-like CSS).
- `-title TITLE` — Specify the page title.
- `-anchor-text string` — Specify the anchor text (default ".").

## Example

To build a simple page with a custom header and footer:

```
minipage -header header.md -footer footer.md content.md > index.html
```

Combine it with `make` for efficient site generation:

```makefile
all:
	minipage -header header.md -footer footer.md index.md > index.html
	minipage -header header.md -footer footer.md about.md > about.html
```

## Example Use Case

The website for [nyagos](https://nyaos.org/nyagos), a command-line shell for Windows, is built using **minipage**.

## Technical Details

- **Markdown to HTML Conversion:** Powered by [goldmark](https://github.com/yuin/goldmark), a CommonMark-compliant Markdown parser written in Go that is easy to extend and well-structured.
- **GitHub-like CSS:** Uses [github-markdown-css](https://github.com/sindresorhus/github-markdown-css), a minimal CSS file that replicates the GitHub Markdown style.

## Why minipage?

If you value:

✅ Clean, distraction-free Markdown rendering  
✅ Straightforward usage without complex configurations  
✅ Simple workflows that integrate with existing tools  

**minipage** is designed for you.


