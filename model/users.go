package model

import (
	"log"
	"os"
	"time"

	"github.com/aryawirasandi/parking-app/entity"
	"github.com/aryawirasandi/parking-app/middleware"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (m Model) GetUser(username string, password string) (entity.User, error) {
	var (
		id        int
		usr       string
		pwd       string
		role      string
		createdAt string
		updatedAt string
	)

	result := m.Database.QueryRow("SELECT * FROM users WHERE username = ?", username)
	result.Scan(&id, &usr, &pwd, &role, &createdAt, &updatedAt)

	if err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password)); err != nil {
		return entity.User{}, err
	}

	claims := middleware.JWTClaims{
		Username: usr,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		Id:       id,
		Username: usr,
		Password: pwd,
		Token:    t,
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
	resultUser.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	return user, nil
}
