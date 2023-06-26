package API

import (
	"fmt"
	"ticketing/models"
)

func DeleteFromlot(location *models.Parkinglot, vno string, vtype string) {
	Id, ok := Allotedlots[vno]
	if ok {
		location.Emptylots[Id] = vtype
		location.Capacity = location.Capacity + 1
		location.TypeCap[vtype] = location.TypeCap[vtype] + 1
		delete(Allotedlots, vno)
		fmt.Println("Remaining lots:", location.Capacity)
		fmt.Println("suv lots: ", location.TypeCap["SUV"], "Sedan lots: ", location.TypeCap["SEDAN"], "Truck lots: ", location.TypeCap["TRUCK"])
	} else {
		fmt.Println("Vehicle not found..!")
	}
}
