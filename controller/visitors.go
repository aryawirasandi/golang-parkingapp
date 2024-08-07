package controller

import (
	"net/http"

	"github.com/aryawirasandi/parking-app/entity"
	"github.com/aryawirasandi/parking-app/model"
)

func (app *Controller) GetListVisitor(c e) error {
	m := &model.Model{
		Database: app.DB,
	}

	result, err := m.GetVisitors()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, entity.ReponseWithData[[]entity.Visitor]{
		Message: "lists of visitors",
		Data:    result,
	})
}
