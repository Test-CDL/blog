package presentation

import (
	"chrombit/features/users"
	_requestUser "chrombit/features/users/presentation/request"
	_helpers "chrombit/helpers"

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
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(dataUser)
	// errFullName := v.Var(dataUser.FullName, "required,alpha")
	// if errFullName != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("fullname can only contains alphabet"))
	// }
	if len(dataUser.UserName) == 0 {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("username must be filled"))
	}

	if len(dataUser.Password) == 0 {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("password must be filled"))
	}

	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed(errValidator.Error()))
	}
	dataUsr := _requestUser.ToCore(dataUser)
	row, err := h.userBusiness.InsertUser(dataUsr)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("username already exist"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helpers.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, _helpers.ResponseSuccesNoData("Succes to insert data"))
}

func (h *UserHandler) Login(c echo.Context) error {
	var userLogin _requestUser.User
	errLog := c.Bind(&userLogin)
	if errLog != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("failed to bind data, check your input"))
	}
	// v := validator.New()
	// errValidator := v.Struct(userLogin)
	// if errValidator != nil {
	// 	return c.JSON(http.StatusBadRequest, _helper.ResponseFailed(errValidator.Error()))
	// }
	token, userName, e := h.userBusiness.LoginUser(userLogin.UserName, userLogin.Password)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helpers.ResponseFailed("username or password incorrect"))
	}
	data := map[string]interface{}{
		"user_name": userName,
		"token":     token,
	}
	return c.JSON(http.StatusOK, _helpers.ResponseSuccesWithData("LOGIN SUCCES", data))
}
