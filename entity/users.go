package entity

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"token"`
	Role     string `json:"role"`
	TimeStamp
}

type Users = []User
