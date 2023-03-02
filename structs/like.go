package structs

import uuid "github.com/satori/go.uuid"

type Like struct {
	Id   uuid.UUID
	User User
	Post Post
}
