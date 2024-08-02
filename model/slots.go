package model

import (
	"github.com/aryawirasandi/parking-app/entity"
)

func (m Model) CreateSlot(visitorId string, slotName string) (entity.Slot, error) {

	_, errDb := m.Database.Exec("INSERT INTO slots (visitor_id, slot_name) VALUES (?, ?)", visitorId, slotName)

	var user entity.Slot

	if errDb != nil {
		return user, errDb
	}
	resultUser := m.Database.QueryRow("SELECT * FROM slots WHERE slot_name = ?", slotName)
	resultUser.Scan(&user.Id, &user.VisitorId, &user.SlotName)
	return user, nil
}
