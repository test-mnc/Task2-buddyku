package data

import (
	"fmt"
	"test/mnc/features/company"
	_bcrypt "test/mnc/plugins"

	"test/mnc/middlewares"

	"gorm.io/gorm"
)

type mysqlCompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(conn *gorm.DB) company.Data {
	return &mysqlCompanyRepository{
		db: conn,
	}
}

func (repo *mysqlCompanyRepository) PostEmployee(input company.Core) (row int, err error) {
	passHash, _ := _bcrypt.HashPassword(input.Password)
	company := Company{
		EmployeeName: input.EmployeeName,
		Password:     passHash,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
	}
	tx := repo.db.Create(&company)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlCompanyRepository) AuthEmployee(email string, password string) (employeeName string, token string, e error) {
	employeeData := Company{}
	repo.db.Where("email = ?", email).First(&employeeData)
	bool1 := _bcrypt.CheckPasswordHash(password, employeeData.Password)

	if !bool1 {
		return "", "", fmt.Errorf("error")
	}
	token, errToken := middlewares.CreateToken(int(employeeData.ID))
	if errToken != nil {
		return "", "", errToken
	}
	return token, employeeData.EmployeeName, nil
}
