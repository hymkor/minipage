package outline

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strings"
	"unicode"
)

type Header struct {
	Level int
	Title string
	ID    string
}

var (
	rxHeader  = regexp.MustCompile(`^(#+)\s+`)
	rxMinuses = regexp.MustCompile(`-+`)
	rxHeader1 = regexp.MustCompile(`^====*\s*$`)
	rxHeader2 = regexp.MustCompile(`^----*\s*$`)
	rxVerb    = regexp.MustCompile("^\\s*```")
)

func slugify(s string) string {
	s = strings.ToLower(s)

	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
		} else if unicode.IsSpace(r) {
			b.WriteRune('-')
		} else if r == '-' || r == '_' {
			b.WriteRune(r)
		}
	}
	s = b.String()
	s = rxMinuses.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

type Slugs struct {
	used map[string]int
}

func NewSlugs() *Slugs {
	return &Slugs{used: map[string]int{}}
}

func (S *Slugs) Make(s string) string {
	s = slugify(s)
	if count, ok := S.used[s]; ok {
		S.used[s] = count + 1
		s = fmt.Sprintf("%s-%d", s, count)
	} else {
		S.used[s] = 1
	}
	return url.PathEscape(s)
}

func Make(fname string) ([]*Header, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	return FromReader(fd)
}

func FromReader(r io.Reader) ([]*Header, error) {
	sc := bufio.NewScanner(r)
	lastlast := ""
	last := ""
	slugs := NewSlugs()
	headers := []*Header{}
	verb := false
	for sc.Scan() {
		line := sc.Text()
		if rxVerb.MatchString(line) {
			verb = !verb
		}
		if !verb {
			if m := rxHeader.FindStringSubmatch(line); m != nil {
				level := len(m[1])
				title := strings.TrimSpace(line[len(m[0]):])
				id := slugs.Make(title)
				header := &Header{
					Level: level,
					Title: title,
					ID:    id,
				}
				headers = append(headers, header)
			} else if lastlast == "" && rxHeader1.MatchString(line) {
				header := &Header{
					Level: 1,
					Title: last,
					ID:    slugs.Make(last),
				}
				headers = append(headers, header)
			} else if lastlast == "" && rxHeader2.MatchString(line) {
				header := &Header{
					Level: 2,
					Title: last,
					ID:    slugs.Make(last),
				}
				headers = append(headers, header)
			}
		}
		lastlast = last
		last = line
	}
	return headers, sc.Err()
}

func (h *Header) WriteTo(baseUrl string, w io.Writer) (int, error) {
	n := 0
	for i := 1; i < h.Level; i++ {
		_n, err := io.WriteString(w, "    ")
		n += _n
		if err != nil {
			return n, err
		}
	}
	_n, err := fmt.Fprintf(w, "- [%s](%s#%s)", h.Title, baseUrl, h.ID)
	n += _n
	return n, err
}

func List(headers []*Header, baseUrl, newline string, w io.Writer) (int, error) {
	n := 0
	for _, h := range headers {
		_n, err := h.WriteTo(baseUrl, w)
		n += _n
		if err != nil {
			return n, err
		}
		_n, err = io.WriteString(w, newline)
		n += _n
		if err != nil {
			return n, err
		}
	}
	return n, nil
}
