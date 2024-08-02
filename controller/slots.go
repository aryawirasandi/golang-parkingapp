package controller

import (
	"net/http"

	"github.com/aryawirasandi/parking-app/entity"
	"github.com/aryawirasandi/parking-app/model"
)

func (app *Controller) GetSlots(c ech) error {
	u := new(entity.Slot)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, entity.Response{
			Message: "Bad Request!",
		})
	}

	m := &model.Model{
		Database: app.DB,
	}

	result, err := m.CreateSlot(u.SlotName, u.VisitorId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.Response{
			Message: "Something Wrong With Server",
		})
	}

	return c.JSON(http.StatusOK, entity.ReponseWithData[entity.Slot]{
		Data:    result,
		Message: "Succesfully Created User",
	})
}
