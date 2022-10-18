package users

import "time"

type Core struct {
	ID          int
	FullName    string
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	InsertUser(data Core) (row int, err error)
	LoginUser(email string, password string) (fullName string, token string, e error)
	SelectUser(id int) (data Core, err error)
	SelectAllUser() (data []Core, err error)
}

type Data interface {
	PostUser(data Core) (row int, err error)
	AuthUser(email string, password string) (fullName string, token string, e error)
	GetUser(id int) (data Core, err error)
	GetAllUser() (data []Core, err error)
}
