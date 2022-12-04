package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lutasam/check_in_sys/biz/common"
	"github.com/lutasam/check_in_sys/biz/model"
)

type JWTStruct struct {
	UserID         uint64 `json:"user_id"`
	Email          string `json:"email"`
	Identity       int    `json:"identity"`
	StandardClaims jwt.StandardClaims
}

func (J JWTStruct) Valid() error {
	return nil
}

// GenerateJWTInUser generates a JWT token by username and user account
func GenerateJWTByUserInfo(user *model.User) (string, error) {
	timeNow := time.Now().Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = JWTStruct{
		UserID:   user.ID,
		Email:    user.Email,
		Identity: user.Identity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeNow + common.EXPIRETIME,
			Issuer:    common.ISSUER,
			NotBefore: timeNow,
		},
	}
	tokenString, err := token.SignedString([]byte(common.OTHERSECRETSALT))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWTToken parse tokenString to JWTStruct
func ParseJWTToken(tokenString string) (*JWTStruct, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTStruct{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(common.OTHERSECRETSALT), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*JWTStruct), nil
}

func GetCtxUserInfoJWT(c *gin.Context) (*JWTStruct, error) {
	jwtStruct, exist := c.Get("jwtStruct")
	if !exist {
		return nil, common.USERNOTLOGIN
	}
	return jwtStruct.(*JWTStruct), nil
}
