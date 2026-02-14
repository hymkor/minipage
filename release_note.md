Release notes
=============
**English** / [Japanese](./release_note_ja.md) / [Top](./README.md)

- Use Go 1.20.14 instead of latest Go (1.24.1) if possible. (#4)
  - Stop using "slices"
  - Downgrade go.abhg.dev/goldmark/anchor from v0.2.0 to v0.1.1

v0.11.0
-------
Feb 13, 2026

- Changed: Markdown files now support raw HTML. Embedded HTML tags will be preserved and rendered directly in the output. (#1)
- Added: Support for inline HTML arguments. You can now pass raw HTML strings directly as command-line arguments without creating a separate file. (#2)

v0.10.0
-------
Oct 19, 2025

- In outline output, the indent depth before the Markdown list marker was previously calculated as **(heading level × 4)**. However:

  - When the heading level is two or more levels deep, it may not be recognized as nested.
  - Some documents start directly with `##`.

  To address this, the indent is now calculated as **((heading level − top-level heading level) × 2)**.

v0.9.0
------
Oct 2, 2025

- Added the option `-outline-in-sidebar` to output the outline in the sidebar

v0.8.1
------
Sep 29, 2025

- Use `hymkor/goldmark-mb-headingids` instead of `hymkor/xnhttpd/idgen`
    - Fix the issue where non-alphanumeric single-byte characters in headings were included in IDs, making them incompatible with GitHub.
    - Fix the issue where, when a heading contained no characters usable for an ID, the placeholder `xheading` was used instead of `heading`.

v0.8.0
------
Sep 28, 2025

- Fix the issue where IDs were not compatible with GitHub when headers contained non-ASCII characters.

v0.7.0
------
Sep 17, 2025

- Changed output format to XHTML
- Stopped wrapping the contents of `<style>..</style>` tags with `<!--..-->`

v0.6.0
------
May 31, 2025

- Adjusted the style of heading anchor links to appear smaller and less intrusive
- Anchor links next to headings are now hidden when the `-anchor-text` option is not specified
- Fixed: CSS output bug: % was not escaped as %%, breaking @media print rule.

v0.5.0
------
May 30, 2025

- Treat `-` as stdin
- **Added** the `-readme-to-index` option, which rewrites file names starting with `README` to start with `index` in relative anchor URLs.  
    This is useful when generating links that point to `index.html` instead of `README.html` (e.g., `README_ja.md` → `index_ja.html`).

v0.4.0
------
May 7, 2025

- Insert `<!DOCTYPE html>` into generated HTML
- Implement `-title-from-file` option to use the contents of a specified file as the HTML title

v0.3.0
------
Mar 22, 2025

- Removed `-header` and `-footer` options. Markdown files are now specified directly in the command:

```
minipage HEADER.md BODY.md FOOTER.md > OUTPUT.html
```

v0.2.0
------
Mar 21, 2025

- Added the option `-anchor-text STRING`

v0.1.0
------
Mar 20, 2025

- Initial release
