package insta

import (
	"math/rand"

	"gomatcha.io/matcha/view/ios"

	golorem "github.com/drhodes/golorem"
)

type App struct {
	Stack *ios.Stack
	Posts []*Post
}

func NewApp() *App {
	posts := []*Post{}
	for i := 0; i < 5; i++ {
		posts = append(posts, GeneratePost())
	}

	return &App{
		Stack: &ios.Stack{},
		Posts: posts,
	}
}

type Post struct {
	Id int64
	// UserId       int64
	UserName     string
	UserImageURL string
	Location     string
	ImageURL     string
	LikeCount    int
	// CommentIds   []int64
	Comments []*Comment

	Liked      bool
	Bookmarked bool
}

func GeneratePost() *Post {
	comments := []*Comment{}
	for i := 0; i < rand.Intn(20); i++ {
		comments = append(comments, GenerateComment())
	}

	return &Post{
		Id:           0,
		UserName:     golorem.Word(5, 15),
		UserImageURL: "http://lorempixel.com/50/50/",
		Location:     golorem.Word(5, 15),
		ImageURL:     "http://lorempixel.com/400/400/",
		LikeCount:    rand.Intn(500),
		Comments:     comments,
	}
}

type Comment struct {
	Id       int64
	UserId   int64
	UserName string
	Text     string
}

func GenerateComment() *Comment {
	return &Comment{
		Id:       0,
		UserId:   0,
		UserName: golorem.Word(5, 15),
		Text:     golorem.Paragraph(1, 5),
	}
}

type User struct {
	Id       int64
	Name     string
	ImageURL string
}
