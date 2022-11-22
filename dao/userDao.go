package dao

import (
	"database/sql"
	"fmt"
	"ginessnital/global"
	"ginessnital/model"
	"log"
)

// 注册
func RegisterUser(username string, password string, age int) error {
	//查询单条数据
	row := global.Mysql.QueryRow("select * from sys_user where username =?", username)
	log.Println(row, "row")
	var u model.User
	err := row.Scan(&u.Username, &u.Password, &u.Age)
	log.Println(err, "err")
	//sql: expected 4 destination arguments in Scan, not 3 err
	//sql: no rows in result set
	if err == sql.ErrNoRows {
		exec, err := global.Mysql.Exec("insert into sys_user(username,password,age) values (?,?,?) ", username, password, age)
		log.Println(&exec, "exec")
		if err != nil {
			return err
		}
		_, err = exec.LastInsertId()
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	} else {
		return err
	}
	return err
}

// 登录
func LoginUser(username string) (user *model.User) {
	var u model.User
	row := global.Mysql.QueryRow("select * from sys_user where username = ?", username)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Age)
	if err != nil {
		fmt.Println(&u, "err")
		//log.Panic(err)
		return nil
	}
	return &u
}

// 获取所有用户数据
func GetAllUsers() []model.User {
	rows, err := global.Mysql.Query("select * from sys_user")
	if err != nil {
		return nil
	}
	var persons = make([]model.User, 0)
	for rows.Next() {
		var a model.User
		err := rows.Scan(&a.ID, &a.Username, &a.Password, &a.Age)
		if err != nil {
			return nil
		}
		persons = append(persons, a)
	}
	return persons
}
