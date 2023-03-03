package structs

type Data struct {
	Posts     []Post
	User      User
	Connected bool
	Post      Post
	Comments  []Comment
}

var Datas = &Data{
	Posts:     []Post{},
	User:      User{},
	Connected: false,
	Post:      Post{},
}
