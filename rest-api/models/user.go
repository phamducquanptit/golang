package models

import (
	"rest-api/pkg/auth"

	validator "gopkg.in/go-playground/validator.v9"
)

// UserModel - User Table
type UserModel struct {
	BaseModel
	Username string `gorm:"column:username;not null" json:"username" validate:"min=1,max=32" binding: "required"`
	Password string `gorm:"column:password;not null" json:"password" validate:"min=5,max=128" binding: "required"`
}

// TableName - mapping name table user in database
func (u *UserModel) TableName() string {
	return "user"
}

// Create a user
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func FindUser(username string) (*UserModel, bool) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	if d.Error != nil {
		return nil, false
	}
	return u, true
}

func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields of UserModel
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
