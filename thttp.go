package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"reflect"
	"time"
)

type MHttp struct {
	supper Mbase
}

func (m MHttp) Supper() Mbase {
	return m.supper
}

var vmod = MHttp{supper: NewSupper("http")}

func init() {

	ModuleRegisterCase(vmod, "Send")
	ModuleRegisterCase(vmod, "Receive")
	RegisterModule(vmod)
}
func (m MHttp) Client(s string) {

	timeout, err := time.ParseDuration(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("timeout:%s", timeout)
	c := http.Client{Timeout: timeout}
	resp, err := c.Get("http://192.168.176.3/status.html")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("%v", resp.Header)
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)

}
func (m MHttp) Rclient(s string) {

	timeout, err := time.ParseDuration(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("timeout:%s", timeout)
	r, err := http.NewRequest("GET", "http://dl.qihooyun.cn/status.html", nil)
	if err != nil {
		log.Fatal(err)
	}
	c := http.Client{Timeout: timeout}
	resp, err := c.Do(r)

	if err != nil {
		log.Printf("err %T %v\n", err, err)
		/*
			v := reflect.ValueOf(err)
			h := v.MethodByName("TimeOut")

			if h.IsValid() {
				log.Printf("timeout %v\n", h.Call(nil))
			}
		*/
		e := err.(*url.Error)
		log.Printf("Op:%v \n err:%T \n", e, e.Err)
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("%v", resp.Header)
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", buf)

}
