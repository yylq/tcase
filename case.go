package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type Funcptr func(s string)

type Mode interface {
	Supper() Mbase
}

type Mbase struct {
	Name  string
	Cases map[string]bool
}

func (m Mbase) GetName() string {
	return m.Name
}

func (m Mbase) AddCase(name string) error {
	m.Cases[name] = true
	return nil
}

func (m Mbase) ShowModCase() {
	fmt.Printf("module:%s\n cases:", m.Name)
	for k, _ := range m.Cases {
		fmt.Printf("\t:%s\n", k)
	}
}

func NewSupper(mname string) Mbase {
	return Mbase{Name: mname, Cases: make(map[string]bool)}
}

var modules = make(map[string]Mode)

func CheckVaildCase(mname string, cname string) (Mode, error) {
	m, ok := modules[mname]
	if !ok {
		return nil, errors.New(mname + " is no exist")
	}
	v := reflect.ValueOf(m)
	h := v.MethodByName(cname)

	if !h.IsValid() {
		m.Supper().ShowModCase()
		return nil, errors.New(cname + " is no exist")
	}

	return m, nil

}
func RunCase(m Mode, cname string, args []string) error {
	v := reflect.ValueOf(m)
	h := v.MethodByName(cname)
	if !h.IsValid() {
		return errors.New(cname + "is not exist")
	}
	var in []reflect.Value
	if len(args) != 0 {
		in = make([]reflect.Value, len(args))
		for i := 0; i < len(args); i++ {
			in[i] = reflect.ValueOf(args[i])
		}
	}
	h.Call(in)
	return nil
}
func ShowModCase() {
	for _, v := range modules {
		v.Supper().ShowModCase()
	}
}
func RegisterModule(m Mode) {
	name := m.Supper().GetName()
	modules[name] = m
}
func ModuleRegisterCase(m Mode, cname string) {
	m.Supper().AddCase(cname)
}
func LocateCase() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "?()"
	}
	fn := runtime.FuncForPC(pc)

	if fn == nil {
		return "?()"
	}
	dotName := filepath.Ext(fn.Name())
	return strings.TrimLeft(dotName, ".") + "()"
}
