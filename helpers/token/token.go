package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/saeedhpro/apisimateb/config"
	"github.com/saeedhpro/apisimateb/constant"
	"github.com/saeedhpro/apisimateb/domain/models"
	"github.com/saeedhpro/apisimateb/domain/responses"
	"strconv"
	"time"
)

func ValidateToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return nil, errors.New(constant.InvalidTokenError)
	}
	tokenClaims := token.Claims.(*UserClaims)
	if tokenClaims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New(constant.TokenIsExpiredError)
	}
	return tokenClaims, nil
}

func GenerateToken(user *models.UserModel) (*responses.UserLoginResponse, error) {
	claims := UserClaims{
		UserID:         user.ID,
		Tel:            user.Tel,
		Fname:       user.Fname,
		Lname:      user.Lname,
		UserGroupID:    user.UserGroupID,
		OrganizationID: user.OrganizationID,
	}
	claims.ExpiresAt = time.Now().Unix() + constant.ExpTime
	claims.Issuer = strconv.Itoa(int(user.ID))
	token, err := claims.GenerateToken()
	if err != nil {
		return nil, err
	}
	response := responses.UserLoginResponse{
		Token:     *token,
		ExpiresIn: claims.ExpiresAt,
		User:      user,
	}
	return &response, nil
}
