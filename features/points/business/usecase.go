package business

import "test/mnc/features/points"

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
