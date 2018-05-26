// 数据库连接
package db

import (
	"common/cfg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func New(config *cfg.Configuration) (db *gorm.DB, err error) {
	if db, err = gorm.Open("mysql", config.Database.Dsn); err != nil {
		return
	}
	if err = db.DB().Ping(); err != nil {
		return
	}
	db.DB().SetMaxOpenConns(config.Database.MaxConn)
	db.LogMode(true)
	return
}
