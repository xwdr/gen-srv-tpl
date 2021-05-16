package tmpl

var User = `package models

import (
	"time"
)

// User
type User struct {
	Id        int64      {{.Quote}}gorm:"primary_key;column:id"{{.Quote}}
	Name      string     {{.Quote}}gorm:"not null;column:name"{{.Quote}}
	Address   string     {{.Quote}}gorm:"not null;column:address"{{.Quote}}
	Age       uint8      {{.Quote}}gorm:"not null;column:age"{{.Quote}}
	CreatedAt time.Time  {{.Quote}}gorm:"column:created_at"{{.Quote}}
	UpdatedAt time.Time  {{.Quote}}gorm:"column:updated_at"{{.Quote}}
	DeletedAt *time.Time {{.Quote}}sql:"index" gorm:"deleted_at"{{.Quote}}
}

func NewUser() *User {
	return &User{}
}

func (User) TableName() string {
	return "tbl_user"
}

// create
func (m *User) Create() error {
	if err := db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

// query
func (m *User) Query() error {
	err := db.Where("id = ?", m.Id).First(m).Error
	return err
}

`

var Models = `package models

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"127.0.0.1:3306",
		"test"))
	if err != nil {
		logrus.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&User{})
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

`