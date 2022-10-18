package request

import "test/mnc/features/users"

type User struct {
	FullName    string `json:"full_name" form:"full_name"`
	UserName    string `json:"user_name" form:"user_name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

func ToCore(req User) users.Core {
	return users.Core{
		FullName:    req.FullName,
		UserName:    req.UserName,
		Password:    req.Password,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
}
