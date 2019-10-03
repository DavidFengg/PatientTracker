package login

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Response struct {
	Token string `json:"token"`
	Username string `json:"user"`
}

/* ** TEMP **
	temporary store valid user creds
*/
var users = map[string]string {
	"test": "test",
}

var signingKey = []byte("secret")

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(2 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(creds.Username)

	res := Response{tokenString, creds.Username}

	json.NewEncoder(w).Encode(res)

	// http.SetCookie(w, &http.Cookie {
	// 	Name: "token",
	// 	Value: tokenString,
	// 	Expires: expirationTime,
	// })
}