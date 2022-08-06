package database

import (
	config "animerest/Config"
	models "animerest/Models"
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DbContext struct {
	Db *pg.DB
}

func Connect() *DbContext {
	data := config.GetConfig()

	postgres := data.Connection.Postgres
	db := pg.Connect(&pg.Options{
		Password: postgres.Password,
		Database: postgres.Database,
		User:     postgres.User,
		Addr:     postgres.Port,
	})

	ctx := db.Context()
	Up(ctx, db)

	return &DbContext{Db: db}
}

func Up(ctx context.Context, db *pg.DB) {
	models := []interface{}{
		(*models.User)(nil),
		(*models.UserAchievement)(nil),
		(*models.UserChat)(nil),
		(*models.UserFriend)(nil),
		(*models.UserFavorite)(nil),
		(*models.UserMessage)(nil),
		(*models.UserWatchLater)(nil),
		(*models.Anime)(nil),
	}

	err := db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		for _, model := range models {
			err := db.Model(model).CreateTable(&orm.CreateTableOptions{
				IfNotExists: true,
			})

			if err != nil {
				panic(err.Error())
			}
		}

		return nil
	})

	if err != nil {
		panic(err.Error())
	}
}
