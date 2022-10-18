package data

import (
	"test/mnc/features/articles"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
	User    User
}

type User struct {
	gorm.Model
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Email       string `gorm:"unique" json:"email"`
	PhoneNumber string `json:"phone_number"`
	Article     []Article
}

func (data *Article) toCore() articles.Core {
	return articles.Core{
		ID:        int(data.ID),
		Title:     data.Title,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: articles.User{
			ID:          int(data.User.ID),
			FullName:    data.User.FullName,
			UserName:    data.User.UserName,
			Email:       data.User.Email,
			PhoneNumber: data.User.PhoneNumber,
		},
	}
}

func toCoreList(data []Article) []articles.Core {
	result := []articles.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core articles.Core) Article {
	return Article{
		Title:   core.Title,
		Content: core.Content,
		UserID:  core.User.ID,
	}
}
