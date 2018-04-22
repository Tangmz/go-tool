package database

import (
	"testing"
	"github.com/tangs-drm/go-tool/log"
	"github.com/tangs-drm/go-tool/util"
	"fmt"
)

func TestMysql(t *testing.T) {
	log.Debug("TestMysql --- START")
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

	db = DefaultDBManger.D()
	if db == nil {
		t.Error(err)
		return
	}

	db = DefaultDBManger.D("default")
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

	// 测试执行多条sql语句
	var sqlString string = `/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2017/8/6 14:47:07                            */
/*==============================================================*/


DROP INDEX LASTTIME ON SESSION;

DROP INDEX TIME ON SESSION;

DROP INDEX UID ON SESSION;

DROP INDEX ID ON SESSION;

DROP TABLE IF EXISTS SESSION;

/*==============================================================*/
/* Table: SESSION                                               */
/*==============================================================*/
CREATE TABLE SESSION
(
   ID                   VARCHAR(64) NOT NULL,
   UID                  VARCHAR(64),
   TIME                 INT,
   LASTTIME             INT,
   PRIMARY KEY (ID)
);

/*==============================================================*/
/* Index: ID                                                    */
/*==============================================================*/
CREATE INDEX ID ON SESSION
(
   ID
);

/*==============================================================*/
/* Index: UID                                                   */
/*==============================================================*/
CREATE INDEX UID ON SESSION
(
   UID
);

/*==============================================================*/
/* Index: TIME                                                  */
/*==============================================================*/
CREATE INDEX TIME ON SESSION
(
   TIME
);

/*==============================================================*/
/* Index: LASTTIME                                              */
/*==============================================================*/
CREATE INDEX LASTTIME ON SESSION
(
   LASTTIME
);

`
	db = D()
	if db == nil {
		t.Error(err)
		return
	}

	err = db.ExecScripts(sqlString, false)
	if err == nil {
		t.Error(err)
		return
	}

	err = db.ExecScripts(sqlString, true)
	if err != nil {
		t.Error(err)
		return
	}

	// check db
	rows, err := db.Query("desc SESSION;")
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()

	var count int = 0
	for rows.Next() {
		count ++
	}
	if count != 4 {
		t.Error(count)
		return
	}

	log.Debug("TestMysql --- END")
}
