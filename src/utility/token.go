package utility

import (
	"github.com/dgrijalva/jwt-go"
	"gopost/src/base"
	"time"
)

func ParseToken(raw string) (string, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, NewError(ERROR_AUTH_CODE, ERROR_MSG_ERROR_SIGN_METHOD)
		}
		return []byte(base.GetConfig().Token.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var username string
		var ok1 bool
		if username, ok1 = claims["username"].(string); !ok1 {
			return "", NewError(ERROR_CONVER_CODE, ERROR_MSG_AUTH_TOKEN_USERNAME_ERROR)
		}
		//貌似自带超时检测
		//var exp int64
		//if exp, ok1 = claims["exp"].(int64); !ok1 {
		//	return "", NewError(ERROR_CONVER_CODE, ERROR_MSG_AUTH_TOKEN_EXP_ERROR)
		//}
		//if time.Unix(exp, 0).After(time.Now()) {
		//	return "", NewError(ERROR_AUTH_CODE, ERROR_MSG_AUTH_TIMEOUT)
		//}
		return username, nil
	}
	return "", NewError(ERROR_CONVER_CODE, ERROR_MSG_AUTH_TOKEN_KNOW_ERROR)
}

func GenerateToken(username string) (res string, err error) {
	tconfig := base.GetConfig().Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	exp := time.Now().Add(time.Second * time.Duration(tconfig.TTL))
	claims["username"] = username
	claims["exp"] = exp.Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if res, err = token.SignedString([]byte(tconfig.SecretKey)); err != nil {
		return "", Wrap(err, ERROR_AUTH_CODE, ERROR_MSG_ERROR_TOKEN_SIGN)
	}
	return res, nil
}
