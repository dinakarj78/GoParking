package models

import (
	uuid "github.com/satori/go.uuid"
)

type Parkinglot struct {
	Location  string               `json:"location"`
	Capacity  int                  `json:"capacity"`
	Lots      []Lot                `json:lots`
	TypeCap   map[string]int       //defines map between type of vehical and number of slots for that type
	Emptylots map[uuid.UUID]string //map between empty slots and type of that slot
}

var Honeycut = []Parkinglot{}

func Parkinginit() {
	Honeycut = []Parkinglot{
		{Location: "clapperlane",
			Capacity: 14,
			Lots: []Lot{{
				SlotType: "SUV",
				Id:       uuid.NewV4()}, {SlotType: "Sedan", Id: uuid.NewV4()}, {SlotType: "Truck", Id: uuid.NewV4()},
			},
			TypeCap:   map[string]int{"Suv": 5, "Sedan": 5, "Truck": 4},
			Emptylots: map[uuid.UUID]string{}},
	}
	for _, lot := range Honeycut[0].Lots {
		Honeycut[0].Emptylots[lot.Id] = lot.SlotType
	}
}
