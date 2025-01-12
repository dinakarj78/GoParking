package API

import (
	// "GoParking/models"
	"GoParking/models"
	"net/http"
 
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var Allotedlots = make(map[string]uuid.UUID)

// add vehicles
func AllotLots(c *gin.Context) {
	var newVehicle = models.Vehicle
	// Route := gin.Default() //Returns the router

	if err := c.BindJSON(&newVehicle); err != nil {
		return
	}
	// c.Redirect(http.StatusPermanentRedirect, "/AddtoQueue")
	for i := 0; i < len(newVehicle); i++ {
		 if newVehicle[i].Location == models.Parkinglotstruct[0].Location && models.Parkinglotstruct[0].Capacity!=0{
			Allotedlots[newVehicle[0].VehicleNo] = uuid.UUID{}
			models.Parkinglotstruct[0].Capacity--
			c.IndentedJSON(http.StatusCreated, "lotalloted")
		}else if models.Parkinglotstruct[0].Capacity==0 {
			c.IndentedJSON(http.StatusCreated, "addingtoqueue")
			c.Redirect(http.StatusPermanentRedirect, "/AddtoQueue")
	    } else {
			i++
		}
	}
	// c.IndentedJSON(http.StatusCreated, Allotedlots)

}
