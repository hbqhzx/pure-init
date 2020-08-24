package db

import (
	"fmt"
	"os"
	"pure-init/lib/obj"

	"github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/olebedev/config"
)

const (
	DBSingularTable = true
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type MyDB struct {
	*gorm.DB
}

var DB MyDB

type Logger struct {
}

func (logger Logger) Print(v ...interface{}) {
	seelog.Trace(gorm.LogFormatter(v...)...)
}

func InitDBWithConfig(cfg *config.Config) error {
	dataSourceName := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?%s",
		cfg.UString("username"),
		cfg.UString("password"),
		cfg.UString("protocol"),
		cfg.UString("address"),
		cfg.UString("db"),
		cfg.UString("params"),
	)

	var db *gorm.DB
	if err := obj.Retry("init db", 10, func() (e error) {
		db, e = gorm.Open("mysql", dataSourceName)
		return e
	}); err != nil {
		fmt.Println(err.Error())
		return err
	}

	db.SingularTable(DBSingularTable)
	db.DB().SetMaxIdleConns(cfg.UInt("max_idle_conns"))
	db.DB().SetMaxOpenConns(cfg.UInt("max_open_conns"))
	DB = MyDB{db}
	db.LogMode(true) //TODO need to be deleted
	if os.Getenv("MODE") == "" {
		db.LogMode(true)
	}
	db.SetLogger(Logger{})
	return nil
}

func Destory() {
	DB.Close()
}
