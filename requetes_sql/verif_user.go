package requetes_sql

func Verifuser(username string, password string) int {
	for _, user := range Getall() {
		if user.Username == username && user.Password == password {
			return user.Id
		}
	}
	return 0
}
