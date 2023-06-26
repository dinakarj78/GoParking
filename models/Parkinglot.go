package models

// "fmt"
import (
	uuid "github.com/satori/go.uuid"
)

type Parkinglot struct {
	Id        uuid.UUID
	Lots      []Lot
	Capacity  int
	Location  string
	TypeCap   map[string]int       //defines map between type of vehical and number of slots for that type
	Emptylots map[uuid.UUID]string //map between empty slots and type of that slot
}
