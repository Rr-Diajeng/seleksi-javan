package middleware

import (
	"net/http"
	"seleksi-javan/util/http/errors"

	"github.com/gin-gonic/gin"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if httpError, ok := err.(*errors.HttpError); ok {
				c.JSON(httpError.Code, gin.H{
					"success": false,
					"message": httpError.Message,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Something went wrong",
				})
			}

			c.Abort()
		}
	}
}
