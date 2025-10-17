package dbrepo

import (
	"database/sql"

	"github.com/vadim-vep/booking/internal/config"
	"github.com/vadim-vep/booking/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testPostgresRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
func NewTestingPostgresRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testPostgresRepo{
		App: a,
	}
}
