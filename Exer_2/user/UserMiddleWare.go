package user

import (
	"context"
	"example/Exer_2/helper"
	"example/Exer_2/model"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")
		if author != "" {
			bearerToken := strings.Split(author, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error")
					}
					return []byte("secretcode"), nil
				})
				if err != nil {
					helper.RespondwithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
					return
				}
				claims, ok := token.Claims.(jwt.MapClaims)
				if ok && token.Valid {
					var user User
					helper.TranDataJson(&claims, &user)
					ctx := context.WithValue(r.Context(), "user", user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					helper.RespondwithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
					return
				}
			}
		} else {
			helper.RespondwithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
			return
		}
	})
}
