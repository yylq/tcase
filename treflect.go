package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type MReflect struct {
	supper Mbase
}

func (m MReflect) Supper() Mbase {
	return m.supper
}

var mreflect = MReflect{supper: NewSupper("reflect")}

func init() {
	fmt.Print("reflect init\n")
	/*
		RegisterCase("test_reflect", TCASE{"test_reflect", test_reflect})
		RegisterCase("test_attr", TCASE{"test_attr", test_attr})
		RegisterCase("test_rand", TCASE{"test_rand", test_rand})
		RegisterCase("test_runtime", TCASE{"test_runtime", test_runtime})
	*/
	RegisterModule(mreflect)
}

var ErrNotFound = errors.New("not found")
var ErrUnsupportedType = errors.New("unsupported type")

func fieldName(f reflect.StructField) string {
	if f.Anonymous {
		return ""
	}
	tag := f.Tag.Get("config")
	fmt.Printf("tag:%s\n", tag)
	if tag != "" {
		if tag == "-" {
			return ""
		}
		tagParts := strings.Split(tag, ",")
		if len(tagParts) >= 1 {
			return strings.TrimSpace(tagParts[0])
		}
	}
	return f.Name
}
func loadStruct(v reflect.Value) error {
	t := v.Type()
	n := t.NumField()
	for i := 0; i < n; i++ {
		f := t.Field(i)

		fmt.Printf("name:%s tag:%s %s\n", f.Name,
			fieldName(f), f.Type)

	}
	return nil
}
func Load(st interface{}) error {
	v := reflect.ValueOf(st)
	k := v.Kind()
	if k != reflect.Ptr && k != reflect.Interface {
		return ErrUnsupportedType
	} else if v.IsNil() {
		return ErrUnsupportedType
	}
	e := v.Elem()

	loadStruct(e)
	return nil

}
func (m MReflect) Tfield() {
	type Q struct {
		Bind     string `config:"bind"`
		Redisip  string `config:"redisip"`
		Redisdb  int    `config:"redisdb"`
		RabbitMq string `config:"rabbit"`
		Mysql    string `config:"mysql"`
		Logpath  string `config:"logpath"`
		Regurl   string `config:"regurl"`
		Binpath  string `config:"binpath"`
	}
	var q Q
	err := Load(&q)
	if err != nil {
		log.Fatal(err)
	}
	var s []string
	v := reflect.ValueOf(s)
	k := v.Kind()
	log.Printf("%v", k)
	log.Printf("%v", v.Type())
}
