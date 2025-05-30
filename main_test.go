package main

import (
	"bytes"
	"testing"
)

func TestEnableReadmeToIndex1(t *testing.T) {
	M := New("#")
	M.EnableReadmeToIndex()

	result := M.rewriteLinks([]byte("[README_ja](./README_ja.md)"))
	expect := []byte("[README_ja](./index_ja.html)")

	if !bytes.Equal(result, expect) {
		t.Fatalf("expect %v, but %v", string(expect), string(result))
	}
}

func TestEnableReadmeToIndex2(t *testing.T) {
	M := New("#")
	M.EnableReadmeToIndex()

	result := M.rewriteLinks([]byte("[README_ja]: ./README_ja.md"))
	expect := []byte("[README_ja]: ./index_ja.html")

	if !bytes.Equal(result, expect) {
		t.Fatalf("expect %v, but %v", string(expect), string(result))
	}
}
