package structs

type Data struct {
	Posts []Post
	User  User
}

var Datas = &Data{
	Posts: []Post{},
	User:  User{},
}
