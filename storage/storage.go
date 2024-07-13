package storage

import (
	"fmt"
	"github.com/dilly3/houdini/config"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Storage object
type Storage struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

var DefaultStorage *Storage

func New(config *config.Configuration, logger *zerolog.Logger) *Storage {

	postgresDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=%s",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresDB,
		config.PostgresPassword,
		config.PostgresTimezone,
	)
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		logger.Error().Err(err).Msg("failed to connect to db")

	}
	logger.Info().Msg("connected to db")
	str := &Storage{
		DB:     db,
		Logger: logger,
	}
	if DefaultStorage == nil {
		DefaultStorage = str
	}
	return str

}
