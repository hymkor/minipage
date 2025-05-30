Release notes
=============

[English](./release_note.md) / **Japanese** / [Top](./README.md)

- ファイル名として与えられた `-` は標準入力として扱うようにした。
- **追加**: `-readme-to-index` オプションを追加した。相対リンク中の `README` で始まるファイル名（例：`README_ja`）を `index` に変換する（例：`index_ja`）。拡張子 `.md → .html` の変換や、相対 URL の処理はこのオプションに関係なく常に行われる。

v0.4.0
------
May 7, 2025

- 生成されるHTMLに `<!DOCTYPE html>` を挿入するようにした
- 指定したファイルの中味を HTML のタイトルとして使うオプション `-title-from-file` を実装

v0.3.0
------
Mar 22, 2025

- `-header`, `-footer` オプションを削除。markdownファイルはコマンドラインで直接指定するようにした。

```
minipage HEADER.md BODY.md FOOTER.md > OUTPUT.html
```

v0.2.0
------
Mar 21, 2025

- オプション `-anchor-text STRING` を追加

v0.1.0
------
Mar 20, 2025

- 初版
