package services

import (
	config "animerest/Config"
	database "animerest/Database"
	helpers "animerest/Helpers"
	models "animerest/Models"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnimeService struct {
	*database.DbContext
}

func (u *AnimeService) FindById(ctx *gin.Context, filters models.AnimeFilter) {
	url := "?"

	if filters.Code != "" {
		url += "code=" + filters.Code
	}

	if filters.Id != 0 {
		url += "id=" + strconv.Itoa(filters.Id)
	}

	data, err := u.request(config.Anilibria, "GET", "getTitle"+url, nil, "")

	if err != nil {
		helpers.ErrorAction(ctx, err.Error(), 400)
		return
	}

	helpers.SuccessAction(ctx, "", 200, data)
}

func (u *AnimeService) GetPaginatedList(ctx *gin.Context) {
	data, err := u.request(config.Anilibria, "GET", "getUpdates", nil, "list")

	if err != nil {
		helpers.ErrorAction(ctx, err.Error(), 400)
		return
	}

	helpers.SuccessAction(ctx, "Успешно", 200, data)
}

func (u *AnimeService) request(site config.ApiConfiguration, method string, url string, body io.Reader, modelType string) (interface{}, error) {
	configuration := config.GetConfig()
	var globalErr error

	for index, link := range configuration.Api[site] {
		client := &http.Client{}
		r, err := http.NewRequest(method, link+url, body)

		if err != nil {
			if index < len(configuration.Api[site])-1 {
				continue
			} else {
				globalErr = err
			}
		} else {
			resp, _ := client.Do(r)

			if resp.StatusCode != 200 {
				globalErr = errors.New(strconv.Itoa(resp.StatusCode))
				continue
			}

			if site == config.Anilibria {
				if modelType == "list" {
					var result []models.Anime

					err := helpers.DecodeJson(resp.Body, &result)

					if err != nil {
						globalErr = err
						continue
					} else {
						globalErr = nil
						return result, nil
					}
				} else {
					var result models.Anime

					err := helpers.DecodeJson(resp.Body, &result)

					if err != nil {
						globalErr = err
						continue
					} else {
						globalErr = nil
						return result, nil
					}
				}
			} else {
				var result interface{}

				err := helpers.DecodeJson(resp.Body, &result)

				if err != nil {
					globalErr = err
					continue
				} else {
					globalErr = nil
					return nil, nil
				}
			}
		}
	}

	return nil, globalErr
}

func (u *AnimeService) FindByName(ctx *gin.Context, name string) {
	data, err := u.request(config.Anilibria, "GET", "searchTitle?search="+name, nil, "")

	if err != nil {
		helpers.ErrorAction(ctx, err.Error(), 400)
		return
	}

	helpers.SuccessAction(ctx, "Успешно", 200, data)
}
