package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	//"unicode/utf8"

	"bytes"

	"github.com/opesun/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type MHtml struct {
	supper Mbase
}

func (m MHtml) Supper() Mbase {
	return m.supper
}

var mhtml = newMHtml()

func newMHtml() Mode {
	return MHtml{supper: NewSupper("html")}
}

func init() {
	fmt.Print("thtml init\n")

	ModuleRegisterCase(mhtml, "Htmlstring")
	ModuleRegisterCase(mhtml, "Htmltemplate")
	ModuleRegisterCase(mhtml, "Urlparse")
	ModuleRegisterCase(mhtml, "Downfile")
	ModuleRegisterCase(mhtml, "Parsecontent")
	ModuleRegisterCase(mhtml, "Parseutf")
	RegisterModule(mhtml)
}
func (m MHtml) Htmlstring() {
	var example = `
<html>
	<head>
		<title>
		</title>
	</head>
<body>
	<div class=hey cust-a="wow"><h2>Title here</h2></div>
	<span><h2>Yoyoyo</h2></span>
	<div id="box">
		<span>
			content<a href=""><div><li></li></div></a>
		</span>
	</div>
	<div class="yo hey">
		<a href="xyz"><div class="cow sheep bunny"><h8>content</h8></div></a>
	</div>
</body>
</html>
`
	tags := []string{"div#box", "div.cow", "[cust-a=wow]"}

	x, err := goquery.ParseString(example)
	if err != nil {
		log.Fatal(err)
	}

	for _, tag := range tags {
		log.Printf("--%s----------------------------\n", tag)
		n := x.Find(tag)
		log.Printf("len:%d \n", len(n))
		if len(n) <= 0 {
			log.Println("not find\n")
			continue
		}
		n.Print()
	}
}

type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

type OnlineUser struct {
	User      []*Person
	LoginTime string
}

func (m MHtml) Htmltemplate() {

	dumx := Person{
		Name:    "zoro",
		Age:     27,
		Emails:  []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role:    "SE"}

	chxd := Person{Name: "chxd", Age: 27, Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, onlineUser)
	if err != nil {
		log.Fatal(err)
	}

}
func (m MHtml) Urlparse(s string) {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(u.Host)
	log.Print(u.RequestURI())

}

type HtmlA struct {
	Href  string
	Title string
}
type HtmlAS struct {
	Nodes []*HtmlA
}

func (m MHtml) save_teml_content(ofile string, tepl string, data interface{}) error {

	ind, err := os.OpenFile(ofile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0x666)
	if err != nil {
		log.Print(err)
		return err
	}
	t, err := template.ParseFiles(tepl)
	if err != nil {
		log.Print(err)
		return err
	}

	err = t.Execute(ind, data)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
func (m MHtml) BiqugeIndex() {
	/*
		ind, err := os.OpenFile("index.html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0x666)
		if err != nil {
			fmt.Println(err)
			return
		}*/
	dumx := HtmlA{
		Href: "562989.html", Title: "第一章 寒门子弟"}

	chxd := HtmlA{
		Href: "562990.html", Title: "第二章 危机"}

	onlineUser := HtmlAS{Nodes: []*HtmlA{&dumx, &chxd}}
	/*
		t, err := template.ParseFiles("bqg_ind_teml.html")
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(ind, onlineUser)
	*/
	err := m.save_teml_content("index.html", "bqg_ind_teml.html", onlineUser)
	if err != nil {
		log.Fatal(err)
	}

}
func (m MHtml) Downfile(http_url string, store_file string) {
	res, err := http.Get(http_url)
	if err != nil {
		log.Fatal(err)

	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(store_file, os.O_CREATE|os.O_WRONLY, 0x666)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(robots)
	file.Close()
}

type Content struct {
	Data string
}

func Encode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	if err == nil {
		dst = string(data)
	}
	return
}
func Decode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}
