package presentation

import (
	"test/mnc/features/company"
	_requestCompany "test/mnc/features/company/presentation/request"
	_helper "test/mnc/helpers"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CompanyHandler struct {
	companyBusiness company.Business
}

func NewCompanyHandler(business company.Business) *CompanyHandler {
	return &CompanyHandler{
		companyBusiness: business,
	}
}

func (h *CompanyHandler) AddEmployee(c echo.Context) error {
	var dataEmployee _requestCompany.Company
	errBind := c.Bind(&dataEmployee)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(dataEmployee)
	// errFullName := v.Var(dataUser.FullName, "required,alpha")
	// if errFullName != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("fullname can only contains alphabet"))
	// }
	if len(dataEmployee.EmployeeName) == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("employeename must be filled"))
	}
	errEmail := v.Var(dataEmployee.Email, "required,email")
	if errEmail != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid format email"))
	}
	if len(dataEmployee.Password) == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("password must be filled"))
	}
	errPhoneNumber := v.Var(dataEmployee.PhoneNumber, "required,numeric")
	if errPhoneNumber != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("phone number must be in numeric"))
	}
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed(errValidator.Error()))
	}
	dataInput := _requestCompany.ToCore(dataEmployee)
	row, err := h.companyBusiness.InsertEmployee(dataInput)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email already exist"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesNoData("Succes to insert data"))
}

func (h *CompanyHandler) Login(c echo.Context) error {
	var employeeLogin _requestCompany.Company
	errLog := c.Bind(&employeeLogin)
	if errLog != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	// v := validator.New()
	// errValidator := v.Struct(userLogin)
	// if errValidator != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed(errValidator.Error()))
	// }
	token, employeeName, e := h.companyBusiness.LoginEmployee(employeeLogin.Email, employeeLogin.Password)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
	}
	data := map[string]interface{}{
		"employee_name": employeeName,
		"token":         token,
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("LOGIN SUCCES", data))
}
