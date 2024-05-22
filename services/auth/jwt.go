package authService

import (
	"context"
	"fmt"
	"jwt-api/config"
	userModels "jwt-api/models/user"
	"jwt-api/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string
const UserKey contextKey = "userId"

func CreateJwtToken(payload interface{}) (string, error) {
	secret := config.ENV.JwtSecret
	expirationTime := time.Now().Add(time.Duration(config.ENV.JwtExpirationTime) * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
		"sub": payload,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyJWTToken(handlerFunc http.HandlerFunc, store userModels.UserStore) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
			token := GetTokenFromRequest(r)
			if token == "" {
				utils.WriteJsonError(w, http.StatusUnauthorized, fmt.Errorf("token is required"))
				return
			}
			claims, err := ValidateToken(token)
			if err != nil {
				utils.WriteJsonError(w, http.StatusUnauthorized, fmt.Errorf("token is required"))
				return
			}
			if !claims.Valid {
				utils.WriteJsonError(w, http.StatusUnauthorized, fmt.Errorf("token is required"))
				return
			}

			claimsMap := claims.Claims.(jwt.MapClaims)
			user := claimsMap["sub"].(userModels.User)
			ctx := r.Context()
			ctx = context.WithValue(ctx, UserKey, user.Id)
			r = r.WithContext(ctx)
			handlerFunc(w, r)
	}
}




func GetTokenFromRequest(r * http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth != ""{
		return tokenAuth
	}
	return ""
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := config.ENV.JwtSecret
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}

func GetUserIdFromContext (ctx context.Context) int {
	userId, ok := ctx.Value(UserKey).(int)
	if !ok {
		return 0
	}
	return userId
}