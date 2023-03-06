package requetes_sql

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

func UpdateUser(id uuid.UUID, pseudo string, email string, prenom string, nom string, avatar string) {
	_, err := DB.Exec("UPDATE User SET Pseudo = ?, Email = ?, Prenom = ?, Nom = ?, Avatar = ? WHERE id = ?", pseudo, email, prenom, nom, avatar, id)
	if err != nil {
		log.Fatal("UpdateUser: ", err)
	}
}
