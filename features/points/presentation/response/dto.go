package response

import (
	"test/mnc/features/points"
	"time"
)

type Point struct {
	ID        int       `json:"id"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	Article   Article   `json:"article"`
	Company   Company   `json:"company"`
	User      User      `json:"user"`
}

type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Company struct {
	ID           int    `json:"id"`
	EmployeeName string `json:"employee_name"`
}

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
}

func FromCore(data points.Core) Point {
	return Point{
		ID:        data.ID,
		Value:     data.Value,
		CreatedAt: data.CreatedAt,
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
