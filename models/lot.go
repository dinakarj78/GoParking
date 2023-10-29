package models

import uuid "github.com/satori/go.uuid"

// "fmt"
// "ticketing/Config"
// "github.com/google/uuid"

type Lot struct {
	SlotType string    `json:"slottype"`
	Id       uuid.UUID `json:"id"`
	Location string    `json:"location"`
}

var Lots = []Lot{
	{
		SlotType: "SUV",
		Id:       uuid.NewV4(),
		Location: "clapper lane",
	},
}
