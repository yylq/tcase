package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	/*
	fmt.Print("tcompress init\n")
	RegisterCase("test_gzip", TCASE{"test_gzip", test_gzip})
	*/
}
func test_gzip(s string) {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create("tt.log.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gzw := gzip.NewWriter(f)

	defer gzw.Close()
	n, err := gzw.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("n:%d", n)
}
