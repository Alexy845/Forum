package structs

import uuid "github.com/satori/go.uuid"

type User struct {
	Id           uuid.UUID
	Username     string
	Password     string
	Email        string
	FirstName    string
	LastName     string
	CreationDate string
}
