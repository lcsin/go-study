package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

type User struct {
	ID                 int
	Nickname, Password string
	Telephone, Email   string
	Birth              time.Time
}

func initDB() (err error) {
	dsn := "root:123456@tcp(192.168.5.130:3306)/kratos_demo?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return err
}

func QueryRows(users *[]User) {
	sql := "select * from tb_user"
	err := db.Select(users, sql)
	if err != nil {
		log.Fatalln("query rows failed, err:", err)
		return
	}

	for idx, u := range *users {
		fmt.Printf("query success, user%d: id:%d name:%s age:%s\n", idx+1, u.ID, u.Nickname, u.Password)
	}
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatalln("initDB failed, err:", err)
	}
	r := gin.Default()

	var users []User
	r.GET("/userList", func(ctx *gin.Context) {
		QueryRows(&users)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "用户列表查询成功",
			"data": users,
		})
	})

	r.Run()
}
