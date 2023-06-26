package models

import uuid "github.com/satori/go.uuid"

// "fmt"
// "ticketing/Config"
// "github.com/google/uuid"

type Lot struct {
	SlotType string
	Id       uuid.UUID
	lotno    int
}
