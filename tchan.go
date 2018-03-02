package main

import (
	"fmt"
	"sync"
	"time"
)

type MChan struct {
	supper Mbase
}

func (m MChan) Supper() Mbase {
	return m.supper
}

var mchan = MChan{supper: NewSupper("chan")}

func init() {

	ModuleRegisterCase(mchan, "Writechan")
	RegisterModule(mchan)
}

func (m MChan) WRchan() {
	var str string
	wc := make(chan string)
	go func() {

		for val := range wc {

			fmt.Printf("childd  %s \n", val)
			wc <- fmt.Sprintf(" childd %s", val)
		}
	}()

	for true {
		fmt.Scanln(&str)

		wc <- str

		fmt.Printf("father %s \n", <-wc)
	}

}
func (m MChan) Wchan() {
	var str string
	wc := make(chan string)
	go func(a <-chan string) {

		for val := range a {

			fmt.Printf("childd  %s \n", val)

		}
	}(wc)

	for true {
		fmt.Scanln(&str)

		wc <- str

	}

}

/*
var ch1 chan<- int  //ch1只能写

var ch2 <-chan int  //ch2只能读
*/
type Tt struct {
	id  int
	str string
}

func (t Tt) Cmd() string {
	return t.str
}

func ReadChan(a chan Tt) <-chan Tt {
	return a
}
func Writechan(a chan Tt) chan<- Tt {
	return a
}
func (m MChan) Comchan() {
	var str string
	wc := make(chan Tt)

	go func(a chan Tt) {
		rd := ReadChan(a)
		for val := range rd {

			fmt.Printf("childd  %v \n", val)

		}
	}(wc)

	id := 0
	wrc := Writechan(wc)
	wrq := Writechan(wc)
	for true {
		fmt.Scanln(&str)
		t := Tt{id, str}
		if str == "quit" {
			wrq <- t
			break
		}

		wrc <- t
		id++

	}

}

type cmd interface {
	Cmd() string
}

func ReadCmd(a chan cmd) <-chan cmd {
	return a
}
func WriteCmd(a chan cmd) chan<- cmd {
	return a
}
func (m MChan) Mulchan() {
	var str string
	wc := make(chan cmd)
	wq := make(chan cmd)
	go func() {
		rc := ReadCmd(wc)
		rq := ReadCmd(wq)
		var val cmd
		for true {
			select {
			case val = <-rc:
				fmt.Printf("child cmd  %v \n", val)
			case val = <-rq:
				fmt.Printf("child quit  %v \n", val)
				break

			}
		}
		fmt.Print("func exit")
	}()
	id := 0
	wr := WriteCmd(wc)
	wrq := WriteCmd(wq)
	for true {
		fmt.Scanln(&str)
		t := Tt{id, str}
		if str == "quit" {
			wrq <- t
			break
		}

		wr <- t
		id++
	}
	time.Sleep(5)
}
func (m MChan) Closechan() {
	var str string
	wc := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a <-chan string) {
		defer wg.Done()

		for true {
			val, isOK := <-a
			fmt.Printf("childd  %s %v\n", val, isOK)
			if !isOK {
				break
			}

		}
		fmt.Println("go func exit")
	}(wc)

	for i := 0; i < 5; i++ {
		fmt.Scanln(&str)

		wc <- str

	}
	close(wc)
	wg.Wait()
	fmt.Println("main exit")
}
