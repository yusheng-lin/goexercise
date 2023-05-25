package api

import (
	"fmt"
	"goexercise/models"
	"goexercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	acctsvc *service.AccountService
	jwtsvc  *service.JWTService
}

func NewAccountController(acctsvc *service.AccountService, jwtsvc *service.JWTService) *AccountController {
	return &AccountController{
		acctsvc: acctsvc,
		jwtsvc:  jwtsvc,
	}
}

// @Summary uers login
// @Tags user
// @version 1.0
// @produce application/json
// @param login body models.Login true "login"
// @Success 200 {object} models.Response 成功後返回的值
// @Router /Login [post]
func (controller *AccountController) Login(ctx *gin.Context) {
	var login models.Account
	err := ctx.BindJSON(&login)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{Msg: "Invalid params"})
		return
	}

	acct, err := controller.acctsvc.Login(login.Acct, login.Pwd)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{Msg: fmt.Sprint(err)})
		return
	}

	token, _ := controller.jwtsvc.GenToken(acct.Acct)

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: token})
}

// @Summary uers sign up
// @Tags user
// @version 1.0
// @produce application/json
// @param singup body models.Account true "singup"
// @Success 200 {object} models.Response 成功後返回的值
// @Router /SignUp [post]
func (controller *AccountController) SignUp(ctx *gin.Context) {
	singup := &models.Account{}
	err := ctx.BindJSON(singup)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{Msg: "Invalid params"})
		return
	}

	err = controller.acctsvc.SignUp(singup)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{Msg: fmt.Sprint(err)})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success"})
}

// @Summary delete account
// @Tags user
// @version 1.0
// @produce application/json
// @Security BearerAuth
// @param acct body models.DeleteAccount true "acct"
// @Success 200 {object} models.Response 成功後返回的值
// @Router /user [delete]
func (controller *AccountController) Delete(ctx *gin.Context) {
	acct := &models.Account{}
	err := ctx.BindJSON(acct)

	if err != nil || acct.Acct == "" {
		ctx.JSON(http.StatusOK, models.Response{Msg: "Invalid params"})
		return
	}

	err = controller.acctsvc.DeleteAccount(acct.Acct)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{Msg: fmt.Sprint(err)})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success"})
}

// @Summary update account
// @Tags user
// @version 1.0
// @produce application/json
// @Security BearerAuth
// @param acct body models.Account true "acct"
// @Success 200 {object} models.Response 成功後返回的值
// @Router /user [put]
func (controller *AccountController) Update(ctx *gin.Context) {
	acct := &models.Account{}
	err := ctx.BindJSON(acct)

	if err != nil || acct.Acct == "" {
		ctx.JSON(http.StatusOK, models.Response{Msg: "Invalid params"})
		return
	}

	err = controller.acctsvc.UpadteAccount(acct)

	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{Msg: fmt.Sprint(err)})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Msg: "Success"})
}
