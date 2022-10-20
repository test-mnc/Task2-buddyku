package business

import (
	"fmt"
	"test/mnc/features/users"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockUserData struct{}

func (mock mockUserData) PostUser(data users.Core) (int, error) {
	return 1, nil
}

func (mock mockUserData) AuthUser(email string, password string) (string, string, error) {
	return "andri gunawan", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTg1MDM2MjQsInVzZXJJZCI6N30.DFxWe18OoFhRAIt41dIDWPQPORI8S5rAkTylJiolnbc", nil
}

func (mock mockUserData) GetUser(id int) (users.Core, error) {
	return users.Core{
		ID: 1, FullName: "Andri G", Email: "andri@gmail.com", PhoneNumber: "081234", CreatedAt: time.Time{},
	}, nil
}

func (mock mockUserData) GetAllUser() ([]users.Core, error) {
	return []users.Core{
		{ID: 1, FullName: "Andri G", Email: "andri@gmail.com", PhoneNumber: "081234", CreatedAt: time.Time{}},
	}, nil
}

type mockUserDataFailed struct{}

func (mock mockUserDataFailed) PostUser(data users.Core) (int, error) {
	return 0, fmt.Errorf("Failed to select data")
}

func (mock mockUserDataFailed) AuthUser(email string, password string) (string, string, error) {
	return "", "", fmt.Errorf("email or password incorrect")
}

func (mock mockUserDataFailed) GetUser(id int) (users.Core, error) {
	return users.Core{}, fmt.Errorf("failed to get data")
}

func (mock mockUserDataFailed) GetAllUser() ([]users.Core, error) {
	return nil, fmt.Errorf("failed to get data")
}

func TestInsertUser(t *testing.T) {
	t.Run("Test Insert User Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			FullName:    "Andri G",
			Email:       "andri@gmail.com",
			Password:    "qwerty",
			PhoneNumber: "081234",
		}
		result, err := userBusiness.InsertUser(newUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Insert Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			FullName:    "Andri G",
			Email:       "andri@gmail.com",
			Password:    "qwerty",
			PhoneNumber: "081234",
		}
		result, err := userBusiness.InsertUser(newUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, result)
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("Test Login User Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		var loginEmail = "andri@gmail.com"
		var loginPass = "qwerty"

		resFullName, resToken, err := userBusiness.LoginUser(loginEmail, loginPass)
		assert.Nil(t, err)
		assert.Equal(t, "andri gunawan", "andri gunawan", resFullName, resToken)
	})

	t.Run("Test Login User Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		var loginEmail = "andri@gmail.com"
		var loginPass = "qwerty"

		resFullName, resToken, err := userBusiness.LoginUser(loginEmail, loginPass)
		assert.NotNil(t, err)
		assert.Equal(t, "", "", resFullName, resToken)
	})
}

func TestSelectUser(t *testing.T) {
	t.Run("Test Select User Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.SelectUser(0)
		assert.Nil(t, err)
		assert.Equal(t, "Andri G", result.FullName)
	})
	t.Run("Test Select User Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.SelectUser(0)
		assert.NotNil(t, err)
		assert.Equal(t, users.Core{}, result)
	})
}

func TestSelectAllUser(t *testing.T) {
	t.Run("Test Select User Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.SelectAllUser()
		assert.Nil(t, err)
		assert.Equal(t, []users.Core{
			{
				ID:          1,
				FullName:    "Andri G",
				Email:       "andri@gmail.com",
				PhoneNumber: "081234",
				CreatedAt:   time.Time{},
			},
		}, result)
	})
	t.Run("Test Select User Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.SelectAllUser()
		assert.NotNil(t, err)
		assert.Nil(t, nil, result)
	})
}
