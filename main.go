//+build main
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var mcase string

func init() {
	flag.StringVar(&mcase, "mcase", "", "-mcase=mode.fname [arg1] [args2]...")

}
func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	if mcase == "" {
		log.Fatal("case name is empty")

	}
	s := strings.Split(mcase, ".")
	if len(s) < 2 {
		log.Fatal("invalid input: -mcase=mode.fname")
		return
	}
	mname := s[0]
	M := []byte(s[1])
	fname := strings.ToUpper(string(M[0])) + string(M[1:])
	fmt.Printf("mname:%s cname:%s %v\n", mname, fname, flag.Args())
	m, err := CheckVaildCase(mname, fname)
	if err != nil {
		log.Fatal("%s\n", err)
		return
	}
	RunCase(m, fname, flag.Args())
}
