package service

import (
	"fmt"

	"github.com/aryawirasandi/parking-app/controller"
	"github.com/aryawirasandi/parking-app/database"
	"github.com/labstack/echo/v4"
)

func Mapper(conn *database.Env) *Server {
	result := database.Conn(conn)
	server := &Server{
		DB: result,
		e:  echo.New(),
		contrl: &controller.Controller{
			DB: result,
		},
	}

	return server
}

func (server Server) Run(portRunning string) {
	fullPathPort := fmt.Sprintf(":%s", portRunning)
	g := server.e.Group("/api/v1/")
	g.GET("/", server.contrl.GetUser)
	server.e.Logger.Fatal(server.e.Start(fullPathPort))
}
