package mService

import (
	"WangYiYunDemo/music/mService/msDef"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/gbrlsnchs/jwt/v3"
	uuid "github.com/satori/go.uuid"
	"time"
)

// 签名算法, 随机, 不保存密钥, 每次都是随机的
var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
var publicKey = &privateKey.PublicKey
var hs = jwt.NewES256(
	jwt.ECDSAPublicKey(publicKey),
	jwt.ECDSAPrivateKey(privateKey),
)

// Sign 签名
func Sign(id uint, username string) (string, error) {
	now := time.Now()
	pl := msDef.LoginToken{
		Payload: jwt.Payload{
			Issuer:         "lian_xu",
			Subject:        "login",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Second)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.NewV4().String(),
		},
		ID:       id,
		Username: username,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

// Verify 验证
func Verify(token []byte) (*msDef.LoginToken, error) {
	pl := &msDef.LoginToken{}
	_, err := jwt.Verify(token, hs, pl)
	return pl, err
}
