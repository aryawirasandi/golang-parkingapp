package controller

import (
	"net/http"

	"github.com/aryawirasandi/parking-app/entity"
	"github.com/aryawirasandi/parking-app/model"
)

func (app *Controller) GetUser(c ech) error {
	u := new(entity.Credential)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Message: "Bad Request!",
		})
	}

	m := &model.Model{
		Database: app.DB,
	}

	result, err := m.GetUser(u.Username, u.Password)

	if err != nil {
		return c.JSON(http.StatusForbidden, entity.Response{
			Message: "Username Or Password is wrong",
		})
	}

	if result.Id == 0 {
		return c.JSON(http.StatusNotFound, entity.Response{
			Message: "Username or Password is wrong",
		})
	}

	return c.JSON(http.StatusOK, result)
}

func (app *Controller) CreateUser(c ech) error {
	u := new(entity.Credential)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Message: "Bad Request!",
		})
	}

	m := &model.Model{
		Database: app.DB,
	}

	result, err := m.CreateUser(u.Username, u.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Message: "Something Wrong With Server",
		})
	}

	return c.JSON(http.StatusOK, entity.ReponseWithData[entity.User]{
		Data:    result,
		Message: "Succesfully Created User",
	})
}
