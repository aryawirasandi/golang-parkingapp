package controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	DB *sql.DB
}

type ech = echo.Context
