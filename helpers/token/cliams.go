package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/config"
	"github.com/saeedhpro/apisimateb/constant"
	"time"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID         uint64 `json:"user_id"`
	OrganizationID uint64 `json:"organization_id"`
	Fname          string `json:"fname"`
	Lname          string `json:"lname"`
	Tel            string `json:"tel"`
	UserGroupID    uint64 `json:"user_group_id"`
}

func (u *UserClaims) GenerateToken() (*string, error) {
	u.ExpiresAt = time.Now().Unix() + constant.ExpTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func GetStaffUser(c *gin.Context) *UserClaims {
	resp, exist := c.Get("claims")
	if !exist {
		return nil
	}
	c2 := resp.(UserClaims)
	return &c2
}
