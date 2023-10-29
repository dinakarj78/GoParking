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
func Addlots(c *gin.Context) {
	var newlot []models.Lot
	if err := c.BindJSON(&newlot); err != nil {
		return
	}
	// c.IndentedJSON(http.StatusCreated, newlot)
	var i int = 0
	for _, lot := range newlot {
		// newlotmap:=make(map[uuid.UUID]string)
		// newlotmap[lot.Id]=lot.Location
		if models.Parkinglotstruct[i].Location == lot.Location {
			models.Parkinglotstruct[i].Lots = append(models.Parkinglotstruct[i].Lots, lot)
			c.IndentedJSON(http.StatusCreated, models.Parkinglotstruct)
		} else {
			i++
		}
	}
}
func Addlocation(c *gin.Context) {
	var locationLots []models.Parkinglot
	if err := c.BindJSON(&locationLots); err != nil {
		c.IndentedJSON(http.StatusCreated, "hello1")
		return
	}
	for _, lot := range locationLots {
		models.Parkinglotstruct = append(models.Parkinglotstruct, lot)
		c.IndentedJSON(http.StatusOK, "hello")
	}
}
