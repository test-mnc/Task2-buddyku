package articles

import "time"

type Core struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}

type User struct {
	ID          int
	FullName    string
	UserName    string
	Email       string
	PhoneNumber string
}

type Business interface {
	GetArticleById(idArticle int) (data Core, err error)
	PostArticle(data Core) (row int, idArticle int, err error)
	GetAllArticles() (data []Core, err error)
}

type Data interface {
	SelectArticleById(idArticle int) (data Core, err error)
	InsertArticle(data Core) (row int, idArticle int, err error)
	SelectAllArticles() (data []Core, err error)
}
