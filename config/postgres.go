package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewPostgresConfig(cfg *Config) (*gorm.DB, error) {
	pgConfig := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSL,
	)

	db, err := gorm.Open("postgres", pgConfig)

	if err != nil {
		panic(err)
	}

	return db, nil
}
