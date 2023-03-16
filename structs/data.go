package structs

type Data struct {
	Posts      []Post
	User       User
	Connected  bool
	Post       Post
	Comments   []Comment
	Comment    Comment
	EditPost   bool
	Result     bool
	Categories []Category
}

var Datas = &Data{
	Posts:      []Post{},
	User:       User{},
	Connected:  false,
	Post:       Post{},
	Comments:   []Comment{},
	Comment:    Comment{},
	EditPost:   false,
	Result:     false,
	Categories: []Category{},
}
