package structs

import uuid "github.com/satori/go.uuid"

type Post struct {
	Id      uuid.UUID
	Auteur  uuid.UUID
	Titre   string
	Contenu string
	Date    string
}
