package request

import "test/mnc/features/company"

type Company struct {
	EmployeeName string `json:"employee_name" form:"employee_name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	PhoneNumber  string `json:"phone_number" form:"phone_number"`
}

func ToCore(req Company) company.Core {
	return company.Core{
		EmployeeName: req.EmployeeName,
		Password:     req.Password,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
	}
}
