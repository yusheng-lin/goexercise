package api

import (
	"fmt"
	"goexercise/models"
	"goexercise/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{
		service: svc,
	}
}

// @Summary list users
// @Tags user
// @version 1.0
// @produce application/json
// @Security BearerAuth
// @Success 200 string string 成功後返回的值
// @Router /user [get]
func (controller *UserController) Users(ctx *gin.Context) {
	users := controller.service.GetUsers()
	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: users})
}

// @Summary get user by name
// @Tags user
// @version 1.0
// @produce application/json
// @Security BearerAuth
// @param name path string true "name"
// @Success 200 string string 成功後返回的值
// @Router /user/{name} [get]
func (controller *UserController) User(ctx *gin.Context) {
	name := ctx.Param("name")
	user, err := controller.service.GetUser(name)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{Msg: fmt.Sprint(err)})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Msg: "Success", Data: user})
}
