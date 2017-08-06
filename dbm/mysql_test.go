package dbm

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tangs-drm/go-tool/log"
	"github.com/tangs-drm/go-tool/util"
	"fmt"
)

func TestMysql(t *testing.T) {
	var err error
	var db *DB
	db, err = Default("testUser:123@tcp(localhost:3306)/testdb?charset=utf8")
	if err != nil {
		t.Error(err)
		return
	}
	err = db.Ping()
	if err != nil {
		t.Error(err)
		return
	}

	res, err := db.Exec("drop table if exists test;")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
	log.Debug("res1 -- %v", util.S2Json(res))

	res, err = db.Exec("create table test(id int, val varchar(16));")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)
	log.Debug("res2 -- %v", util.S2Json(res))

	// 手动注册数据库
	db, err = NewDB("", "")
	if err == nil {
		t.Error(err)
		return
	}

	db = DefaultDBManger.Db()
	if db == nil {
		t.Error(err)
		return
	}

	db = DefaultDBManger.Db("default")
	if db == nil {
		t.Error(nil)
		return
	}

	db, err = NewDBWithManager(DefaultDBManger, "", "testUser:123@tcp(localhost:3306)/testdb?charset=utf8")
	if err == nil {
		t.Error(err)
		return
	}

	db, err = NewDBWithManager(DefaultDBManger, "halo", "testUser:123@tcp(localhost:3306)/testdb?charset=utf8")
	if err != nil {
		t.Error(err)
		return
	}

	err = db.Ping()
	if err != nil {

		t.Error(err)
		return
	}
}
