package service

import (
	"fmt"
	"os"

	"github.com/aryawirasandi/parking-app/controller"
	"github.com/aryawirasandi/parking-app/database"
	ml "github.com/aryawirasandi/parking-app/middleware"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Mapper(conn *database.Env) *Server {
	result := database.Conn(conn)
	server := &Server{
		DB: result,
		e:  echo.New(),
		contrl: &controller.Controller{
			DB: result,
		},
		config: &echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(ml.JWTClaims)
			},
			SigningKey: []byte(os.Getenv("JWT_SECRET")),
		},
	}
	return server
}

func (server Server) Run(portRunning string) {
	fullPathPort := fmt.Sprintf(":%s", portRunning)
	server.e.Use(middleware.Logger())
	server.e.Use(middleware.Recover())
	g := server.e.Group("/api/v1")
	r := g.Group("")
	r.Use(echojwt.WithConfig(*server.config))

	g.POST("/auth/login", server.contrl.GetUser)
	g.POST("/auth/register", server.contrl.CreateUser)

	// get lists of a
	r.POST("/slots", server.contrl.GetSlots)

	server.e.Logger.Fatal(server.e.Start(fullPathPort))
}
