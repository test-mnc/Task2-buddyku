package response

import (
	"test/mnc/features/points"
	"time"
)

type Point struct {
	ID        int    `json:"id"`
	Value     string `json:"value"`
	ArticleID int    `json:"article_id"`
	CompanyID int    `json:"company_id"`
	UserID    int    `json:"user_id"`
	CreatedAt time.Time
	Article   Article
	Company   Company
	User      User
}

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Company struct {
	ID           int    `json:"id"`
	EmployeeName string `json:"employee_name"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
}

type User struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Email       string `gorm:"unique" json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func FromCore(data points.Core) Point {
	return Point{
		ID:    data.ID,
		Value: data.Value,
		Article: Article{
			ID:    data.Article.ID,
			Title: data.Article.Title,
		},
		User: User{
			ID:       data.User.ID,
			UserName: data.User.UserName,
			FullName: data.User.FullName,
		},
		Company: Company{
			ID:           data.Company.ID,
			EmployeeName: data.Company.EmployeeName,
		},
	}
}

func FromCoreList(data []points.Core) []Point {
	result := []Point{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
