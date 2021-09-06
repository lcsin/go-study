package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_test?charset=utf8mb4&parseTime=True"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func Close() {
	DB.Close()
}
