package controllers

import (
	helpers "animerest/Helpers"
	models "animerest/Models"
	services "animerest/Services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*services.UserService
}

// @BasePath /api/
// @Summary login
// @Schemes
// @Tags UserController
// @Accept json
// @Produce json
// @Param model body models.UserLoginDto required "model"
// @Router /user/login [post]
func (u *UserController) Login(ctx *gin.Context) {
	var model models.UserLoginDto
	err := helpers.DecodeJson(ctx.Request.Body, &model)

	if err != nil {
		helpers.ErrorAction(ctx, err.Error(), 400)
	} else {
		u.UserService.Login(ctx, model)
	}
}

// @BasePath /api/
// @Summary register
// @Schemes
// @Tags UserController
// @Accept json
// @Produce json
// @Param model body models.UserRegisterDto required "model"
// @Router /user/register [post]
func (u *UserController) Register(ctx *gin.Context) {
	var model models.UserRegisterDto
	err := helpers.DecodeJson(ctx.Request.Body, &model)

	if err != nil {
		helpers.ErrorAction(ctx, err.Error(), 400)
	} else {
		u.UserService.Register(ctx, model)
	}
}
