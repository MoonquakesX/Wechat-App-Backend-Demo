package msDef

import (
	"github.com/gbrlsnchs/jwt/v3"
)

// LoginToken 记录登录信息的 JWT
type LoginToken struct {
	jwt.Payload
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type Token struct {
	Token string `json:"token"`
}
