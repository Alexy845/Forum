package requetes_sql

import (
	"fmt"
	"forum/hash"
	uuid2 "forum/uuid"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func AddUser(pseudo string, prenom string, nom string, mdp string, email string) uuid.UUID {
	sqL := "INSERT INTO User (Id, Pseudo, Prenom, Nom, Date_membre, MDP, Email) values (?, ?, ?, ?, ?, ?, ?)"
	req, err := DB.Prepare(sqL)
	if err != nil {
		fmt.Println("AddUser1")
		log.Fatal(err)
	}
	id := uuid2.CreateUUID()
	_, err = req.Exec(id, pseudo, prenom, nom, time.Now().Format("2006-01-02 15:04:05"), hash.Hash(mdp), email)
	if err != nil {
		fmt.Println("AddUser2")
		log.Fatal(err)
	}
	return id
}
