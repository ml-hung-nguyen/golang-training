package main

import (
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
					respondwithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorize"})
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					respondwithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorize"})
					return
				}
			}
		} else {
			respondwithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorize"})
			return
		}
	})
}
