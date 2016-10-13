package main

import (

	//"os"
	//"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	//"os"
	//"math"

	"time"
)

type MGob struct {
	supper Mbase
}

func (m MGob) Supper() Mbase {
	return m.supper
}

var mod = MGob{supper: NewSupper("gob")}

func init() {

	ModuleRegisterCase(mod, "Testgob")
	ModuleRegisterCase(mod, "Receive")
	RegisterModule(mod)
}

type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y *int32
	Name string
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func (m MGob) Testgob(s string) {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) some values.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) and print the values.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	// Output:
	// "Pythagoras": {3, 4}
	// "Treehouse": {1782, 1841}
}

func (m MGob) Tjson() {
	type Cmd struct {
		Tid   string `json:"tid"`
		state int
	}
	var c Cmd
	c.Tid = "aaaaaa"
	c.state = 1
	body, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", body)
	type CmdRes struct {
		int `json:"stat"`
	}
	var a CmdRes
	a = CmdRes{10}
	fmt.Print(a)
	body, err = json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", body)
}
func (m MGob) Tcomjson() {
	/*
		type state struct {
			Hash           string `json:"hash"`
			Total_wanted   int64  `json:"total_wanted"`
			Total_download int64  `json:"total_download"`
			Total_upload   int64  `json:"total_upload"`
		}
		type Suminfo struct {
			Host string  `json:"host"`
			Info []state `json:"info"`
		}
		hashs := []string{"73694de7fc2149770d626cbd011947adbc49c02f",
			"343734273286bfdb03846baed9de239fd96a3c4e",
			"fd44a542d717b955604ed4f327eea9095b7516c6"}

		states := make([]state, 3)
		for i := 0; i < 3; i++ {
			states[i].Hash = hashs[i]
			states[i].Total_wanted = (1024 * int64(math.Pow10(i+1)))
		}
		fmt.Print(states)
		var info Suminfo
		info.Host = "192.168.145.225"
		info.Info = states

		body, err := json.Marshal(info)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", body)
		fmt.Println("\n*******************************\n")
		var ts Suminfo
		err = json.Unmarshal(body, &ts)
		fmt.Print(ts)
		ls := ts.Info
		fmt.Print(ls)
	*/
	//enc := json.NewEncoder(os.Stdout)
	//enc.Encode(states)
	/*
		stteas[0].hash =

		 := state{"", 1024, 256, 0}
		states[1] := state{"8e35b08b3be69d4d792dd31ece12493b4a52b18e", 10240, 2560, 0}
		states[2] := state{"", 102400, 25600, 0}
		body, err := json.Marshal(states[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", body)
	*/
	ms := map[string]string{
		"hello": "world",
		"aaaa":  "1111",
		"bbbb":  "22222",
	}
	body, err := json.Marshal(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", body)

}
func (m MGob) Selfjson() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
		{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
	]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		log.Printf("%s\n", c.Point)
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}

}

type UmMsgReport struct {
	Tid   string          `json:"tid"`
	Bid   string          `json:"bid"`
	Aid   string          `json:"aid"`
	Ctime int             `json:"ctime"`
	Data  json.RawMessage `json:"data"`
}

type MsgReport struct {
	Tid   string      `json:"tid"`
	Bid   string      `json:"bid"`
	Aid   string      `json:"aid"`
	Ctime int         `json:"ctime"`
	Data  interface{} `json:"data"`
	data  json.RawMessage
}
type DownRinfo struct {
	Hid    string `json:"hid"`
	Status int32  `json:"st"`
	Dur    int32  `json:"dur"`
}

func (m *MsgReport) String() string {
	var p interface{}
	if m.Data != nil {
		p = *m
	} else {
		p = UmMsgReport{Tid: m.Tid, Bid: m.Bid, Data: m.data}
	}
	buff, _ := json.Marshal(p)
	return string(buff)
}
func (m *MsgReport) UnmarshalJSON(buff []byte) error {
	var um UmMsgReport

	if err := json.Unmarshal(buff, &um); err != nil {
		return err
	}
	log.Printf("%s\n", um.Data)
	m.Tid = um.Tid
	m.Bid = um.Bid
	m.Aid = um.Aid
	m.Ctime = um.Ctime
	m.data = um.Data

	return nil
}
func (m MGob) Selfum() {
	hinfo := []DownRinfo{
		DownRinfo{Hid: "h1", Status: 1, Dur: 100},
		DownRinfo{Hid: "h2", Status: 1, Dur: 100},
	}
	fmt.Print(hinfo)
	data, _ := json.Marshal(hinfo)
	fmt.Printf("data:%s\n", data)
	mrt := MsgReport{Tid: "111", Bid: "222", Aid: "w10", Ctime: int(time.Now().Unix()), Data: hinfo}
	log.Printf("%s\n", mrt.String())

	buff, err := json.Marshal(mrt)
	fmt.Printf("%s\n", buff)
	if err != nil {
		log.Fatal(err)
	}
	var um MsgReport
	if err = json.Unmarshal(buff, &um); err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", um.data)
	log.Printf("%s\n", um.String())

}

type RawString string

func (m *RawString) MarshalJSON() ([]byte, error) {
	return []byte(*m), nil
}
func (m *RawString) UnmarshalJSON(buff []byte) error {
	*m = RawString(buff)
	return nil
}

type Rawst struct {
	Dur  []byte    `json:"dur"`
	Data RawString `json:"data"`
}

func (m MGob) Rawjson() {
	var j = []byte(`[
		{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
	]`)
	s := Rawst{Data: RawString(j), Dur: j}
	buff, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", buff)

}
func (m MGob) One() {

	t := []bool{true, false, true}
	buff, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("buff:%s", buff)
	var ut []bool
	err = json.Unmarshal(buff, &ut)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("buff:%v", ut)
}
