package model

import (
	"errors"

	"github.com/aryawirasandi/parking-app/entity"
)

func (m Model) CreateSlot(slotName string) (entity.Slot, error) {

	var user entity.Slot

	findFirst := m.Database.QueryRow("SELECT * FROM slots WHERE slot_name = ?", slotName)

	findFirst.Scan(&user.Id, &user.VisitorId, &user.SlotName)

	if user.Id != 0 {
		return user, errors.New("can't duplicate parking code")
	}

	_, errDb := m.Database.Exec("INSERT INTO slots (slot_name) VALUES (?)", slotName)

	if errDb != nil {
		return user, errDb
	}
	resultUser := m.Database.QueryRow("SELECT * FROM slots WHERE slot_name = ?", slotName)
	resultUser.Scan(&user.Id, &user.VisitorId, &user.SlotName)
	return user, nil
}
