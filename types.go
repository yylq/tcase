package main

import (
	"fmt"
)

type MTypes struct {
	supper Mbase
}

func (m MTypes) Supper() Mbase {
	return m.supper
}

var mtypes = MTypes{supper: NewSupper("types")}

func init() {

	ModuleRegisterCase(mtypes, "Testslice")
	ModuleRegisterCase(mtypes, "Byteop")
	RegisterModule(mtypes)
}

func (m MTypes) Testslice() {
	var pieces []int
	for i := 0; i < 10; i++ {
		pieces = append(pieces, i)

		fmt.Printf("%d arr len %d cap :%d\n", i, len(pieces), cap(pieces))

	}

	for i, v := range pieces {
		fmt.Printf("%d = %d\n", i, v)
	}
}
func (m MTypes) Byteop() {
	var a, b, c, d uint32

	a = 0x7FFFFFFF
	b = a << 1
	c = b << 1
	d = c << 1

	fmt.Printf("%x %x %x %x", a, b, c, d)
}
func test_map(s string) {
	type P struct {
		Id int
		Ns string
	}
	var ps = map[P]int{}
	var q1, q2 P
	q1.Id = 1
	q1.Ns = "q1"
	q2.Id = 2
	q2.Ns = "q2"
	ps[q1] = q1.Id
	ps[q2] = q2.Id
	for k, v := range ps {
		fmt.Println(k, v)
	}
	var q3 P
	q3.Id = 1
	q3.Ns = "q1"
	var q4 P
	q4.Id = 2
	fmt.Println(q1 == q3)
	fmt.Println(q1 == q2)
	i, ok := ps[q3]
	fmt.Println(i, ok)
	i, ok = ps[q4]
	fmt.Println(i, ok)

}
func test_array(s string) {
	paths := [][2]string{
		[2]string{"/host/", "base"},
		[2]string{"/host//conf", "conf"},
		[2]string{"/host//task", "conf"},
	}

	for i := 0; i < len(paths); i++ {
		fmt.Println(paths[i])
	}
}
