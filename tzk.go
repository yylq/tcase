package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func init() {
	fmt.Print("zk_test init\n")
	/*
	RegisterCase("test_watchchild", TCASE{"test_watchchild", test_watchchild})
	RegisterCase("test_zkcreatetemp", TCASE{"test_create", test_zkcreatetemp})
	*/
}
func test_watchchild(s string) {
	c, _, err := zk.Connect([]string{"192.168.145.224"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	for {
		children, _, ch, err := c.ChildrenW("/zk_test")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", children)
		<-ch
		//fmt.Printf("E:%+v\n", e)
	}
}
func test_createtemp(s string) {
	//need use other way ceate temp node such as  zkCli.sh
	// cmd: create -e /zk_test/005 ""
	c, _, err := zk.Connect([]string{"192.168.145.224"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	for {
		ok, _, ch, err := c.ExistsW("/zk_test/005")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", ok)
		<-ch
	}
}
func test_zkcreatetemp(s string) {
	c, _, err := zk.Connect([]string{"192.168.145.224"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	defer c.Close()
}
