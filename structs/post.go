package structs

import (
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	Id       uuid.UUID
	Auteur   User
	Titre    string
	Contenu  string
	Date     string
	Likes    int
	Dislikes int
}
