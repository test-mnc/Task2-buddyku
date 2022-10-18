package data

import (
	"test/mnc/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Email       string `gorm:"unique" json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:          int(data.ID),
		FullName:    data.FullName,
		UserName:    data.UserName,
		Password:    data.Password,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core users.Core) User {
	return User{
		FullName:    core.FullName,
		UserName:    core.UserName,
		Password:    core.Password,
		Email:       core.Email,
		PhoneNumber: core.PhoneNumber,
	}
}
