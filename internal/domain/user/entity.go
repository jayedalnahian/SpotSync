package user

import (


	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     string `gorm:"type:varchar(20);not null;default:'driver';check:role IN ('driver', 'admin')" json:"role"`
}


func (u *User) hashPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
	
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) checkPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}