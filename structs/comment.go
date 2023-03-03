package structs

import uuid "github.com/satori/go.uuid"

type Comment struct {
	Id      uuid.UUID
	Auteur  User
	Post    Post
	Contenu string
	Date    string
}
