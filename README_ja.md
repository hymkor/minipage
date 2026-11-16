# minipage - シンプルな静的ページジェネレータ

( [English](./README.md) / Japanese )

<!-- stdout: go run github.com/hymkor/example-into-readme/cmd/badges@latest -->
[![Go Test](https://github.com/hymkor/minipage/actions/workflows/go.yml/badge.svg)](https://github.com/hymkor/minipage/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/License-MIT-red)](https://github.com/hymkor/minipage/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/hymkor/minipage.svg)](https://pkg.go.dev/github.com/hymkor/minipage)
[![GitHub](https://img.shields.io/badge/github-repo-blue?logo=github)](https://github.com/hymkor/minipage)
<!-- -->

**minipage** は、余計な複雑さを排除し、簡単に使えることを目指した静的ページジェネレータです。Markdown ファイルを手軽に HTML に変換したい方に向いています。

**GitHub Pages** を使いたいけれど、既存の静的サイトジェネレータは複雑すぎる、あるいは重すぎると感じている方にも、**minipage** は軽量な選択肢となります。

**minipage** は、サイドバーやバナーなどの余計な装飾を排した、GitHub 風のすっきりした HTML を Markdown から生成します。1ページずつ生成する方式なので、`make` などのツールを使って、サイト全体の構成を自由に管理できます。

## 特徴

- **高速・シンプル:** 1つまたは複数の Markdown ファイルを、1つの HTML ページにまとめて1コマンドで変換できます。
- **レイアウトのカスタマイズ:** 必要に応じてヘッダー、フッター、サイドバーを追加できます。
- **軽量な設計:** 設定ファイルは不要で、シンプルなオプションだけで利用できます。
- **柔軟なワークフロー:** `make` などのビルドツールから呼び出して、Web サイト全体を構築する用途を想定しています。

## インストール

### GitHub Releases から

[リリースページ](https://github.com/hymkor/minipage/releases)から、お使いの OS とアーキテクチャに対応する `.zip` ファイルをダウンロードしてください。

アーカイブを展開し、`minipage` 実行ファイルを `PATH` に登録されているディレクトリに配置します。

### `go install` を使う場合（全 OS 対応）

Go を使ってインストールする場合は、次のコマンドを実行してください。

```sh
go install github.com/hymkor/minipage@latest
```

### Scoop を使う場合（Windows）

Windows では、[Scoop](https://scoop.sh/) を使ってインストールすることもできます。

```powershell
scoop install https://raw.githubusercontent.com/hymkor/minipage/master/minipage.json
```

または、`hymkor` バケットを追加してからインストールします。

```powershell
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install hymkor/minipage
```

## 使い方

```sh
minipage {options} SOURCE1 [SOURCE2 ...] > OUTPUT.html
```

`SOURCE` には、**Markdown ファイルのパス**または**HTML の文字列**を指定できます。

`minipage` は、指定された順番に入力を連結します。動作は `cat` コマンドに似ています。

### オプション

- `-sidebar SIDEBAR.MD` — Markdown ファイルをサイドバーとして追加します。
- `-css CSSURL` — CSS の URL を指定します（デフォルトは GitHub 風の CSS）。
- `-title TITLE` — ページのタイトルを指定します。
- `-anchor-text string` — アンカーのテキストを指定します（デフォルトは空文字列 `""`）。
- `-title-from-file file` — 指定したファイルから HTML のタイトルを読み込みます。
- `-readme-to-index` — 相対アンカー URL のうち、ファイル名が `README` で始まるものを `index` に置き換えます。
- `-outline-in-sidebar` — サイドバーに目次を出力します。

## 使用例

### 複数ファイルからの基本的な生成

複数のセクションからなるページを生成する場合：

```sh
minipage header.md content.md footer.md > index.html
```

### HTML 文字列を直接渡す

コマンドラインから HTML の断片を直接渡すこともできます。

これは、元の `README.md` には含めたくないロゴなどの要素を追加する場合に便利です。

入力が `<` で始まる場合、minipage はそれをファイル名ではなく HTML として認識します。そのため、シェルによるリダイレクトと解釈されないよう、文字列を引用符で囲んでください。

```bash
# 先頭の '<' によって、ファイルではなく HTML として認識される
minipage "<img src='logo.png' align='left' width='100' />" README.md > index.html
```

### `make` との併用

`make` と組み合わせることで、効率よくサイトを生成できます。

```makefile
all:
	minipage header.md index.md footer.md > index.html
	minipage header.md about.md footer.md > about.html
```

## 使用例

Windows 用コマンドラインシェルである [nyagos](https://nyaos.org/nyagos) の Web サイトは、**minipage** を使って構築されています。

また、minipage 自身のプロジェクトサイト用に、GitHub Pages のコンテンツを生成することもできます。

**Makefile**

```makefile
docs:
	"./minipage" -outline-in-sidebar -readme-to-index -title "minipage - Minimal Static Page Generator" README.md > docs/index.html
	"./minipage" -outline-in-sidebar -readme-to-index -title "CHANGELOG" CHANGELOG.md > docs/CHANGELOG.html
	"./minipage" -outline-in-sidebar -readme-to-index -title "CHANGELOG" CHANGELOG_ja.md > docs/CHANGELOG_ja.html
```

この例では、Markdown のソースファイルから、GitHub Pages で公開するためのシンプルな HTML ページを生成しています。

## 技術的な詳細

- **Markdown から HTML への変換:** Go 言語で実装された、CommonMark 準拠の拡張性・構造性に優れた Markdown パーサー [goldmark](https://github.com/yuin/goldmark) を使用しています。
- **GitHub 風 CSS:** GitHub の Markdown 表示に似せた、軽量な独自 CSS を使用しています。
- **アンカー URL の書き換え:** 相対アンカー URL に含まれる `.md` 拡張子は、自動的に `.html` に書き換えられます。これにより、変換後も Markdown ファイル間のリンクが機能します。

## minipage が向いている方

次のような点を重視する方に向けて、**minipage** は設計されています。

- 余計な装飾のない、すっきりとした Markdown 表示
- 複雑な設定を必要としない、簡単な使い方
- 既存のツールと組み合わせられるシンプルなワークフロー

## CHANGELOG

- [英語版](CHANGELOG.md)
- [日本語版](CHANGELOG_ja.md)

## 作者

[hymkor (HAYAMA Kaoru)](https://github.com/hymkor)

## ライセンス

MIT License
