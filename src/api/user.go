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
// @param page body models.Paging true "page"
// @Success 200 {object} models.Response 成功後返回的值
// @Router /user [post]
func (controller *UserController) Users(ctx *gin.Context) {
	page := &models.Paging{}
	err := ctx.BindJSON(page)
	if err != nil || !page.Validate() {
		ctx.JSON(http.StatusOK, models.Response{Msg: "Invalid params"})
		return
	}
	rows, pages, users := controller.service.GetUsers(page)
	ctx.JSON(http.StatusOK,
		models.Response{Msg: "Success",
			Data: models.PagingResponse{
				TotalRows:  rows,
				TotalPages: pages,
				PageNo:     page.PageNo,
				Data:       users,
				Rows:       len(*users),
			}})
}

// @Summary get user by name
// @Tags user
// @version 1.0
// @produce application/json
// @Security BearerAuth
// @param name path string true "name"
// @Success 200 {object} models.Response 成功後返回的值
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
