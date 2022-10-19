package data

import (
	"test/mnc/features/points"

	"gorm.io/gorm"
)

type Point struct {
	gorm.Model
	Value     string  `json:"value"`
	ArticleID int     `json:"article_id"`
	CompanyID int     `json:"company_id"`
	UserID    int     `json:"user_id"`
	Article   Article `json:"article"`
	Company   Company `json:"company"`
	User      User    `json:"user"`
}

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Company struct {
	gorm.Model
	EmployeeName string `json:"employee_name"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
}

type User struct {
	gorm.Model
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Email       string `gorm:"unique" json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (data *Point) toCore() points.Core {
	return points.Core{
		ID:        int(data.ID),
		Value:     data.Value,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Article: points.Article{
			ID:      int(data.Article.ID),
			Title:   data.Article.Title,
			Content: data.Article.Content,
		},
		User: points.User{
			ID:          int(data.User.ID),
			FullName:    data.User.FullName,
			UserName:    data.User.UserName,
			Email:       data.User.Email,
			PhoneNumber: data.User.PhoneNumber,
		},
		Company: points.Company{
			ID:           int(data.Company.ID),
			EmployeeName: data.Company.EmployeeName,
			Email:        data.Company.Email,
			PhoneNumber:  data.Company.PhoneNumber,
		},
	}
}

func toCoreList(data []Point) []points.Core {
	result := []points.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core points.Core) Point {
	return Point{
		Value:     core.Value,
		ArticleID: core.Article.ID,
		CompanyID: core.Company.ID,
		UserID:    core.User.ID,
	}
}
