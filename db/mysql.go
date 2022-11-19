package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/xormplus/xorm"
)

type MySQLConf struct {
	Host         string
	Port         uint16
	Db           string
	Username     string
	Password     string
	Charset      string
	Timeout      string
	ParseTime    bool
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

var Mysql *xorm.Engine

func InitMySQL(cfg *MySQLConf) (err error) {
	mysqlLink := fmt.Sprintf("%v:%v@(%v:%d)/%v?charset=%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Db, cfg.Charset)
	db, err := xorm.NewEngine("mysql", mysqlLink)
	if err != nil {
		return
	}
	// Sets the maximum number of idle connections in the connection pool.
	db.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	// Sets the maximum number of connections to the database.
	db.DB().SetMaxOpenConns(cfg.MaxOpenConns)
	// Sets the maximum reusable time for the connection.
	db.DB().SetConnMaxLifetime(time.Hour)
	//Presentation logs
	db.ShowSQL(true)
	Mysql = db
	return
}

var MysqlGroup *xorm.EngineGroup

func InitEngineGroup(conf ...*MySQLConf) (err error) {
	if len(conf) <= 0 {
		err = errors.New("config is nil ")
		return
	}
	var Conf MySQLConf
	var cfg []string
	for k, v := range conf {
		if k == 0 {
			Conf = *v
		}
		mysqlLink := fmt.Sprintf("%v:%v@(%v:%d)/%v?charset=%v", v.Username, v.Password, v.Host, v.Port, v.Db, v.Charset)
		cfg = append(cfg, mysqlLink)
	}

	db, err := xorm.NewEngineGroup("mysql", cfg)
	if err != nil {
		return
	}
	// Sets the maximum number of idle connections in the connection pool.
	db.DB().SetMaxIdleConns(Conf.MaxIdleConns)
	// Sets the maximum number of connections to the database.
	db.DB().SetMaxOpenConns(Conf.MaxOpenConns)
	// Sets the maximum reusable time for the connection.
	db.DB().SetConnMaxLifetime(time.Hour)
	//Presentation logs
	db.ShowSQL(true)
	MysqlGroup = db
	return
}
