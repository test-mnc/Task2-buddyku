package presentation

import (
	"test/mnc/features/users"
	_requestUser "test/mnc/features/users/presentation/request"
	_responseUser "test/mnc/features/users/presentation/response"
	_helper "test/mnc/helpers"
	"test/mnc/middlewares"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) AddUser(c echo.Context) error {
	var dataUser _requestUser.User
	errBind := c.Bind(&dataUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(dataUser)

	if len(dataUser.FullName) == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("fullname must be filled"))
	}

	errUserName := v.Var(dataUser.UserName, "required,alphanumunicode")
	if errUserName != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("username must without space"))
	}

	errEmail := v.Var(dataUser.Email, "required,email")
	if errEmail != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid format email"))
	}
	if len(dataUser.Password) == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("password must be filled"))
	}
	errPhoneNumber := v.Var(dataUser.PhoneNumber, "required,numeric")
	if errPhoneNumber != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("phone number must be in numeric"))
	}
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed(errValidator.Error()))
	}
	dataUsr := _requestUser.ToCore(dataUser)
	row, err := h.userBusiness.InsertUser(dataUsr)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email already exist"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("Succes to insert data"))
}

func (h *UserHandler) Login(c echo.Context) error {
	var userLogin _requestUser.User
	errLog := c.Bind(&userLogin)
	if errLog != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	token, fullName, e := h.userBusiness.LoginUser(userLogin.Email, userLogin.Password)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
	}
	data := map[string]interface{}{
		"full_name": fullName,
		"token":     token,
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("login success", data))
}

func (h *UserHandler) GetUser(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}

	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("unauthorized"))
	}
	result, err := h.userBusiness.SelectUser(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data user"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	idToken, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("token in unvalid"))
	}

	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, _helper.ResponseFailed("missing token"))
	}

	res, err := h.userBusiness.SelectAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data all users"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success to get data  all users", _responseUser.FromCoreList(res)))
}
