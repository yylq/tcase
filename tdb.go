package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MDb struct {
	supper Mbase
}

func (m *MDb) Supper() Mbase {
	return m.supper
}

var mDb = newMode()

func newMode() Mode {
	return &MDb{supper: NewSupper("db")}
}

func init() {
	fmt.Print("tdb init\n")

	ModuleRegisterCase(mDb, "Query")

	RegisterModule(mDb)
}
func (m *MDb) Query() {
	db, err := sql.Open("mysql", "360cdn:123345@tcp(192.168.176.3:3306)/monitor?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from table_node_mon")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var bid string
		var ctime string
		var host string
		var dur int
		rows.Columns()
		err = rows.Scan(&id, &bid, &ctime, &host, &dur)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("bid:%v ctime:%v host:%v dur:%v\n", bid, ctime, host, dur)

	}

}
func (m *MDb) Trunc() {
	db, err := sql.Open("mysql", "360cdn:123345@tcp(192.168.176.3:3306)/monitor?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`truncate table table_node_mon`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	q, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(q)

}
func (m *MDb) Insert() {
	db, err := sql.Open("mysql", "360cdn:123345@tcp(192.168.176.3:3306)/monitor?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`insert into table_node_mon(bid,ctime,host,dur) values(?,now(),?,?)`)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		res, err := stmt.Exec("box_01", "192.168.176.6", 500000)
		if err != nil {
			log.Fatal(err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id:%v\n", id)
	}

}

func (m *MDb) Cal() {
	db, err := sql.Open("mysql", "360cdn:123345@tcp(192.168.176.3:3306)/monitor?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlstr := fmt.Sprintf(`select count(bid)/%d as cnum, sum(dur) as cav from table_node_mon where ctime>=? and host =? and bid=?`,
		5)
	fmt.Printf("sql:%s\n", sqlstr)
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query("2016-08-16 16:48:58", "192.168.176.6", "box_01")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var cav int32
		var cnum int32
		rows.Columns()
		err = rows.Scan(&cnum, &cav)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" cnum :%d ca:%v \n", cnum, cav)

	}

}
