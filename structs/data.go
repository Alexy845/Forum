package structs

type Data struct {
	Posts     []Post
	User      User
	Connected bool
}

var Datas = &Data{
	Posts:     []Post{},
	User:      User{},
	Connected: false,
}
