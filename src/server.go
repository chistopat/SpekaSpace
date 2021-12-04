package speka_space

import (
	"context"

	entSql "entgo.io/ent/dialect/sql"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent"
	entMigrate "git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/migrate"
	"github.com/rs/zerolog"
)

func NewDBClient(drv *entSql.Driver, log *zerolog.Logger, debug bool) (client *ent.Client, err error) {
	client = ent.NewClient(ent.Driver(drv))
	if debug {
		err = client.Debug().Schema.Create(
			context.Background(),
			entMigrate.WithDropColumn(true),
		)
	} else {
		err = client.Schema.Create(
			context.Background(),
			entMigrate.WithDropColumn(true),
		)
	}
	if err != nil {
		log.Error().Err(err).Msg(err.Error())

		return nil, err
	}

	return client, nil
}
