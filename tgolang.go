package main

import (
	"log"

	"strings"

	"golang.org/x/net/html"
)

type MGolang struct {
	supper Mbase
}

func (m MGolang) Supper() Mbase {
	return m.supper
}

var m = MGolang{supper: NewSupper("golang")}

func init() {

	ModuleRegisterCase(m, "HtmlParse")
	RegisterModule(m)
}
func (m MGolang) Htmlparse() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					log.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
