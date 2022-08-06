package services

import (
	database "animerest/Database"
	helpers "animerest/Helpers"
	models "animerest/Models"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	*database.DbContext
}

func (u *UserService) GetAll() {
	u.Db.Model().Select()
}

func (u *UserService) Login(ctx *gin.Context, model models.UserLoginDto) {
	var user models.User

	err := u.Db.Model(&user).
		Where("email = ?", model.Login).
		Select()

	if err != nil {
		pgErr := helpers.ExtractPgException(err.Error(), "login")

		helpers.ErrorAction(ctx, pgErr.Msg, 400)
		return
	}

	isCompare := helpers.ComparePasswords(user.Password, []byte(model.Password))

	if isCompare {
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, models.UserClaims{
			Uid: user.Uid,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: 15000,
			},
		})

		// #TODO get str from config
		token, err := at.SignedString([]byte("Test"))
		if err != nil {
			helpers.ErrorAction(ctx, "Произошла ошибка", 500)
			return
		}

		helpers.SuccessAction(ctx, "Success", 200, token)
	} else {
		helpers.ErrorAction(ctx, "Не правильный логин или пароль", 400)
	}
}

func (u *UserService) Register(ctx *gin.Context, model models.UserRegisterDto) {
	if model.Password != model.ConfirmPassword {
		helpers.ErrorAction(ctx, "Пароли не совпадают", 400)
		return
	}

	err := u.Db.Model(&models.User{}).
		Where("email = ?", model.Email).
		First()

	if err != nil {
		pgErr := helpers.ExtractPgException(err.Error(), "register")

		if pgErr.Code != models.NotFound {
			helpers.ErrorAction(ctx, pgErr.Msg, 400)
			return
		}

		uid, err := u.Db.Model(&models.User{}).Count()
		if err != nil {
			helpers.ErrorAction(ctx, err.Error(), 500)
			return
		}

		uuid, err := uuid.NewV4()
		if err != nil {
			helpers.ErrorAction(ctx, err.Error(), 500)
			return
		}

		bytes, err := bcrypt.GenerateFromPassword([]byte(model.Password), 14)
		if err != nil {
			helpers.ErrorAction(ctx, "Произошла ошибка", 500)
			return
		}

		obj := &models.User{
			Id:         uuid,
			Uid:        uid,
			Username:   string(uid),
			FirstName:  model.FirstName,
			LastName:   model.LastName,
			MiddleName: model.MiddleName,
			Role:       models.Member,
			Rating:     0.0,
			Email:      model.Email,
			Password:   string(bytes),
		}

		_, dbErr := u.Db.Model(obj).Insert()
		if dbErr != nil {
			pgErr := helpers.ExtractPgException(dbErr.Error(), "register")

			helpers.ErrorAction(ctx, pgErr.Msg, 500)
			return
		} else {
			helpers.SuccessAction(ctx, "Пользователь создан", 200, uuid)
		}
	} else {
		helpers.ErrorAction(ctx, "Пользователь уже существует", 400)
	}
}
