package structs

import uuid "github.com/satori/go.uuid"

type Category struct {
	Id   uuid.UUID
	Name string
}
