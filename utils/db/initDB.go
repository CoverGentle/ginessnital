package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "admin.1234"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "gotest"
)

// 初始化数据库连接，返回数据库连接的指针引用
func InitDB() *sql.DB {
	//Golang数据连接："用户名:密码@tcp(IP:端口号)/数据库名?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := sql.Open("mysql", path)
	if err != nil {
		//如果打开数据库错误，直接panic
		panic(err)
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	//验证连接
	if err := db.Ping(); err != nil {
		panic(err)
	}

	//将数据库连接的指针引用返回
	return db
}
