package response

import (
	"test/mnc/features/company"
	"time"
)

type Company struct {
	ID           int       `json:"id"`
	EmployeeName string    `json:"employee_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
}

func FromCore(data company.Core) Company {
	return Company{
		ID:           data.ID,
		EmployeeName: data.EmployeeName,
		Email:        data.Email,
		PhoneNumber:  data.PhoneNumber,
		CreatedAt:    data.CreatedAt,
	}
}

func FromCoreList(data []company.Core) []Company {
	result := []Company{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
