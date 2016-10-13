package main

import (
	"fmt"
	"log"

	"github.com/yylq/config"
)

type MConf struct {
	supper Mbase
}

func (m MConf) Supper() Mbase {
	return m.supper
}

var mconf = newConf()

func newConf() Mode {
	return MConf{supper: NewSupper("conf")}
}

func init() {
	fmt.Print("tconf init\n")

	ModuleRegisterCase(mconf, "First")

	RegisterModule(mconf)
}
func (m MConf) One() {
	configFile := "./test.ini"
	cfg, err := config.ReadDefault(configFile)
	if err != nil {
		log.Fatalf("Fail to find:%s err:%v", configFile, err)
	}
	//set config file std End
	var TOPIC = make(map[string]string)
	//Initialized topic from the configuration
	if cfg.HasSection("main") {
		section, err := cfg.SectionOptions("main")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("main", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	//Initialized topic from the configuration END

	fmt.Print(TOPIC)
	fmt.Println(TOPIC["debug"])
}
func (m MConf) Tsec() {
	type Q struct {
		Addr  string `config:"main-addr"`
		Debug bool   `config:"main-debug"`
		T1    int32  `config:"other-t1"`
	}
	configFile := "test.ini"
	cfg, err := config.ReadDefault(configFile)
	if err != nil {
		log.Fatalf("Fail to find:%s err:%v", configFile, err)
	}
	var q Q
	err = cfg.ParseConf(&q)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(q)
}
func (m MConf) Two() {
	type Q struct {
		Runserv   []string  `config:"main-runserv"`
		Runflag   []bool    `config:"main-runflag"`
		Runstatus []int     `config:"main-runstatus"`
		Runrate   []float64 `config:"main-runrate"`
	}
	configFile := "test.ini"
	cfg, err := config.ReadDefault(configFile)
	if err != nil {
		log.Fatalf("Fail to find:%s err:%v", configFile, err)
	}
	var q Q
	err = cfg.ParseConf(&q)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(q)
}
