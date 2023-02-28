package requetes_sql

import (
	"fmt"
	"forum/fonctions"
	"log"
	"strconv"
	"time"
)

func AddUser(pseudo string, prenom string, nom string, mdp string, email string) {
	sqL := "INSERT INTO User (Id, Pseudo, Prenom, Nom, Date_membre, MDP, Email) values (?, ?, ?, ?, ?, ?, ?)"
	req, err := DB.Prepare(sqL)
	if err != nil {
		fmt.Println("AddUser1")
		log.Fatal(err)
	}
	_, err = req.Exec(strconv.Itoa(GetTailleUser()+1), pseudo, prenom, nom, time.Now().Format("2006-01-02 15:04:05"), fonctions.Hash(mdp), email)
	if err != nil {
		fmt.Println("AddUser2")
		log.Fatal(err)
	}
}
