package API

import (
	"GoParking/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// uuid "github.com/satori/go.uuid"
)

func GetDetails(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Parkinglotstruct)
}
//add lots to a location
func Addlots(c *gin.Context) {
	var newlot []models.Lot
	if err := c.BindJSON(&newlot); err != nil {
		return
	}
	var i int = 0
	var addedFlag=false

	for _, lot := range newlot {
		if models.Parkinglotstruct[i].Location == lot.Location {
			models.Parkinglotstruct[i].Lots = append(models.Parkinglotstruct[i].Lots, lot)
			models.Parkinglotstruct[i].Capacity++
			addedFlag=true
		} else {
			i++
		}
	}
	if(!addedFlag) {
		c.IndentedJSON(http.StatusCreated, "adding a new location feature is not available yet")
	}else{
	c.IndentedJSON(http.StatusCreated, models.Parkinglotstruct)
	}
}
//adding a location..yet to be written
func Addlocation(c *gin.Context) {
	var locationLots []models.Parkinglot
	if err := c.BindJSON(&locationLots); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	for _, lot := range locationLots {
		models.Parkinglotstruct = append(models.Parkinglotstruct, lot)
		c.IndentedJSON(http.StatusOK, "hello")
	}
}
