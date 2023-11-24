package models

import (
	uuid "github.com/satori/go.uuid"
)

	type Parkinglot struct {
		Location string         `json:"location"`
		Capacity int            `json:"capacity"`
		Lots     []Lot          `json:lots`
		TypeCap  map[string]int //defines map between type of vehicle and number of slots for that type
	}

var Parkinglotstruct = []Parkinglot{}

func Parkinginit() {
		Parkinglotstruct = []Parkinglot{
			{Location: "clapperlane",
				Capacity: 0,
				Lots: []Lot{
					{
						SlotType: "SUV",
						Id:       uuid.NewV4(),
						Location: "clapperlane",
					},
				},
				TypeCap: map[string]int{"SUV": 1, "Sedan": 0, "Truck": 0},
			},
	}
}
