package response

import (
	"test/mnc/features/articles"
	"time"
)

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	User      User
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func FromCore(data articles.Core) Article {
	return Article{
		ID:        data.ID,
		Title:     data.Title,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		User: User{
			ID:          data.User.ID,
			FullName:    data.User.FullName,
			UserName:    data.User.UserName,
			Email:       data.User.Email,
			PhoneNumber: data.User.PhoneNumber,
		},
	}
}

func FromCoreList(data []articles.Core) []Article {
	result := []Article{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
