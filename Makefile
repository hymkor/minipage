ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=set
    DEL=del
    NUL=nul
else
    SET=export
    DEL=rm
    NUL=/dev/null
endif

ifndef GO
    SUPPORTGO=go1.20.14
    GO:=$(shell $(WHICH) $(SUPPORTGO) 2>$(NUL) || echo go)
endif

NAME:=$(subst go-,,$(notdir $(CURDIR)))
VERSION:=$(shell git describe --tags 2>$(NUL) || echo v0.0.0)
GOOPT:=-ldflags "-s -w -X main.version=$(VERSION)"
EXE:=$(shell $(GO) env GOEXE)

all:
	$(GO) fmt ./...
	$(SET) "CGO_ENABLED=0" && $(GO) build $(GOOPT)

test:
	$(GO) test -v

_dist:
	$(SET) "CGO_ENABLED=0" && $(GO) build $(GOOPT)
	zip -9 $(NAME)-$(VERSION)-$(GOOS)-$(GOARCH).zip $(NAME)$(EXE)

dist:
	$(SET) "GOOS=linux"   && $(SET) "GOARCH=amd64" && $(MAKE) _dist
	$(SET) "GOOS=windows" && $(SET) "GOARCH=amd64" && $(MAKE) _dist

clean:
	$(DEL) *.zip $(NAME)$(EXE)

manifest:
	$(GO) run github.com/hymkor/make-scoop-manifest@latest *-windows-*.zip > $(NAME).json

release:
	$(GO) run github.com/hymkor/latest-notes@latest | gh release create -d --notes-file - -t $(VERSION) $(VERSION) $(wildcard $(NAME)-$(VERSION)-*.zip)

docs:
	"./minipage" -outline-in-sidebar -readme-to-index -title "minipage - Minimal Static Page Generator" README.md > docs/index.html
	"./minipage" -outline-in-sidebar -readme-to-index -title "Release Notes" release_note.md > docs/release_note.html
	"./minipage" -outline-in-sidebar -readme-to-index -title "Release Notes(ja)" release_note_ja.md > docs/release_note_ja.html

.PHONY: all test dist _dist clean manifest release docs
