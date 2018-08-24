package user

import (
	"context"
	"example/Exer_1/golang-training/Exer_2/helper"
	"example/Exer_1/golang-training/Exer_2/model"
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
			if len(bearerToken) == 2 && strings.HasPrefix(author, "Bearer ") {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error")
					}
					return []byte("secretcode"), nil
				})
				if err != nil {
					helper.RespondWithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
					return
				}
				claims, ok := token.Claims.(jwt.MapClaims)
				if ok && token.Valid {
					var user User
					helper.TranDataJson(&claims, &user)
					ctx := context.WithValue(r.Context(), "user", user)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					helper.RespondWithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
					return
				}
			} else {
				helper.RespondWithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
				return
			}
		} else {
			helper.RespondWithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
			return
		}
	})
}
