package main

import (
	// "GoParking/API"
	"GoParking/models"
	"GoParking1/API"
	"fmt"

	"github.com/gin-gonic/gin"
)

// var Honeycut = models.Parkinglot{}

func main() {
	models.Parkinginit()
	router := gin.Default() //Returns the router
	router.GET("/getDetails", API.GetDetails)
	router.POST("/Addlocation", API.Addlocation)
	router.POST("/addLots", API.Addlots)
	router.Run("localhost:8080")
}

// func Lotinit(Type string, capacity int, lots []models.Lot) []models.Lot {
// 	Lot := models.Lot{}

// 	for i := 0; i < capacity; i++ {
// 		Lot.SlotType = Type
// 		Lot.Id = uuid.NewV4()
// 		lots = append(lots, Lot)
// 	}
// 	return lots
// }

// func Parkinginit(location string, capacity int, truckCapacity int, suvcapacity int, sedancapacity int) {
// 	Honeycut.Id = uuid.NewV4()
// 	Honeycut.Capacity = capacity
// 	Honeycut.Location = location

// 	Honeycut.Lots = []models.Lot{}
// 	Honeycut.Lots = Lotinit("SUV", suvcapacity, Honeycut.Lots)
// 	Honeycut.Lots = Lotinit("SEDAN", sedancapacity, Honeycut.Lots)
// 	Honeycut.Lots = Lotinit("TRUCK", truckCapacity, Honeycut.Lots)

// 	Honeycut.Emptylots = make(map[uuid.UUid]string)

// 	for j := 0; j < capacity; j++ { //adding empty lots to the map
// 		Honeycut.Emptylots[Honeycut.Lots[j].id] = Honeycut.Lots[j].SlotType
// 	}
// 	Honeycut.TypeCap = make(map[string]int)
// 	Honeycut.TypeCap["SUV"] = suvcapacity
// 	Honeycut.TypeCap["SEDAN"] = sedancapacity
// 	Honeycut.TypeCap["TRUCK"] = truckCapacity
// }

//	func Input() {
//		i := 0
//		for i < 4 {
//			fmt.Println("Enter your choice:\n 1: ADD a car to slot \n 2: Remove your vehicle")
//			var choice int
//			switch fmt.Scanln(&choice); choice {
//			case 1:
//				vno, vtype := VehicleDetails()
//				API.AddtoSlot(&Honeycut, vno, vtype)
//			case 2:
//				vno, vtype := VehicleDetails()
//				API.DeleteFromlot(&Honeycut, vno, vtype)
//			default:
//				fmt.Println("Enter a correct choice")
//			}
//			i++
//		}
//	}
func VehicleDetails() (string, string) {
	var vno string
	var vtype string

	fmt.Println("Enter  vehicle number")
	fmt.Scanln(&vno)

	fmt.Println("Enter vehicle type ")
	fmt.Scanln(&vtype)

	return vno, vtype
}
