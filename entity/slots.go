package entity

type Slot struct {
	Id        int    `json:"id"`
	VisitorId string `json:"visitor_id"`
	SlotName  string `json:"slot_name"`
}
