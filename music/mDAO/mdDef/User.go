package mdDef

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null; column:name" validate:"min=1,max=32"`
	PassWord string `json:"passWord" gorm:"not null; column:password" validate:"min=5,max=128"`
}

// Compare 比较密码是否相同, 前提是 UserModel.Password 是已经哈希过的
func (u *UserBasic) Compare(pwd string) error {
	return Compare(u.PassWord, pwd)
}

// Encrypt 加密用户密码
func (u *UserBasic) Encrypt() error {
	password, err := Encrypt(u.PassWord)
	if err == nil {
		u.PassWord = password
	}
	return err
}

// Validate 验证字段
func (u *UserBasic) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
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
