package data

import (
	"test/mnc/features/company"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	EmployeeName string `json:"employee_name"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phone_number"`
}

func (data *Company) toCore() company.Core {
	return company.Core{
		ID:           int(data.ID),
		EmployeeName: data.EmployeeName,
		Password:     data.Password,
		Email:        data.Email,
		PhoneNumber:  data.PhoneNumber,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func toCoreList(data []Company) []company.Core {
	result := []company.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core company.Core) Company {
	return Company{
		EmployeeName: core.EmployeeName,
		Password:     core.Password,
		Email:        core.Email,
		PhoneNumber:  core.PhoneNumber,
	}
}
