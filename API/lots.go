package API

import (
	"GoParking/models"
	"net/http"
      
	"github.com/gin-gonic/gin"
)

func GetDetails(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Honeycut)
}
func Addlots(c *gin.Context) {
	var newlot models.Lot
	if err := c.BindJSON(&newlot); err != nil {
		return
	}
	c.IndentedJSON(http.StatusCreated, newlot)
	models.Honeycut[0].Lots = append(models.Honeycut[0].Lots, newlot)
}

