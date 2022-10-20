package business

import (
	"test/mnc/features/points"
)

type pointUsecase struct {
	pointData points.Data
}

func NewPointBusiness(pntData points.Data) points.Business {
	return &pointUsecase{
		pointData: pntData,
	}
}

func (uc *pointUsecase) PostPoint(inputPoint points.Core) (res int, err error) {
	res, err = uc.pointData.InsertPoint(inputPoint)
	return res, err
}

func (uc *pointUsecase) GetPointByIdArticle(idArticle int) (res []points.Core, err error) {
	res, err = uc.pointData.SelectPointByIdArticle(idArticle)
	return res, err
}

func (uc *pointUsecase) GetPointByIdUser(idUser int) (res []points.Core, err error) {
	res, err = uc.pointData.SelectPointByIdUser(idUser)
	return res, err
}
