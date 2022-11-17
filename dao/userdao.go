package dao

import (
	"ginessnital/global"
	"log"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func GetAllUsers() []user {
	rows, err := global.Mysql.Query("SELECT * FROM tb_user")
	log.Println(err)
	log.Println(&rows, "rows")
	if err != nil {
		return nil
	}
	var persons = make([]user, 0)
	for rows.Next() {
		var a user
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
