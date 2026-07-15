package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PsqlDB.User,
		cfg.PsqlDB.Password,
		cfg.PsqlDB.Host,
		cfg.PsqlDB.Port,
		cfg.PsqlDB.DBName,
	)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to Postgres host %s: %v", cfg.PsqlDB.Host, err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to get sql.DB from gorm.DB: %v", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdle)

	return &Postgres{DB: db}, nil
}
