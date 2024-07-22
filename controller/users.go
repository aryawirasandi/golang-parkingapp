package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	DB *sql.DB
}

func (app *Controller) GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
