Release notes
=============
[English](./release_note.md) / **Japanese** / [Top](./README.md)

v0.9.0
======
Oct 2, 2025

- サイドバーにページ内目次を追加するオプション: `-outline-in-sidebar` を追加

v0.8.1
======
Sep 29, 2025

- `hymkor/xnhttpd/idgen` のかわりに `hymkor/goldmark-mb-headingids` を使用
    - 見出し文字のうち、英数字ではないシングルバイト文字が ID に入っていて、GitHub 非互換になっていた問題を修正
    - 見出しに ID として使える文字が一文字もない時、`heading` ではなく、調査用に仮に設定していた`xheading` が ID として使われていた問題を修正

v0.8.0
------
Sep 28, 2025

- 見出しが日本語だった時、ID が GitHub 非互換になっていた問題を修正

v0.7.0
------
Sep 17, 2025

- XHTML形式で出力するようにした
- `<style>..</style>`タグの中身を`<!--..-->` で囲まないようにした

v0.6.0
------
May 31, 2025

- 見出しのアンカーリンクのスタイルを調整して、小さくて邪魔にならないように見えるようにした
- `-anchor-text` オプションが指定されていないときに、見出しの横にあるアンカーリンクを隠すようにした
- CSSの出力不具合: `%` が `%%` とエスケープされていなかったため、@media print ルールを壊していた

v0.5.0
------
May 30, 2025


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
