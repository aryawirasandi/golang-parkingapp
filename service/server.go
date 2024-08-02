package service

import (
	"database/sql"

	"github.com/aryawirasandi/parking-app/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Server struct {
	DB     *sql.DB
	e      *echo.Echo
	contrl *controller.Controller
	config *echojwt.Config
}
