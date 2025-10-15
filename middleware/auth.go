package middleware

import (
	"net/http"
	"os"
	ucuser "seleksi-javan/usecase/uc_user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

func NewAuthMiddleware(uc ucuser.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(t *jwt.Token) (interface{}, error) {
			secret := os.Getenv("JWT_SECRET_KEY")

			return []byte(secret), nil
		})

		if err != nil {
			UnauthorizedResponse(c)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["user_id"].(float64)

			user, err := uc.GetUserByID(uint(userId))
			if err != nil {
				UnauthorizedResponse(c)
				return
			}

			c.Set("user", user.ID)
		} else {
			UnauthorizedResponse(c)
			return
		}

		c.Next()
	}
}

func UnauthorizedResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"message": "Unauthorized",
	})
}
