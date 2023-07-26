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

var Parkinglotstruct = []Parkinglot{}

func Parkinginit() {
	Parkinglotstruct = []Parkinglot{
		{Location: "clapperlane",
			Capacity:  14,
			Lots:      []Lot{
				{
				SlotType:"SUV",
				Id:uuid.NewV4(),
				Location: "clapperlane",
			},
			},
			TypeCap:   map[string]int{"Suv": 5, "Sedan": 5, "Truck": 4},
			Emptylots: map[uuid.UUID]string{}},
	}
    
	for _, lot := range Parkinglotstruct[0].Lots {
		Parkinglotstruct[0].Emptylots[lot.Id] = lot.SlotType
	}
}
