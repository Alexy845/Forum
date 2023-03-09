package structs

import uuid "github.com/satori/go.uuid"

type Dislike struct {
	Id   uuid.UUID
	User User
	Post Post
}
