package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Conn(env *Env) *sql.DB {
	host := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", env.Username, env.Password, env.Host, env.Port, env.Database)
	result, err := sql.Open("mysql", host)
	fmt.Print(host)
	if err != nil {
		fmt.Println(err.Error())
	}

	return result
}
