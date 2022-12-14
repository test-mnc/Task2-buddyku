package points

import "time"

type Core struct {
	ID        int
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	Article   Article
	Company   Company
	User      User
}

type Article struct {
	ID      int
	Title   string
	Content string
}

type Company struct {
	ID           int
	EmployeeName string
	Email        string
	PhoneNumber  string
}

type User struct {
	ID          int
	FullName    string
	UserName    string
	Email       string
	PhoneNumber string
}

type Business interface {
	PostPoint(data Core) (row int, err error)
	GetPointByIdArticle(idArticle int) (data []Core, err error)
	GetPointByIdUser(idUser int) (data []Core, err error)
	GetPointAllUsers() (data []Core, err error)
}

type Data interface {
	InsertPoint(data Core) (row int, err error)
	SelectPointByIdArticle(idArticle int) (data []Core, err error)
	SelectPointByIdUser(idUser int) (data []Core, err error)
	FirstPoint(idUser int, ArticleID int, CompanyID int, value float64) (row int, idArticle int, err error)
	SelectPointAllUsers() (data []Core, err error)
}
