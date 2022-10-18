package presentation

import (
	"net/http"
	"test/mnc/features/points"
	_requestPoint "test/mnc/features/points/presentation/request"
	_helper "test/mnc/helpers"
	"test/mnc/middlewares"

	"github.com/labstack/echo/v4"
)

type PointHandler struct {
	pointBusiness points.Business
}

func NewPointHandler(business points.Business) *PointHandler {
	return &PointHandler{
		pointBusiness: business,
	}
}

func (h *PointHandler) InsertValue(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("token is unvalid "))
	}

	var givePoint = _requestPoint.Point{
		CompanyID: idToken,
	}

	errBind := c.Bind(&givePoint)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to binding data"))
	}

	res, err := h.pointBusiness.PostPoint(_requestPoint.ToCore(givePoint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert value"))
	}

	if res != 1 {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert/record data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("success to insert data"))
}
