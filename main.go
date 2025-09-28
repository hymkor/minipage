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
	ghtml "github.com/yuin/goldmark/renderer/html"

	"go.abhg.dev/goldmark/anchor"

	"github.com/hymkor/exregexp-go"
	"github.com/hymkor/xnhttpd/idgen"
)

//go:embed github.css
var gitHubCss string

//go:embed our.css
var ourCss string

// Markdown is an object that converts Markdown documents to HTML.
// It supports optional anchor links next to headings and URL rewriting for links pointing to README.
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

// New creates a new Markdown instance.
// The anchorText argument specifies the text to be used for anchor links next to headings (e.g., "#").
// If an empty string is given, anchor links will not be generated.
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
		goldmark.WithRendererOptions(
			ghtml.WithXHTML(),
		),
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
	mdCtx := parser.NewContext(parser.WithIDs(idgen.New()))
	err = M.Convert(source, w, parser.WithContext(mdCtx))
	if class != "" {
		fmt.Fprintf(w, "</div><!-- \"%s\" -->\n", class)
	}
	return err
}

// Make generates an HTML page from the given Markdown file paths.
// The sidebar argument specifies the path to a sidebar Markdown file.
// The bodies argument is a list of paths to main content Markdown files.
// The title argument specifies the HTML <title> element content.
// The generated HTML is written to the writer w.
func (M *Markdown) Make(bodies []string, sidebar, css, title string, w io.Writer) error {
	fmt.Fprintln(w, `<!DOCTYPE html><html><head>`)
	if title != "" {
		fmt.Fprintf(w, "<title>%s</title>\n", html.EscapeString(title))
	}
	if css != "" {
		fmt.Fprintf(w, "<link rel=\"stylesheet\" href=\"%s\" />\n", css)
	} else {
		io.WriteString(w, "<style type=\"text/css\">\n")
		io.WriteString(w, gitHubCss)
		io.WriteString(w, ourCss)
		io.WriteString(w, "\n</style>\n")
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
	fmt.Fprintln(w, `</body></html>`)
	return nil
}

// EnableReadmeToIndex enables URL rewriting so that links ending with "README" are rewritten to "index".
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
