package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"strings"

	"../models"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Usuario models.Usuario     `json:"usuario"`
	Menu    []models.MenuModel `json:"menu"`
	jwt.StandardClaims
}

var jwtKey = []byte("Un4M4s@cr33sL4M3j0r0pc!0n2")

func Create_JWT(usuario models.Usuario) (string, error) {

	expirationTime := time.Now().Add(480 * time.Minute)
	OpcionesMenu, _ := GetOpcionesDeMenu(usuario.IdColaborador)

	claims := &Claims{
		Usuario: usuario,
		Menu:    OpcionesMenu,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("Credenciales incorrectas")
	}

	return tokenString, nil
}

func IsLogginMiddleWare(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization, X-Requested-With, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Screen")

		notAuth := []string{"/prueba", "/login", "/loginWithToken"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		reqToken := r.Header.Get("Authorization")

		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "No se ha proporcionado el token")
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "No autenticado")
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "No autenticado")
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := jwtKey // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func GetToken(r *http.Request) (Claims, bool) {

	var result Claims

	reqToken := r.Header.Get("Authorization")

	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return result, true
	}

	reqToken = strings.TrimSpace(splitToken[1])

	token, _ := ExtractClaims(reqToken)

	data, _ := json.Marshal(token)

	json.Unmarshal(data, &result)
	return result, false
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
