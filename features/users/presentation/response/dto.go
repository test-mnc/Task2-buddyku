package response

import (
	"test/mnc/features/users"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:          data.ID,
		FullName:    data.FullName,
		UserName:    data.UserName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		CreatedAt:   data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
