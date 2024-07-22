package service

import (
	"database/sql"

	"github.com/aryawirasandi/parking-app/controller"
	"github.com/labstack/echo/v4"
)

type Server struct {
	DB     *sql.DB
	e      *echo.Echo
	contrl *controller.Controller
}
