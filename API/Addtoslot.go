package API

import (
	"fmt"
	"ticketing/models"

	uuid "github.com/satori/go.uuid"
)

var Allotedlots = make(map[string]uuid.UUID)

func AddtoSlot(location *models.Parkinglot, vno string, vtype string) {
	if capacity, ok := location.TypeCap[vtype]; ok && capacity != 0 {
		for i := range location.Emptylots {
			if location.Emptylots[i] == vtype {
				location.Capacity = location.Capacity - 1
				location.TypeCap[vtype] = location.TypeCap[vtype] - 1
				delete(location.Emptylots, i)
				Allotedlots[vno] = i
				break
			}
		}
	} else {
		fmt.Println("No such type vehicle or  number of lots")
	}
	fmt.Println("Remaining lots:", location.Capacity)
	fmt.Println("suv lots: ", location.TypeCap["SUV"], "Sedan lots: ", location.TypeCap["SEDAN"], "Truck lots: ", location.TypeCap["TRUCK"])

}
