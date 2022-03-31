package manager

import (
	"go-wmb/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(dsn string) Infra {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = conn.AutoMigrate(&model.Food{}, &model.Table{}, &model.Order{}, &model.Payment{})
	if err != nil {
		panic(err)
	}
	return &infra{
		db: conn,
	}
}
