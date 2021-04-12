package token

import "github.com/dgrijalva/jwt-go"

func GenToken(iat int64, secretKey string, payloads map[string]interface{}, expireSecends int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expireSecends
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
