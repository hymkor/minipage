Release notes
=============

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

