package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	. "golang-training/common"
	"golang-training/user"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization != "" && strings.HasPrefix(authorization, "Bearer ") {
			bearerToken := strings.Split(authorization, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error")
					}
					return []byte("somesecretcode"), nil
				})
				if err == nil && token.Valid {
					var user user.User
					claims, _ := token.Claims.(jwt.MapClaims)
					mapstructure.Decode(claims, &user)
					ctx := context.WithValue(r.Context(), "user", user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(ErrorResponse{Message: "Unauthorize"})
					return
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Unauthorize"})
			return
		}
	})
}
