package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.New("token 过期")
	ErrTokenNotValidYet = errors.New("token 未生效")
	ErrTokenMalformed   = errors.New("这不是一个 token")
	ErrTokenInvalid     = errors.New("无法处理此 token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(Settings.JWTKey.SigningKey),
	}
}

// CreateToken 创建一个token
func CreateToken(id uint, nickName string, role uint) (string, error) {
	j := NewJWT()
	cliams := CustomClaims{
		ID:          id,
		NickName:    nickName,
		AuthorityId: role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + (30 * 60),
			Issuer:    "gin",
		},
	}
	return j.token(cliams)
}

// token 生成一个token
func (j *JWT) token(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if cliams, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return cliams, nil
		}
		return nil, ErrTokenInvalid
	} else {
		return nil, ErrTokenInvalid
	}
}
