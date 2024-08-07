package model

import (
	"database/sql"
	"log"

	"github.com/aryawirasandi/parking-app/entity"
)

func (m Model) GetVisitors() ([]entity.Visitor, error) {
	slots := []entity.Visitor{}

	vl, err := m.Database.Query("SELECT * FROM visitors")

	if err != nil {
		return slots, err
	}

	for vl.Next() {
		var v entity.Visitor
		var nco sql.NullTime
		var nci sql.NullTime
		var createdAt sql.NullString
		var updatedAt sql.NullString
		if err := vl.Scan(&v.Id, &v.VehicleNumber, &v.TypeVehicle, &nci, &nco, &createdAt, &updatedAt); err != nil {
			log.Fatal(err)
		} else {
			v.ClockIn = &nci.Time
			v.ClockOut = &nco.Time
			v.TimeStamp.CreatedAt = &createdAt.String
			v.TimeStamp.UpdatedAt = &updatedAt.String
		}
		slots = append(slots, v)
	}

	defer vl.Close()

	return slots, nil
}
