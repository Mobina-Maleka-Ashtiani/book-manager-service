package DataAccess

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gender string

const (
	Female      Gender = "female"
	Male        Gender = "male"
	NonBinary   Gender = "non-binary"
	Transgender Gender = "transgender"
	Intersex    Gender = "intersex"
	Other       Gender = "i prefer not to say"
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
	Books       []Book `gorm:"foreignKey:UserId"`
}

type Author struct {
	gorm.Model
	FirstName   string `gorm:"varchar(50)" json:"first_name"`
	LastName    string `gorm:"varchar(50)" json:"last_name"`
	Birthday    string `gorm:"varchar(50)" json:"birthday"`
	Nationality string `gorm:"varchar(50)" json:"nationality"`
}

type Book struct {
	gorm.Model
	Name            string `gorm:"varchar(50)" json:"name"`
	Author          Author `gorm:"foreignKey:AuthorID" json:"author"`
	AuthorID        uint   `gorm:"int" json:"author_id"`
	Category        string `gorm:"varchar(20)" json:"category"`
	Volume          int    `gorm:"int" json:"volume"`
	PublishedAt     string `gorm:"varchar(40)" json:"published_at"`
	Summary         string `gorm:"varchar(100)" json:"summary"`
	TableOfContents string `gorm:"varchar(50)" json:"table_of_contents"`
	Publisher       string `gorm:"varchar(50)" json:"publisher"`
	UserId          uint   `json:"user_id"`
}

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
	err := gdb.db.AutoMigrate(&User{}, &Author{}, &Book{})
	if err != nil {
		return err
	}

	return nil
}
