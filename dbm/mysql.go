package dbm

import (
	"database/sql"
	"github.com/tangs-drm/go-tool/util"
)

type DB struct  {
	*sql.DB
}

var Db *DB

func NewDB(name string, url string) (*DB, error) {
	if len(name) < 1 {
		return nil, util.Error("name is invalid: %v", name)
	}
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	var db_ = &DB{DB: db}
	DbMangerV.Register(name, db_)
	return db_, nil
}

func Default(url string) (*DB, error) {
	var err error
	Db, err =  NewDB("default", url)
	if err != nil {
		return nil, err
	}
	return Db, nil
}

var DbMangerV = NewDbManager()

type DbManager struct {
	DbMap map[string]*DB
}

func NewDbManager() *DbManager {
	return &DbManager{
		DbMap:map[string]*DB{},
	}
}

func (dm *DbManager) Register(name string, db *DB) {
	dm.DbMap[name] = db
}

func (dm *DbManager) Db(name... string) *DB {
	if len(name) < 1 {
		return dm.DbMap["default"]
	}
	nm := name[0]
	return dm.DbMap[nm]
}