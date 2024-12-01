package middleware

import (
	"fmt"
	"net/http"
	"os"
	"rentalMobil/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization") // BEARER eykwaiowoakawjo ['bear', 'eykawoakw']

		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		tokenString = strings.Split(tokenString, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			secret := os.Getenv("SECRET_KEY")
			return []byte(secret), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, nil)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("username", claims["username"])
			ctx.Set("role", claims["role"])
		}

		ctx.Next()
	}
}

func MustAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		if role != models.ROLE_ADMIN {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		ctx.Next()
	}
}
