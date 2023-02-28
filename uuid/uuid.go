package uuid

import (
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
)

func CreateUUID() uuid.UUID {
	return uuid.NewV4()
}
