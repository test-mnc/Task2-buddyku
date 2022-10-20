package business

import (
	"fmt"
	"test/mnc/features/points"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestPostPoint(t *testing.T) {
	t.Run("Test Post Point Success", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointData{})
		input := points.Core{
			Value: 4,
			Article: points.Article{
				ID: 1,
			},
			User: points.User{
				ID: 2,
			},
			Company: points.Company{
				ID: 1,
			},
		}
		res, err := pointBusiness.PostPoint(input)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
	})
	t.Run("Test Post Point Failed", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointDataFailed{})
		input := points.Core{
			Value: 4,
			Article: points.Article{
				ID: 1,
			},
			User: points.User{
				ID: 2,
			},
			Company: points.Company{
				ID: 1,
			},
		}
		res, err := pointBusiness.PostPoint(input)
		assert.NotNil(t, err)
		assert.Equal(t, 0, res)
	})
}

func TestGetPointByIdArticle(t *testing.T) {
	t.Run("Test Get Point By Id Article Success", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointData{})
		res, err := pointBusiness.GetPointByIdArticle(0)
		assert.Nil(t, err)
		assert.Equal(t, []points.Core{
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
		}, res)
	})
	t.Run("Test Get Point By Id Article Failed", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointDataFailed{})
		res, err := pointBusiness.GetPointByIdArticle(0)
		assert.NotNil(t, err)
		assert.Nil(t, nil, res)
	})
}

func TestGetPointByIdUser(t *testing.T) {
	t.Run("Test Get Point By Id User Success", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointData{})
		res, err := pointBusiness.GetPointByIdUser(0)
		assert.Nil(t, err)
		assert.Equal(t, []points.Core{
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
		}, res)
	})
	t.Run("Test Get Point By Id User Failed", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointDataFailed{})
		res, err := pointBusiness.GetPointByIdUser(0)
		assert.NotNil(t, err)
		assert.Nil(t, nil, res)
	})
}

func TestGetPointAllUsers(t *testing.T) {
	t.Run("Test Get Point All Users", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointData{})
		res, err := pointBusiness.GetPointAllUsers()
		assert.Nil(t, err)
		assert.Equal(t, []points.Core{
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
		}, res)
	})
	t.Run("Test Get Point All Users Failed", func(t *testing.T) {
		pointBusiness := NewPointBusiness(mockPointDataFailed{})
		res, err := pointBusiness.GetPointAllUsers()
		assert.NotNil(t, err)
		assert.Nil(t, nil, res)
	})
}
