package API

import (
	"GoParking1/models"
	// "fmt"
)

func DeleteFromlot(location *models.Parkinglot, vno string, vtype string) {
	// if Id, ok := Allotedlots[vno]; ok {
	// 	location.Emptylots[Id] = vtype
	// 	location.Capacity = location.Capacity + 1
	// 	location.TypeCap[vtype] = location.TypeCap[vtype] + 1
	// 	delete(Allotedlots, vno)
	// 	vno = SearchQueue(vtype)
	// 	if vno != "" {
	// 		fmt.Println("added from wait queue")
	// 		// AddtoSlot(location, vno, vtype)
	// 	}
	// 	fmt.Println("Remaining lots:", location.Capacity)
	// 	fmt.Println("suv lots: ", location.TypeCap["SUV"], "Sedan lots: ", location.TypeCap["SEDAN"], "Truck lots: ", location.TypeCap["TRUCK"])
	// } else {
	// 	fmt.Println("Vehicle not found..!")
	// }
}
