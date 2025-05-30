package main

import (
	"testing"
)

func testRewrite(t *testing.T, M *Markdown, source, expect string) {
	t.Helper()

	result := string(M.rewriteLinks([]byte(source)))

	if result != expect {
		t.Fatalf("expect %v, but %v", expect, result)
	}
}

func TestRewriteLinks(t *testing.T) {
	M := &Markdown{}
	testRewrite(t, M, "[README_ja](./README_ja.md)", "[README_ja](./README_ja.html)")
	testRewrite(t, M, "[README_ja]: ./README_ja.md", "[README_ja]: ./README_ja.html")
}

func TestEnableReadmeToIndex(t *testing.T) {
	M := &Markdown{}
	M.EnableReadmeToIndex()
	testRewrite(t, M, "[README_ja](./README_ja.md)", "[README_ja](./index_ja.html)")
	testRewrite(t, M, "[README_ja]: ./README_ja.md", "[README_ja]: ./index_ja.html")
}
