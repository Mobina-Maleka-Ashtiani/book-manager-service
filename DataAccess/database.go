package DataAccess

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	cfg Config
	db  *gorm.DB
}

func NewGormDB(cfg Config) (*GormDB, error) {
	c := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Name,
		cfg.Database.Password,
	)

	// Create a new connection
	db, err := gorm.Open(postgres.Open(c), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &GormDB{
		cfg: cfg,
		db:  db,
	}, nil
}
