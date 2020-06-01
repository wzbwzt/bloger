package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

//Init 初始化
func Init(dns string) (err error) {
	DB, err = sqlx.Connect("mysql", dns)
	if err != nil {
		return
	}
	DB.SetMaxOpenConns(10) //设置连接池中最大的连接数
	DB.SetMaxIdleConns(10) //设置连接池中的最大闲置连接数
	return nil
}
