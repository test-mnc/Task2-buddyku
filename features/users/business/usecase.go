package business

import (
	"test/mnc/features/users"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (uc *userUsecase) InsertUser(input users.Core) (row int, err error) {
	row, err = uc.userData.PostUser(input)
	return row, err
}

func (uc *userUsecase) LoginUser(email string, password string) (fullName string, token string, e error) {
	fullName, token, e = uc.userData.AuthUser(email, password)
	return fullName, token, e
}

func (uc *userUsecase) SelectUser(id int) (resp users.Core, err error) {
	resp, err = uc.userData.GetUser(id)
	return resp, err
}

func (uc *userUsecase) SelectAllUser() (res []users.Core, err error) {
	res, err = uc.userData.GetAllUser()
	return res, err
}
