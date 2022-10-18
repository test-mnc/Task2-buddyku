package company

import "time"

type Core struct {
	ID           int
	EmployeeName string
	Email        string
	Password     string
	PhoneNumber  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Business interface {
	InsertEmployee(data Core) (row int, err error)
	LoginEmployee(email string, password string) (employeeName string, token string, e error)
}

type Data interface {
	PostEmployee(data Core) (row int, err error)
	AuthEmployee(email string, password string) (employeeName string, token string, e error)
}
