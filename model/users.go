package model

import (
	"fmt"
	"log"
	"time"

	"github.com/aryawirasandi/parking-app/entity"
	"golang.org/x/crypto/bcrypt"
)

func (m Model) GetUser(username string, password string) (entity.User, error) {
	var (
		id        int
		usr       string
		pwd       string
		token     string
		role      string
		createdAt string
		updatedAt string
	)

	result := m.Database.QueryRow("SELECT * FROM users WHERE username = ?", username)
	result.Scan(&id, &usr, &pwd, &token, &role, &createdAt, &updatedAt)
	fmt.Printf("Haaaah %s", role)

	if err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password)); err != nil {
		return entity.User{}, err
	}
	return entity.User{
		Id:       id,
		Username: usr,
		Password: pwd,
		Token:    token,
		Role:     role,
		TimeStamp: entity.TimeStamp{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}, nil
}

func (m Model) CreateUser(username string, password string) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)

	if err != nil {
		return entity.User{}, err
	}
	createAt := time.Now()

	_, errDb := m.Database.Exec("INSERT INTO users (username, password, createdAt) VALUES (?, ?, ?)", username, hash, createAt)

	if errDb != nil {
		log.Fatal(err)
		return entity.User{}, errDb
	}

	var user entity.User
	resultUser := m.Database.QueryRow("SELECT * FROM users WHERE username = ?", username)
	resultUser.Scan(&user.Id, &user.Username, &user.Password, &user.Token, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	return user, nil
}
