package models

type vehicle struct {
	VehicleNo   string `json:"VehicleNo"`
	VehicleType string `json:"VehicleType"`
	Location    string `json:"location"`
}

var Vehicle = []vehicle{}
