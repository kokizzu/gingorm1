package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserConn struct {
	*gorm.DB
}

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *UserConn) GetUserByEmail(email string) (user *User, err error) {
	err = u.DB.Where("email = ?", email).First(&user).Error
	return
}

func (u *UserConn) InsertUser(email, password string) (err error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &User{
		Email:    email,
		Password: string(hashedPass),
	}
	err = u.DB.Create(user).Error
	return
}
