package dbm

import (
	"database/sql"
	"github.com/tangs-drm/go-tool/util"
	"strings"
)

type DB struct  {
	*sql.DB
}

var Db *DB

// NewDB根据name新建一个数据库的对象
func NewDB(name string, url string) (*DB, error) {
	if len(name) < 1 {
		return nil, util.Error("name is invalid: %v", name)
	}
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	var db_ = &DB{DB: db}
	DefaultDBManger.Register(name, db_)
	return db_, nil
}

// Default使用默认的数据库对象
func Default(url string) (*DB, error) {
	var err error
	Db, err =  NewDB("default", url)
	if err != nil {
		return nil, err
	}
	return Db, nil
}

var DefaultDBManger = NewDbManager()

type DbManager struct {
	DbMap map[string]*DB
}

// NewDbManager 新建一个数据库对象管理者
func NewDbManager() *DbManager {
	return &DbManager{
		DbMap:map[string]*DB{},
	}
}

// Register 通过name注册数据库对象
func (dm *DbManager) Register(name string, db *DB) {
	dm.DbMap[name] = db
}

// Db 根据name返回对应的数据库对象,name为空则返回默认的数据库对象
func (dm *DbManager) D(name... string) *DB {
	if len(name) < 1 {
		return dm.DbMap["default"]
	}
	nm := name[0]
	return dm.DbMap[nm]
}

func D(name... string) *DB {
	if len(name) < 1 {
		return DefaultDBManger.DbMap["default"]
	}
	nm := name[0]
	return DefaultDBManger.DbMap[nm]
}

// NewDBWithManager 新建数据库对象,并注册到指定的数据库管理者里
func NewDBWithManager(manager *DbManager, name string, url string) (*DB, error) {
	if len(name) < 1 {
		return nil, util.Error("name is invalid: %v", name)
	}
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	var db_ = &DB{DB: db}
	manager.Register(name, db_)
	return db_, nil
}

// ExecScripts 处理了多条sql命令的执行
// 参数1: 需要执行的sql语句
// 参数2: 如果发生错误是否继续执行下去, false: 不继续执行, true: 继续执行
func (db *DB) ExecScripts(scripts string, Continue bool) error {
	commands := strings.Split(scripts, ";")
	if len(commands) < 1 {
		return util.Error("scripts get command empty")
	}

	for _, command := range commands {
		command = strings.TrimSpace(command)
		_, err := db.Exec(command)
		if err != nil && !Continue {
			return err
		}
	}
	return nil
}