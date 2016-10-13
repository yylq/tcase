package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

type MRedis struct {
	supper Mbase
}

func (m MRedis) Supper() Mbase {
	return m.supper
}

var mredis = newMRedis()

func newMRedis() Mode {
	return MRedis{supper: NewSupper("redis")}
}

func init() {
	fmt.Print("tredis init\n")

	ModuleRegisterCase(mredis, "Dial")
	ModuleRegisterCase(mredis, "Get")
	ModuleRegisterCase(mredis, "Hgetall")
	ModuleRegisterCase(mredis, "Hmget")

	RegisterModule(mredis)
}
func (m MRedis) Dial(s string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(2))
	if err != nil {
		panic(err)
	}
	defer c.Close()

}
func (m MRedis) Get(s string) {
	c, err := redis.Dial("tcp", "192.168.176.3:6379", redis.DialDatabase(2))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	ok, err := redis.String(c.Do("GET", "service_role_code"))
	if err != nil {
		panic(err)
	}
	fmt.Print(ok)
}
func (m MRedis) Set() {
	c, err := redis.Dial("tcp", "192.168.176.3:6379", redis.DialDatabase(3))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	ok, err := redis.String(c.Do("SET", "testkey", "test"))
	if err != nil {
		panic(err)
	}
	fmt.Print(ok)
}
func (m MRedis) Hgetall(s string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(2))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	ok, err := redis.StringMap(c.Do("HGETALL", "*.pplive.com"))
	if err != nil {
		panic(err)
	}
	for k, v := range ok {
		fmt.Printf("%s = %s\n", k, v)
	}

}
func (m MRedis) Hmget(s string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(2))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	ok, err := redis.Strings(c.Do("HMGET", "w.pplive.com", "cache_key_without_args", "sorted_src_ip_usability_list"))
	if err != nil {
		panic(err)
	}
	for k, v := range ok {
		fmt.Printf("%d = %s\n", k, v)
	}
}

func (m MRedis) pub(ch string, msg string) {

	c, err := redis.Dial("tcp", "192.168.176.3:6379", redis.DialReadTimeout(100*time.Second),
		redis.DialWriteTimeout(100*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	log.Printf("ch:%s msg:%s\n", ch, msg)
	ok, err := redis.Int64(c.Do("PUBLISH", ch, msg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(ok)
	/*
		var wg sync.WaitGroup
		wg.Add(2)
		go func()
		 {
			defer wg.Done()
			for {
				switch n := psc.Receive().(type) {
				case redis.Message:
					fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
				case redis.PMessage:
					fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
				case redis.Subscription:
					fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
					if n.Count == 0 {
						return
					}
				case error:
					fmt.Printf("error: %v\n", n)
					return
				}
			}
		}()
		go func() {
			defer wg.Done()

			psc.Subscribe("example")
			psc.PSubscribe("p*")

			// The following function calls publish a message using another
			// connection to the Redis server.
			publish("example", "hello")
			publish("example", "world")
			publish("pexample", "foo")
			publish("pexample", "bar")

			// Unsubscribe from all connections. This will cause the receiving
			// goroutine to exit.
			psc.Unsubscribe()
			psc.PUnsubscribe()
		}()
		wg.Wait()
		return
	*/
}
func (m MRedis) Sub(ch string) {

	c, err := redis.Dial("tcp", "192.168.176.3:6379", redis.DialConnectTimeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	psc := redis.PubSubConn{Conn: c}
	if err = psc.Subscribe(ch); err != nil {
		log.Fatal(err)
	}
	go func() {
		defer wg.Done()
		for {
			switch n := psc.Receive().(type) {
			case redis.Message:
				fmt.Printf("Message: %s %s\n", n.Channel, n.Data)
			case redis.PMessage:
				fmt.Printf("PMessage: %s %s %s\n", n.Pattern, n.Channel, n.Data)
			case redis.Subscription:
				fmt.Printf("Subscription: %s %s %d\n", n.Kind, n.Channel, n.Count)
				if n.Count == 0 {
					return
				}
			case error:
				fmt.Printf("error: %v\n", n)
				return
			}
		}
	}()
	wg.Wait()

}
