package API

import (
	"GoParking1/models"
	"net/http"


	"github.com/gin-gonic/gin"
)

var waitQueue = make(map[string]string)
var count = 0

func AddtoQueue(C *gin.Context) {
	C.IndentedJSON(http.StatusCreated,"addingtoqueue")
	var newVehicle1 = models.Vehicle
	if err := C.BindJSON(&newVehicle1); err != nil {
	C.IndentedJSON(http.StatusBadGateway,err)
		
	}
	for i:=range(newVehicle1){
	if _, ok := waitQueue[newVehicle1[i].VehicleNo]; ok {
	C.IndentedJSON(http.StatusAccepted,newVehicle1[i].VehicleNo+"is already in queue")
	}else{
	waitQueue[newVehicle1[i].VehicleNo] = newVehicle1[i].VehicleType
	C.IndentedJSON(http.StatusAccepted,newVehicle1[i].VehicleNo+"is added to queue"+"queue is:"+waitQueue[newVehicle1[i].VehicleNo])
}
	}

	// if _, ok := waitQueue[vno]; ok {
	// 	fmt.Println("Already in the queue")
	// 	return
	// }
	// waitQueue[vno] = vtype
	// count = count + 1
	// fmt.Println("added to queue")
	}

func SearchQueue(vtype string) string {
	var vno string
	for key, value := range waitQueue {
		if value == vtype {
			vno = key
		} else {
			return "none"
		}
	}
	return vno
}
