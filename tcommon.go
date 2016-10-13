//+build main
package main

import (
	"strconv"

	//"strings"
	"fmt"
	"log"
	"myhttp"
	"net/http"
	"os"
	"time"
)

type MCommon struct {
	supper Mbase
}

func (m MCommon) Supper() Mbase {
	return m.supper
}

var mcommon = newMCommon()

func newMCommon() Mode {
	return MCommon{supper: NewSupper("common")}
}
func init() {

	ModuleRegisterCase(mcommon, "Hello")
	RegisterModule(mcommon)
}

func (m MCommon) Hello(args string) {

	fmt.Printf("hello %s\n", args)
}
func (m MCommon) Httpget(url string) {

	fmt.Println("url:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Write(os.Stdout)
}
func (m MCommon) Mhttp(url string, proxy string) {

	client := myhttp.NewClient()
	resp, err := client.Get(url, proxy, "nil")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
func (m MCommon) Testlog(s string) {
	file, err := os.OpenFile("./tt.log", os.O_CREATE|os.O_WRONLY, 0x666)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetPrefix("testlog")
	log.SetOutput(file)
	log.Printf(s)
}

func (m MCommon) Time() {
	const TimeFormat = "2006-01-02 15:04:05"
	cur := time.Now()
	fmt.Printf("%v\n %s\n", cur, cur.Format(TimeFormat))
	fmt.Printf("%d\n ", cur.Unix())
	bf := time.Unix(cur.Unix()-300, 0)
	fmt.Printf("%v\n %s\n", bf, bf.Format(TimeFormat))

}
func (m MCommon) Ttimer(timeout string, count string) {
	d, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatal(err)
	}
	c, err := strconv.Atoi(count)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("dru:%v c:%v\n", d, c)
	t1 := time.NewTimer(d)
	exit := false
	i := 0
	for !exit {
		select {

		case <-t1.C:
			i += 1
			fmt.Printf("timeout cout :%d\n", i)
			if i < c {

				t1.Reset(d)
				continue
			} else {
				exit = true
			}

		}
	}
}
