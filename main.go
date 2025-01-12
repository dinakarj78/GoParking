package main

import (
	// "GoParking/API"
	"GoParking/API"
	"GoParking/models"
	"github.com/gin-gonic/gin"
)

// var Honeycut = models.Parkinglot{}


func main() {
	models.Parkinginit()
	Router := gin.Default()                      //Returns the router
	Router.GET("/getDetails", API.GetDetails)    //get details on lots
	Router.POST("/addLots", API.Addlots)         //add lots to a location
	Router.POST("/allotLots", API.AllotLots)     //add vehicles to lots
	Router.POST("/AddtoQueue", API.AddtoQueue)   //add vehicles to queue
	Router.Run("localhost:8080")
}

