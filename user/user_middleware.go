package user

import (
	"context"
	"fmt"
	"golang-training/helper"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization != "" {
			bearer := strings.Split(authorization, " ")
			if len(bearer) == 2 && strings.HasPrefix(authorization, "Bearer ") {
				log.Println(bearer)
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
				claims, ok := token.Claims.(jwt.MapClaims)
				if ok && token.Valid {
					var user User
					helper.TranDataJson(&claims, &user)
					ctx := context.WithValue(r.Context(), "user", user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
					return
				}
			} else {
				helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
				return
			}
		} else {
			helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			return
		}
	})
}
