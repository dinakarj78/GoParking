package models

import (
	// uuid "github.com/satori/go.uuid"
)

	type Parkinglot struct {
		Location string         `json:"location"`
		Capacity int            `json:"capacity"`
		Lots     []Lot          `json:lots`
	}

var Parkinglotstruct = []Parkinglot{}

func Parkinginit() {
		Parkinglotstruct = []Parkinglot{
			{Location: "clapperlane",
				Capacity: 0,
				Lots: []Lot{
					// {
					// 	SlotType: "SUV",
					// 	Id:       uuid.NewV4(),
					// 	Location: "clapper lane",
					// },
				},

			},
	}
}
