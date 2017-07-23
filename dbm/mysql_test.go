package dbm

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestMysql(t *testing.T) {
	var err error
	_, err = Default("user:123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		t.Error(err)
		return
	}
}
