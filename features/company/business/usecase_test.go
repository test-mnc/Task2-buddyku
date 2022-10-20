package business

import (
	"fmt"
	"test/mnc/features/company"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCompanyData struct{}

func (mock mockCompanyData) PostEmployee(input company.Core) (int, error) {
	return 1, nil
}

func (mock mockCompanyData) AuthEmployee(email string, password string) (string, string, error) {
	return "employee 1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTg1MDM2MjQsInVzZXJJZCI6N30.DFxWe18OoFhRAIt41dIDWPQPORI8S5rAkTylJiolnbc", nil
}

type mockCompanyDataFailed struct{}

func (mock mockCompanyDataFailed) PostEmployee(input company.Core) (int, error) {
	return 0, fmt.Errorf("failed to insert data employee")
}

func (mock mockCompanyDataFailed) AuthEmployee(email string, password string) (string, string, error) {
	return "", "", fmt.Errorf("failed to login employee")
}

func TestInsertEmployee(t *testing.T) {
	t.Run("Test Insert Employee Success", func(t *testing.T) {
		companyBusiness := NewCompanyBusiness(mockCompanyData{})
		input := company.Core{
			EmployeeName: "Employee 1",
			Email:        "employee1@gmail.com",
			Password:     "employee1",
			PhoneNumber:  "08342167",
		}
		row, err := companyBusiness.InsertEmployee(input)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
	t.Run("Test Insert Employee Failed", func(t *testing.T) {
		companyBusiness := NewCompanyBusiness(mockCompanyDataFailed{})
		input := company.Core{
			EmployeeName: "Employee 1",
			Email:        "employee1@gmail.com",
			Password:     "employee1",
			PhoneNumber:  "08342167",
		}
		row, err := companyBusiness.InsertEmployee(input)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestLoginEmployee(t *testing.T) {
	t.Run("Test Login Employee Success", func(t *testing.T) {
		companyBusiness := NewCompanyBusiness(mockCompanyData{})
		fullName, token, err := companyBusiness.LoginEmployee("employee1@gmail.com", "employee1")
		assert.Nil(t, err)
		assert.Equal(t, "employee 1", "employee 1", fullName, token)
	})
	t.Run("Test Login Employee Failed", func(t *testing.T) {
		companyBusiness := NewCompanyBusiness(mockCompanyDataFailed{})
		fullName, token, err := companyBusiness.LoginEmployee("employee1@gmail.com", "employee1")
		assert.NotNil(t, err)
		assert.Equal(t, "", "", fullName, token)
	})
}
