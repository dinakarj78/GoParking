package models

import uuid "github.com/satori/go.uuid"

type vehicle struct {
	VehicleNo   string
	VehicleType string
	AllotedLot  uuid.UUID
}
