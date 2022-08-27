package htmlparser

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// HTMLPath содержит в себе путь до html документа
// и абсолютный путь (без # и тд) документа на сайте
type HTMLPath struct {
	FilePath  string
	ClearLink string
}

// BuildClearLink знает как удалить из ссылки все лишнее
func BuildClearLink(link *url.URL) string {
	var sb strings.Builder
	sb.WriteString(link.Scheme)
	sb.WriteString("://")
	sb.WriteString(link.Hostname())

	return sb.String()
}

// ParseHTML знает как распарсить документ
// и вернуть найденные абсолютные ссылки
func ParseHTML(htmlpath HTMLPath) ([]string, error) {
	links := []string{}
	file, err := os.Open(htmlpath.FilePath)
	if err != nil {
		return links, err
	}
	defer file.Close()
	doc, err := html.Parse(file)
	if err != nil {
		return links, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "a" || n.Data == "link" || n.Data == "img" {
				attrkey := "href"
				if n.Data == "img" {
					attrkey = "src"
				}
				link := ParseTag(n.Attr, attrkey)
				link, err := makeFullLink(link, htmlpath.ClearLink)
				if err != nil {
					fmt.Println(err)
					return
				}
				if link != "" {
					links = append(links, link)
				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links, nil
}

// ParseTag знает как вытащить ссылку из тега по переданному ключу
func ParseTag(attrs []html.Attribute, attrkey string) string {
	var link string
	for _, attr := range attrs {
		if attr.Key == attrkey {
			link = attr.Val
			break
		}
	}

	return link
}

// makeFullLink знает как сделать абсолютный путь из ссылки.
// Если передан абсолютный путь, то возвращает его без изменений
func makeFullLink(link string, absolutepath string) (string, error) {
	isAbsolute, err := regexp.MatchString("^https{0,1}://", link)
	if err != nil {
		return "", err
	}

	if isAbsolute {
		return link, nil
	}

	if strings.HasPrefix(link, "/") {
		return fmt.Sprintf("%v%v", absolutepath, link), nil
	}

	return "", nil
}
