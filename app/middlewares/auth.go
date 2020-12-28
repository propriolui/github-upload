package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var ()

//GenerateJWT : permette di generare un token jwt
func GenerateJWT(userID int) (string, error) {
	//recupero la chiave segreta dal file
	err := godotenv.Load("../../.env")
	if err != nil {
		return "error: ", fmt.Errorf("Error loading env file")
	}
	secretKey := os.Getenv("JWT_Secret")
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

//IsAuthorized : controlla utilizzando il token jwt se la richiesta è autorizzata
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//recupero la chiave segreta dal file
		err := godotenv.Load("../../.env")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		secretKey := os.Getenv("JWT_Secret")
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return secretKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
