package main

import (
	"dev09/htmlparser"
	"dev09/wget"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"testing"

	"golang.org/x/net/html"
)

func TestWget(t *testing.T) {
	w, err := wget.NewWget("https://go.dev/", wget.Flags{
		R: false,
	})

	if err != nil {
		t.Fatal(err)
	}

	dir := w.Dir()
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	if dir == "" {
		t.Fatal("директория не должна быть пустой")
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 1 {
		t.Fatal("ожидаемая количество файлов 1. Получили:", len(files))
	}
}

func TestWgetR(t *testing.T) {
	w, err := wget.NewWget("https://go.dev/learn/", wget.Flags{
		R: true,
	})

	if err != nil {
		t.Fatal(err)
	}

	dir := w.Dir()
	assets := w.DirAssets()
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	if dir == "" {
		t.Fatal("директория не должна быть пустой")
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) == 0 {
		t.Fatal("главная директория не должна быть пустой")
	}

	files, err = ioutil.ReadDir(assets)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) == 0 {
		t.Fatal("ассеты не должны быть пустыми")
	}
}

func TestParseHTML(t *testing.T) {
	html := htmlparser.HTMLPath{
		FilePath:  path.Join("assets", "test.html"),
		ClearLink: "https://test.com",
	}

	got, err := htmlparser.ParseHTML(html)
	if err != nil {
		t.Fatal(err)
	}

	want := []string{
		"https://test.com/1",
		"https://test.com/2",
		"https://test.com/3",
		"https://test.com/4",
		"https://test.com/5",
		"https://test.com/6",
	}

	for i, l := range want {
		if l != got[i] {
			t.Fatalf("want[%d] != got[%d]; want: %v, got: %v", i, i, want, got)
		}
	}
}

func TestBuildClearLink(t *testing.T) {
	url, _ := url.Parse("https://123.22.com/aaaa#4?d=2")
	want := "https://123.22.com"
	got := htmlparser.BuildClearLink(url)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func TestParseTags(t *testing.T) {
	attrs := []html.Attribute{
		{Namespace: "", Key: "href", Val: "123"},
	}
	want := "123"

	got := htmlparser.ParseTag(attrs, "href")

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
