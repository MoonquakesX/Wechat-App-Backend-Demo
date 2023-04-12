package mdDef

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null; column:name"`
	PassWord string `json:"passWord" gorm:"not null; column:password"`
}

// Compare 比较密码是否相同, 前提是 UserModel.Password 是已经哈希过的
func (u *User) Compare(pwd string) error {
	return Compare(u.PassWord, pwd)
}

// Encrypt 加密用户密码
func (u *User) Encrypt() error {
	password, err := Encrypt(u.PassWord)
	if err == nil {
		u.PassWord = password
	}
	return err
}

// Encrypt 加密字符串
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare 比较加密字符串和原始字符串是否相同
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
