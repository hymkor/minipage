package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"html"
	"io"
	"os"
	"regexp"
	"runtime"
	"slices"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	// goldmarkHtml "github.com/yuin/goldmark/renderer/html"

	"go.abhg.dev/goldmark/anchor"

	"github.com/hymkor/exregexp-go"
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
		div.main{ width:100%% }
	}
	a.permalink{
		font-size: 0.75em;
		vertical-align: super;
		text-decoration: none;
		margin-left: 0.3em;
		color: #888;
	}
	a.permalink:hover{
		color: #000;
		text-decoration: underline;
	}
// -->
</style>`

const htmlFooter = `</body></html>`

type Markdown struct {
	goldmark.Markdown
	linkRewriters []func([]byte) []byte
}

type customTexter []byte

func (c customTexter) AnchorText(h *anchor.HeaderInfo) []byte {
	if h.Level == 1 {
		return nil
	}
	return []byte(c)
}

func New(anchorText string) *Markdown {
	ext := []goldmark.Extender{
		extension.Table,
		extension.NewLinkify(
			extension.WithLinkifyAllowedProtocols([][]byte{
				[]byte("http:"),
				[]byte("https:"),
			})),
		extension.TaskList,
		extension.Footnote,
		meta.New(meta.WithTable()),
	}
	if anchorText != "" {
		ext = append(ext, &anchor.Extender{
			Texter:     customTexter(anchorText),
			Attributer: anchor.Attributes{"class": "permalink"}})
	}
	options := []goldmark.Option{
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID()),
		goldmark.WithExtensions(ext...),
	}
	return &Markdown{
		Markdown: goldmark.New(options...),
	}
}

var (
	rxAnchor1 = regexp.MustCompile(`(\[.*?\]\()(.*?)\.md\)`)
	rxAnchor2 = regexp.MustCompile(`(?m)^(\[.*?\]:\s+)(.*?)\.md\s*$`)
)

func readFileOrStdin(path string) ([]byte, error) {
	if path == "-" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(path)
}

func (M *Markdown) rewriteLinks(source []byte) []byte {
	source = exregexp.ReplaceAllSubmatchFunc(rxAnchor1, source, func(s [][]byte) []byte {
		url := s[2]
		if bytes.HasPrefix(url, []byte("http")) {
			return s[0]
		}
		for _, f := range M.linkRewriters {
			url = f(url)
		}
		return slices.Concat(s[1], url, []byte(".html)"))
	})
	source = exregexp.ReplaceAllSubmatchFunc(rxAnchor2, source, func(s [][]byte) []byte {
		url := s[2]
		if bytes.HasPrefix(url, []byte("http")) {
			return s[0]
		}
		for _, f := range M.linkRewriters {
			url = f(url)
		}
		return slices.Concat(s[1], url, []byte(".html"))
	})
	return source
}

func (M *Markdown) makePage(path, class string, w io.Writer) error {
	if path == "" {
		return nil
	}
	source, err := readFileOrStdin(path)
	if err != nil {
		return err
	}
	source = M.rewriteLinks(source)
	if class != "" {
		fmt.Fprintf(w, "<div class=\"%s\">\n", class)
	}
	err = M.Convert(source, w)
	if class != "" {
		fmt.Fprintf(w, "</div><!-- \"%s\" -->\n", class)
	}
	return err
}

func (M *Markdown) Make(bodies []string, sidebar, css, title string, w io.Writer) error {
	fmt.Fprintln(w, `<!DOCTYPE html><html><head>`)
	if title != "" {
		fmt.Fprintf(w, "<title>%s</title>\n", html.EscapeString(title))
	}
	if css != "" {
		fmt.Fprintf(w, "<link rel=\"stylesheet\" href=\"%s\" />\n", css)
	} else {
		fmt.Fprintf(w, htmlHeader, gitHubCss)
	}
	fmt.Fprintln(w, "</head><body>")

	fmt.Fprintln(w, "<div class=\"main markdown-body\">")
	for _, body := range bodies {
		if err := M.makePage(body, "", w); err != nil {
			return err
		}
	}
	fmt.Fprintln(w, "</div><!-- \"main\" -->")
	if err := M.makePage(sidebar, "sidebar markdown-body", w); err != nil {
		return err
	}
	fmt.Fprintln(w, htmlFooter)
	return nil
}

func (M *Markdown) EnableReadmeToIndex() {
	M.linkRewriters = append(M.linkRewriters, func(s []byte) []byte {
		return bytes.ReplaceAll(s, []byte("README"), []byte("index"))
	})
}

var (
	flagSidebar    = flag.String("sidebar", "", "Include a Markdown file as the sidebar")
	flagCSS        = flag.String("css", "", "Specify a custom CSS URL (default: GitHub-like CSS).")
	flagTitle      = flag.String("title", "", "Specify the page title")
	flagAnchorText = flag.String("anchor-text", "", "Specify the anchor text")
	flagTitleFile  = flag.String("title-from-file", "", "Read the HTML title from the specified `file`")

	flagReadmeToIndex = flag.Bool("readme-to-index", false,
		"Replace file names starting with 'README' with 'index' in relative anchor URLs")
)

func mains(args []string) error {
	m := New(*flagAnchorText)
	if len(args) <= 0 {
		fmt.Fprintf(os.Stderr, "minipage %s-%s-%s/%s\n\n",
			version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		flag.Usage()
		return nil
	}
	if *flagReadmeToIndex {
		m.EnableReadmeToIndex()
	}
	title := *flagTitle
	if title == "" && *flagTitleFile != "" {
		if b, err := readFileOrStdin(*flagTitleFile); err != nil {
			return err
		} else {
			title = string(b)
		}
	}
	return m.Make(args, *flagSidebar, *flagCSS, title, os.Stdout)
}

var version string

func main() {
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
