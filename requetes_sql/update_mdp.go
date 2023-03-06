package requetes_sql

import (
	"forum/structs"
	"log"
)

func UpdateMDP(mdp string) {
	_, err := DB.Exec("UPDATE User SET MDP = ? WHERE id = ?", mdp, structs.Datas.User.Id)
	if err != nil {
		log.Fatal("UpdateMDP: ", err)
	}
}
