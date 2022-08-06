package controllers

import (
	models "animerest/Models"
	services "animerest/Services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnimeController struct {
	*services.AnimeService
}

// @BasePath /api/
// @Summary get anime list
// @Schemes
// @Tags AnimeController
// @Accept json
// @Produce json
// @Router /anime/ [get]
func (u *AnimeController) GetAll(ctx *gin.Context) {
	u.AnimeService.GetPaginatedList(ctx)
}

// @BasePath /api/
// @Summary get anime by id
// @Schemes
// @Tags AnimeController
// @Accept json
// @Produce json
// @Param id query string required "model"
// @Param code query string unrequired "model"
// @Param t query string unrequired "model"
// @Router /anime/title [get]
func (u *AnimeController) GetById(ctx *gin.Context) {
	id := ctx.Query("id")
	code := ctx.Query("code")
	t := ctx.Query("t")

	idNum, err := strconv.Atoi(id)
	if err == nil {
		u.AnimeService.FindById(ctx, models.AnimeFilter{
			Id:   idNum,
			Code: code,
			T:    t,
		})
	}
}
