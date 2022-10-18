package business

import (
	"test/mnc/features/company"
)

type companyUsecase struct {
	companyData company.Data
}

func NewCompanyBusiness(cmpyData company.Data) company.Business {
	return &companyUsecase{
		companyData: cmpyData,
	}
}

func (uc *companyUsecase) InsertEmployee(input company.Core) (row int, err error) {
	row, err = uc.companyData.PostEmployee(input)
	return row, err
}

func (uc *companyUsecase) LoginEmployee(email string, password string) (employeeName string, token string, e error) {
	employeeName, token, e = uc.companyData.AuthEmployee(email, password)
	return employeeName, token, e
}
