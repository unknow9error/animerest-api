package helpers

import (
	models "animerest/Models"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PgException struct {
	Msg  string
	Code models.PgExceptionStatusCode
}

func ExtractPgException(msg string, t string) *PgException {
	result := &PgException{
		Msg:  "",
		Code: ExtractPgExceptionCode(msg),
	}

	switch t {
	case "login":
		if result.Code == models.NotFound {
			result.Msg = "Не правильный логин или пароль"
		}
	}

	return result
}

func ExtractPgExceptionCode(msg string) models.PgExceptionStatusCode {
	if strings.Contains(msg, "pg: no rows in result set") {
		return models.NotFound
	}

	if strings.Contains(msg, "pg: Model(unsupported") {
		return models.UnsupportedModel
	}

	return models.Unknown
}

func ErrorAction(ctx *gin.Context, msg string, statusCode int) {
	localizedError := msg

	if strings.Contains(localizedError, "json: Unmarshal") {
		str := strings.Split(strings.Split(localizedError, "models.")[1], ")")[0]
		localizedError = "Не удалось переобразовать модель " + str
	}

	obj := &models.HttpError{
		Reason:     localizedError,
		Status:     nil,
		StatusCode: statusCode,
	}

	ctx.JSON(statusCode, obj)
}

func SuccessAction(ctx *gin.Context, msg string, statusCode int, data interface{}) {
	obj := &models.HttpSuccess{
		Message:    msg,
		StatusCode: statusCode,
		Status:     nil,
		Data:       data,
	}

	ctx.JSON(statusCode, obj)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	} else {
		return true
	}
}

func DecodeJson(body io.ReadCloser, model interface{}) error {
	decoder := json.NewDecoder(body)

	defer body.Close()

	if err := decoder.Decode(model); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
