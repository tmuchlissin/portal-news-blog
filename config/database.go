package config

import (
	"fmt"
	"portal-news-blog/database/seed"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PsqlDB.User,
		cfg.PsqlDB.Password,
		cfg.PsqlDB.Host,
		cfg.PsqlDB.Port,
		cfg.PsqlDB.DBName,
		cfg.PsqlDB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Str("host", cfg.PsqlDB.Host).Msg("failed to connect to Postgres")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("failed to get sql.DB from gorm.DB")
		return nil, err
	}

	seed.SeedUsers(db)

	sqlDB.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdle)

	return &Postgres{DB: db}, nil
}
