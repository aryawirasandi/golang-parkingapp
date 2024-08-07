package entity

import "time"

type Visitor struct {
	Id            int        `json:"id"`
	VehicleNumber string     `json:"vehicle_number"`
	ClockIn       *time.Time `json:"clock_in"`
	ClockOut      *time.Time `json:"clock_out"`
	TypeVehicle   string     `json:"type_vehicle"`
	TimeStamp
}
