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
	var i int =0
    for _,lot:=range newlot{
		// newlotmap:=make(map[uuid.UUID]string)
        // newlotmap[lot.Id]=lot.Location
		if models.Parkinglotstruct[i].Location==lot.Location{
	models.Parkinglotstruct[i].Lots=append(models.Parkinglotstruct[i].Lots, lot)
   }else{
    i++
   }
}
c.IndentedJSON(http.StatusCreated, models.Parkinglotstruct)

	// for _, lots := range models.Parkinglotstruct {
	// 	if lots.Location == newlot.Location {
	// 		c.IndentedJSON(http.StatusCreated, newlot)
	// 		models.Parkinglotstruct[i].Lots = append(models.Parkinglotstruct[i].Lots, newlot)
	// 		fmt.Println("HELLO")
	// 		i++
	// 	}
}
