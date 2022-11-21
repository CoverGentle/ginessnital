package model

// 查询所有
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

// 注册用户
type RegisterUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
