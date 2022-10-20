package business

import (
	"fmt"
	"test/mnc/features/articles"
	"test/mnc/features/points"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockArticleData struct{}

func (mock mockArticleData) SelectArticleById(idArticle int) (articles.Core, error) {
	return articles.Core{
		ID:        1,
		Title:     "Testing Judul",
		Content:   "Testing Content.....",
		CreatedAt: time.Time{},
		User: articles.User{
			ID:          1,
			FullName:    "Musa King",
			UserName:    "musaking",
			Email:       "musa@gmail.com",
			PhoneNumber: "08232147896",
		},
	}, nil
}

func (mock mockArticleData) InsertArticle(input articles.Core) (int, int, error) {
	return 1, 1, nil
}

func (mock mockArticleData) SelectAllArticles() ([]articles.Core, error) {
	return []articles.Core{
		{
			ID:        1,
			Title:     "Judul 1",
			Content:   "Content 1",
			CreatedAt: time.Time{},
			User: articles.User{
				ID:          1,
				FullName:    "Musa",
				UserName:    "musa",
				Email:       "musa@gmail.com",
				PhoneNumber: "082346789",
			},
		},
	}, nil
}

type mockArticleDataFailed struct{}

func (mock mockArticleDataFailed) SelectArticleById(idArticle int) (articles.Core, error) {
	return articles.Core{}, fmt.Errorf("failed to get data article")
}

func (mock mockArticleDataFailed) InsertArticle(input articles.Core) (int, int, error) {
	return 0, 0, fmt.Errorf("failed to insert data article")
}

func (mock mockArticleDataFailed) SelectAllArticles() ([]articles.Core, error) {
	return nil, fmt.Errorf("failed to get all articles")
}

type mockPointData struct{}

func (mock mockPointData) InsertPoint(inputPoint points.Core) (int, error) {
	return 1, nil
}

func (mock mockPointData) SelectPointByIdArticle(idArticle int) ([]points.Core, error) {
	return []points.Core{
		{
			ID:        1,
			Value:     1.0,
			CreatedAt: time.Time{},
			User: points.User{
				ID:       2,
				FullName: "Kurniawan",
				UserName: "kurniawan",
			},
			Article: points.Article{
				ID:      1,
				Title:   "Judul 1",
				Content: "Content 1 .....",
			},
			Company: points.Company{
				ID:           1,
				EmployeeName: "employee 1",
			},
		},
	}, nil
}

func (mock mockPointData) SelectPointByIdUser(idUser int) ([]points.Core, error) {
	return []points.Core{
		{
			ID:        2,
			Value:     4.0,
			CreatedAt: time.Time{},
			User: points.User{
				ID:       2,
				FullName: "Kurniawan",
				UserName: "kurniawan",
			},
			Article: points.Article{
				ID:      1,
				Title:   "Judul 1",
				Content: "Content 1 .....",
			},
			Company: points.Company{
				ID:           1,
				EmployeeName: "employee 1",
			},
		},
	}, nil
}

func (mock mockPointData) SelectPointAllUsers() ([]points.Core, error) {
	return []points.Core{
		{
			ID:        1,
			Value:     4.0,
			CreatedAt: time.Time{},
			User: points.User{
				ID:       2,
				FullName: "Kurniawan",
				UserName: "kurniawan",
			},
			Article: points.Article{
				ID:      1,
				Title:   "Judul 1",
				Content: "Content 1 .....",
			},
			Company: points.Company{
				ID:           1,
				EmployeeName: "employee 1",
			},
		},
	}, nil
}

func (mock mockPointData) FirstPoint(idUser int, ArticleID int, CompanyID int, value float64) (row int, idArticle int, err error) {
	return 1, 1, nil
}

type mockPointDataFailed struct{}

func (mock mockPointDataFailed) InsertPoint(inputPoint points.Core) (int, error) {
	return 0, fmt.Errorf("failed to insert data")
}

func (mock mockPointDataFailed) SelectPointByIdArticle(idArticle int) ([]points.Core, error) {
	return nil, fmt.Errorf("failed to get point")
}

func (mock mockPointDataFailed) SelectPointByIdUser(idUser int) ([]points.Core, error) {
	return nil, fmt.Errorf("failed to get point")
}

func (mock mockPointDataFailed) SelectPointAllUsers() ([]points.Core, error) {
	return nil, fmt.Errorf("failed to get all point")
}

func (mock mockPointDataFailed) FirstPoint(idUser int, ArticleID int, CompanyID int, value float64) (row int, idArticle int, err error) {
	return 0, 0, fmt.Errorf("failed to insert point")
}

func TestGetArticleById(t *testing.T) {
	t.Run("Test Get Article By Id Success", func(t *testing.T) {
		articleBusiness := NewArticleBusiness(mockArticleData{}, mockPointData{})
		result, err := articleBusiness.GetArticleById(0)
		assert.Nil(t, err)
		assert.Equal(t, "Testing Judul", result.Title)
	})
	t.Run("Test Get Article By Id Failed", func(t *testing.T) {
		articleBusiness := NewArticleBusiness(mockArticleDataFailed{}, mockPointDataFailed{})
		result, err := articleBusiness.GetArticleById(0)
		assert.NotNil(t, err)
		assert.Equal(t, articles.Core{}, result)
	})
}

func TestPostArticle(t *testing.T) {
	t.Run("Test Post Article Success", func(t *testing.T) {
		articleBusiness := NewArticleBusiness(mockArticleData{}, mockPointData{})
		input := articles.Core{
			Title:   "Testing Judul",
			Content: "Testing Content",
			User: articles.User{
				ID: 1,
			},
		}
		result, idArticle, err := articleBusiness.PostArticle(input)
		assert.Nil(t, err)
		assert.Equal(t, 1, 1, result, idArticle)
	})
	t.Run("Test Post Article Failed", func(t *testing.T) {
		articleBusiness := NewArticleBusiness(mockArticleDataFailed{}, mockPointDataFailed{})
		input := articles.Core{
			Title:   "Testing Judul",
			Content: "Testing Content",
			User: articles.User{
				ID: 1,
			},
		}
		result, idArticle, err := articleBusiness.PostArticle(input)
		assert.NotNil(t, err)
		assert.Equal(t, 0, 0, result, idArticle)
	})
}

func TestGetAllArticles(t *testing.T) {
	t.Run("Test Get All Articles Success", func(t *testing.T) {
		articleBusiness := NewArticleBusiness(mockArticleData{}, mockPointData{})
		res, err := articleBusiness.GetAllArticles()
		assert.Nil(t, err)
		assert.Equal(t, []articles.Core{
			{
				ID:        1,
				Title:     "Judul 1",
				Content:   "Content 1",
				CreatedAt: time.Time{},
				User: articles.User{
					ID:          1,
					FullName:    "Musa",
					UserName:    "musa",
					Email:       "musa@gmail.com",
					PhoneNumber: "082346789",
				},
			},
		}, res)
	})
	t.Run("Test Get All Articles Failed", func(t *testing.T) {
		articleBusiness := NewArticleBusiness(mockArticleDataFailed{}, mockPointDataFailed{})
		res, err := articleBusiness.GetAllArticles()
		assert.NotNil(t, err)
		assert.Nil(t, nil, res)
	})
}
