all:
	go fmt ./...
	go build

github.css :
	curl https://raw.githubusercontent.com/sindresorhus/github-markdown-css/gh-pages/github-markdown-light.css > github.css
