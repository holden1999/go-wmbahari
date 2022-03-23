package manager

import "github.com/jmoiron/sqlx"

type Infra interface {
	SqlDb() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func (i *infra) SqlDb() *sqlx.DB {
	return i.db
}

func NewInfra(dataSourcename string) Infra {
	conn, err := sqlx.Connect("pgx", dataSourcename)
	if err != nil {
		panic(err)
	}
	return &infra{
		db: conn,
	}
}
