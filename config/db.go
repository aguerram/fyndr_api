package config

import (
	"context"
	"github.com/phuslu/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewDatabaseConnection(env *AppEnv) (*gorm.DB, func(ctx context.Context)) {
	db, err := gorm.Open(postgres.Open(env.DbDsn), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect to database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic().Err(err).Msg("failed to get sql.DB")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, func(ctx context.Context) {
		if err := sqlDB.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close database connection")
		}
	}
}
