package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"myrpc"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"reflect"
	"time"
)

func init() {
	/*
	fmt.Print("tclient init\n")
	RegisterCase("test_mclient", TCASE{"test_mclient", test_mclient})

	RegisterCase("test_socket5", TCASE{"test_socket5", test_socket5})

	RegisterCase("test_rpc", TCASE{"test_rpc", test_rpc})
	*/

}
func test_mclient(s string) {
	var url = os.Args[2]
	var proxy = os.Args[3]
	os.Setenv("http_proxy", proxy)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.Header)
	var buf []byte
	//buf = make([]byte, 50)
	var b = bytes.NewBuffer(buf)
	log.Print(reflect.TypeOf(resp.Body))
	io.Copy(b, resp.Body)
	//b.ReadFrom(resp.Body)
	fmt.Printf("len:%d\n", len(buf))
	fmt.Print(b)

}

func test_socket5(s string) {
	var host string
	host = os.Args[2]

	conn, err := net.DialTimeout("tcp", host, 200*time.Millisecond)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//conn.SetDeadline(time.Now().Add(5 * time.Second))
	io.WriteString(conn, "hello world")
	///io.Copy(os.Stdout, conn)
	buf := make([]byte, 1024)
	//br := bytes.NewBuffer(buf)
	/*
		buf := make([]byte, 1024*5)
		b := bytes.NewBuffer(buf)

		n, err := b.ReadFrom(conn)
		if err != nil {
			fmt.Print(err)
			return
		}*/
	i := 0
	for {
		n, err := conn.Read(buf[i:])
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		i += n
		fmt.Printf("i:%d n:%d\n", i, n)
	}
	//io.Copy(br, conn)
	fmt.Printf("\n%s\n ", buf)
	/*
		buf := make([]byte, 50)
		for {
			//var buf []byte

			n, err := conn.Read(buf)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Printf("n=%d \n buff:%s\n", n, buf)
			buf = make([]byte, 50)
		}
	*/
	//fmt.Printf("len=%d \n buff:%s\n", n, buf)

	/*
		var buf []byte
		//buf = make([]byte, 50)
		var b = bytes.NewBuffer(buf)

		n, err := b.ReadFrom(conn)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("len=%d \n buff:", n, buf)
	*/
	os.Geteuid()

}
func run_rpcarith(client *rpc.Client) {

	args := &myrpc.Args{7, 8}
	var reply int
	err := client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	args1 := &myrpc.Args{100, 8}
	quotient := new(myrpc.Quotient)
	err = client.Call("Arith.Divide", args1, quotient)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Divide: %d*%d=%d %d \n", args1.A, args1.B, quotient.Quo, quotient.Rem)

}
func run_argint(client *rpc.Client) {

	arg := &myrpc.Sarg{Na: "aaaa", Id: 10}
	var reply int
	err := client.Call("MSeed.Add", arg, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	arg1 := &myrpc.Sarg{Na: "aaaa"}
	err = client.Call("MSeed.Get", arg1, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("mseed: [%d]=%d reply :%d\n", arg1.Na, arg1.Id, reply)
}
func run_argstring(client *rpc.Client) {

	var reply int
	arg := &myrpc.Sarg{Na: "aaaa", Id: 10, Va: "bbbbbbb"}
	fmt.Printf("mseed: key:%s value:%s\n", arg.Na, arg.Va)

	arg.List = make([]int, 5)
	for i := 0; i < len(arg.List); i++ {
		arg.List[i] = i
	}
	fmt.Printf("len:%d\n", len(arg.List))

	err := client.Call("MSeed.AddString", arg, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	/*
		var arg1 string;

		err = client.Call("MSeed.GetString",arg.Na, &arg1)
		if err != nil {
			log.Fatal("arith error:", err)
		}

		fmt.Printf("mseed: key:%s value:%s\n",arg.Na,arg1)
	*/
	//sed := new(myrpc.Args)
	var sed myrpc.Sarg
	err = client.Call("MSeed.GetSeed", "aaaa", &sed)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(sed)
}
func test_rpc(s string) {
	var host = os.Args[2]

	client, err := rpc.DialHTTP("tcp", host)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()
	run_rpcarith(client)
	run_argstring(client)
}
