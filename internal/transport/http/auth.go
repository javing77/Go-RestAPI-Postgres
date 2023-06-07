package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func JWTAuth(
	orginal func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "no authorized", http.StatusUnauthorized)
			return
		}

		//Bearer: Token-string
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "no authorized", http.StatusUnauthorized)
			return
		}

		if validateToken(authHeaderParts[1]) {
			orginal(w, r)
		} else {
			http.Error(w, "no authorized", http.StatusUnauthorized)
			return
		}

	}
}

func validateToken(accesToken string) bool {
	var mySignuingKey = []byte("missionimpossible")
	token, err := jwt.Parse(accesToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not vadlidate auth token")
		}
		return mySignuingKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid

}
