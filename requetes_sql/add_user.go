package requetes_sql

import (
	"database/sql"
	"forum/fonctions"
	"log"
	"strconv"
	"time"
)

func AddUser(pseudo string, prenom string, nom string, mdp string, email string) {
	req, err := DB.Query("INSERT INTO User (Id, Pseudo, Prenom, Nom, Date_membre, MDP, Email) values (?, ?, ?, ?, ?, ?, ?)"+strconv.Itoa(GetTailleUser()+1), pseudo, prenom, nom, time.Now(), fonctions.Hash(mdp), email)
	if err != nil {
		log.Fatal(err)
	}
	defer func(req *sql.Rows) {
		err := req.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(req)
}
