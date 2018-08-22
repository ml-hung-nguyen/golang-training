package middleware

import (
	"baitapgo_ngay1/golang-training/model"
	user "baitapgo_ngay1/golang-training/user"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("middleware")
		authorization := r.Header.Get("Authorization")
		if authorization != "" {
			bearerToken := strings.Split(authorization, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error")
					}
					return []byte("somesecretcode"), nil
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(model.MessageResponse{Message: "Unauthorize"})
					return
				}
				if token.Valid {
					//fmt.Println("Valid")
					var u user.User
					claims, _ := token.Claims.(jwt.MapClaims)
					mapstructure.Decode(claims, &u)
					//fmt.Println(u)
					ctx := context.WithValue(r.Context(), "user", u)
					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(model.MessageResponse{Message: "Unauthorize"})
					return
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.MessageResponse{Message: "Unauthorize"})
			return
		}
	})
}
