package user

import (
	"context"
	"fmt"
	"golang-training/helper"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		// log.Println(authorization)
		if authorization != "" {
			bearer := strings.Split(authorization, " ")
			if len(bearer) == 2 {
				token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error")
					}
					return []byte("secretcode"), nil
				})
				if err != nil {
					helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
					return
				}
				if token.Valid {
					// log.Println(token.Valid)
					var user User
					claims, _ := token.Claims.(jwt.MapClaims)
					mapstructure.Decode(claims, &user)
					ctx := context.WithValue(r.Context(), "user", user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
					return
				}
			}
		} else {
			helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}
	})
}
