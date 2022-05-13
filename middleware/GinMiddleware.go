package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/saeedhpro/apisimateb/constant"
	"github.com/saeedhpro/apisimateb/helpers/token"
	"net/http"
	"strings"
)

const (
	claimKey = "claims"
)

func GinJwtAuth(function gin.HandlerFunc, selfAccess, optional bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := strings.Replace(context.GetHeader("Authorization"), "Bearer ", "", -1)
		if tokenString == "" && optional {
			context.Set(claimKey, &token.UserClaims{})
			function(context)
			context.Next()
		} else {
			claims, err := token.ValidateToken(tokenString)
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"error": constant.UnAuthorizedError,
				})
				return
			}

			context.Set(claimKey, *claims)
			function(context)
			context.Next()
		}
	}
}
