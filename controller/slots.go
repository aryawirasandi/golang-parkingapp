package controller

import (
	"net/http"

	"github.com/aryawirasandi/parking-app/entity"
	"github.com/aryawirasandi/parking-app/model"
)

func (app *Controller) CreateSlot(c e) error {
	u := new(entity.Slot)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Message: "Bad Request!",
		})
	}

	m := &model.Model{
		Database: app.DB,
	}

	result, err := m.CreateSlot(u.SlotName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.ReponseWithData[entity.Slot]{
		Message: "Succesfully Created User",
		Data:    result,
	})
}

func (app *Controller) GetListOfSlot(c e) error {

	m := &model.Model{
		Database: app.DB,
	}

	result, err := m.GetListOfSlot()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.ReponseWithData[[]entity.Slot]{
		Message: "Successfully Get a lists of user",
		Data:    result,
	})
}
