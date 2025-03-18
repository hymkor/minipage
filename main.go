package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	// goldmarkHtml "github.com/yuin/goldmark/renderer/html"
)

//go:embed github.css
var gitHubCss string

const htmlHeader = `<style type="text/css"><!--
%s
	.markdown-body {
		box-sizing: border-box;
		max-width: 980px;
		margin: 0 auto;
		padding: 45px;
	}
	body{
		display:flex;
		overflow-y: scroll;
	}
	@media screen{
		div.sidebar{ width:30%% }
		div.main{ width:70%% }
	}

	@media print{
		div.sidebar,div.footer{ display:none }
		div.main{ width:100% }
	}

// -->
</style>`

const htmlFooter = `</body></html>`

type Markdown struct {
	goldmark.Markdown
}

func New() *Markdown {
	options := []goldmark.Option{
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID()),
		goldmark.WithExtensions(
			extension.Table,
			extension.NewLinkify(
				extension.WithLinkifyAllowedProtocols([][]byte{
					[]byte("http:"),
					[]byte("https:"),
				})),
			extension.TaskList,
			extension.Footnote,
			meta.New(meta.WithTable())),
	}
	return &Markdown{
		Markdown: goldmark.New(options...),
	}
}

var (
	rxAnchor1 = regexp.MustCompile(`(\[.*?\]\(.*?\.)md\)`)
	rxAnchor2 = regexp.MustCompile(`(?m)^\[.*?\]:\s+.*?\.md\s*$`)
)

func concat(a []byte, b []byte) []byte {
	r := make([]byte, 0, len(a)+len(b))
	r = append(r, a...)
	r = append(r, b...)
	return r
}

func chomp(s []byte) []byte {
	for len(s) > 0 && bytes.IndexByte([]byte{'\r', '\n', ' '}, s[len(s)-1]) >= 0 {
		s = s[:len(s)-1]
	}
	return s
}

func (M *Markdown) makePage(path, class string, w io.Writer) error {
	if path == "" {
		return nil
	}
	source, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	source = rxAnchor1.ReplaceAllFunc(source, func(s []byte) []byte {
		if len(s) < len("md)") {
			return s
		}
		return concat(s[:len(s)-3], []byte("html)"))
	})
	source = rxAnchor2.ReplaceAllFunc(source, func(s []byte) []byte {
		s = chomp(s)
		if len(s) < len("md") {
			return s
		}
		return concat(s[:len(s)-2], []byte("html"))
	})
	// println(string(source))

	fmt.Fprintf(w, "<div class=\"%s\">\n", class)
	err = M.Convert(source, w)
	fmt.Fprintf(w, "</div><!-- \"%s\" -->\n", class)
	return err
}

func withoutExt(path string) string {
	bodyLen := len(path) - len(filepath.Ext(path))
	return path[:bodyLen]
}

func (M *Markdown) Make(body, header, sidebar, footer, css string, w io.Writer) error {
	fmt.Fprintln(w, `<html><head>`)
	if css != "" {
		fmt.Fprintf(w, "<link rel=\"stylesheet\" href=\"%s\" />\n", css)
	} else {
		fmt.Fprintf(w, htmlHeader, gitHubCss)
	}
	fmt.Fprintln(w, "</head><body>")

	fmt.Fprintln(w, "<div class=\"main markdown-body\">")
	if err := M.makePage(header, "header", w); err != nil {
		return err
	}
	if err := M.makePage(body, "body", w); err != nil {
		return err
	}
	if err := M.makePage(footer, "footer", w); err != nil {
		return err
	}
	fmt.Fprintln(w, "</div><!-- \"main\" -->")
	if err := M.makePage(sidebar, "sidebar markdown-body", w); err != nil {
		return err
	}
	fmt.Fprintln(w, htmlFooter)
	return nil
}

var (
	flagSidebar = flag.String("sidebar", "", "Specify sidebar")
	flagHeader  = flag.String("header", "", "Specify header")
	flagFooter  = flag.String("footer", "", "Specify footer")
	flagCSS     = flag.String("css", "", "Specify CSS URL")
)

func mains(args []string) error {
	m := New()
	if len(args) <= 0 {
		return io.EOF
	}
	return m.Make(args[0], *flagHeader, *flagSidebar, *flagFooter, *flagCSS, os.Stdout)
}

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
