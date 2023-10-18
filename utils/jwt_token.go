package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	data "web_server_2.0/data/mysql"
	"web_server_2.0/types"
)

func CreateToken(user data.SysUser, key string) (string, error) {
	claims := types.JWT{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     //生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			Issuer:    "web_server_2.0",                                   //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//return token.SignedString(MySecret)
	return token.SignedString([]byte(key))
}

// 作用：验证token是否有效

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check if the signing method is HMAC
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// 作用：从token中提取用户ID

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check if the signing method is HMAC
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	return claims["id"].(string), nil
}
