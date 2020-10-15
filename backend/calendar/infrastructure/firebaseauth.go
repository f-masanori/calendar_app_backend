package firebaseauth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseUID string

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		opt := option.WithCredentialsFile("/go/src/golang/calendarapp-dcbd5-firebase-adminsdk-9fhv7-acb24a0067.json")
		// fmt.Print(opt)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Errorf("error initializing app: %v", err)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		authHeader := r.Header.Get("Authorization")

		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// // JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {

			u := fmt.Sprintf("error verifying ID token: %v\n", err)
			fmt.Print(u)
		}
		// fmt.Print(token)
		uid := token.Claims["user_id"]
		// fmt.Println(uid)
		FirebaseUID = uid.(string)

		next.ServeHTTP(w, r)
	}
}
