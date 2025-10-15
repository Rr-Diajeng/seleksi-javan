package security

import (
	"os"
	"seleksi-javan/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GeneratedToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var (
	GenerateToken = generateToken
)

func generateToken(user *model.User) (*GeneratedToken, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	accessTokenDuration, _ := time.ParseDuration(
		os.Getenv("JWT_ACCESS_TOKEN_EXPIRED_AT"),
	)

	refreshTokenDuration, _ := time.ParseDuration(
		os.Getenv("JWT_REFRESH_TOKEN_EXPIRED_AT"),
	)

	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		GenerateClaimsWithExp(user, time.Now().Add(accessTokenDuration).Unix()),
	)

	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		GenerateClaimsWithExp(user, time.Now().Add(refreshTokenDuration).Unix()),
	)

	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &GeneratedToken{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func GenerateClaimsWithExp(user *model.User, exp int64) jwt.MapClaims {
	claims := jwt.MapClaims{
		"iss":     "javan-api",
		"sub":     user.Username,
		"user_id": user.ID,
		"exp":     exp,
	}

	return claims
}
