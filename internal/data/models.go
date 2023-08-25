package data

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies      MovieModel
	Permissions PermissionModel
	Users       UserModel
	Tokens      TokenModel
}

func NewModels(pool *pgxpool.Pool) Models {
	return Models{
		Movies:      MovieModel{DB: pool},
		Permissions: PermissionModel{DB: pool},
		Users:       UserModel{DB: pool},
		Tokens:      TokenModel{DB: pool},
	}
}
