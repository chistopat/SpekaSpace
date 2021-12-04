package repository

import (
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent"

	"github.com/rs/zerolog"
)

type Repository struct {
	Client *ent.Client
	log    *zerolog.Logger
}

func NewRepository(db *ent.Client, logger *zerolog.Logger) Repository {
	return Repository{
		Client: db,
		log:    logger,
	}
}
