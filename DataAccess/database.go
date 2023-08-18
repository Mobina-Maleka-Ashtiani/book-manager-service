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

type User struct {
	gorm.Model
	FirstName   string `gorm:"varchar(50)" json:"first_name"`
	LastName    string `gorm:"varchar(50)" json:"last_name"`
	Gender      Gender `gorm:"varchar(25)" json:"gender"`
	PhoneNumber string `gorm:"varchar(30),unique" json:"phone_number"`
	Username    string `gorm:"varchar(50),unique" json:"username"`
	Email       string `gorm:"varchar(50),unique" json:"email"`
	Password    string `gorm:"varchar(64)" json:"password"`
}

type Gender string

const (
	Female      Gender = "female"
	Male        Gender = "male"
	NonBinary   Gender = "non-binary"
	Transgender Gender = "transgender"
	Intersex    Gender = "intersex"
	Other       Gender = "i prefer not to say"
)

func NewGormDB(cfg Config) (*GormDB, error) {
	c := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Name,
		cfg.Database.Password,
	)

	db, err := gorm.Open(postgres.Open(c), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &GormDB{
		cfg: cfg,
		db:  db,
	}, nil
}

func (gdb *GormDB) CreateSchemas() error {
	err := gdb.db.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}
