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
	var u model.RegisterUser
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
func LoginUser(username string, password string) (err error, reg model.RegisterUser) {
	rows, err := global.Mysql.Query("select username,password from sys_user where username =?", username)
	if err != nil {
		fmt.Println("查询失败")
		return err, reg
	}
	var u model.RegisterUser
	for rows.Next() {
		//var u model.RegisterUser
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Age)
		if err != nil {
			fmt.Println(err)

		}
	}
	return err, u
	//return err, reg
}

// 获取所有用户数据
func GetAllUsers() []model.User {
	rows, err := global.Mysql.Query("select * from tb_user")
	if err != nil {
		return nil
	}
	var persons = make([]model.User, 0)
	for rows.Next() {
		var a model.User
		err := rows.Scan(&a.ID, &a.Name, &a.Password)
		log.Println(&a.ID, "&a.ID")
		log.Println(&a.ID, "&a.Username")
		log.Println(&a.ID, "&a.Password")
		//if err != nil {
		//	return nil
		//}
		log.Println(err, "err")
		log.Println(persons, "persons")
		persons = append(persons, a)

	}
	return persons
}
