package firebaseauth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseUID string

type FirebaseAuth struct {
	serviceAccountKeyPath string
	authClient            *auth.Client
}

func NewFirebaseAuth(serviceAccountKeyPath string) *FirebaseAuth {
	return &FirebaseAuth{
		serviceAccountKeyPath: serviceAccountKeyPath,
		authClient:            nil,
	}
}
func (f *FirebaseAuth) Init(serviceAccountKeyPath string) error {
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	f.authClient, err = fb.Auth(context.Background())
	if err != nil {
		return fmt.Errorf("failed init auth client. %s", err)
	}
	return nil
}

func (f *FirebaseAuth) FBAuth(next http.Handler) http.Handler {
	//現在改善中
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		opt := option.WithCredentialsFile(f.serviceAccountKeyPath)
		fmt.Print(opt)
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
	})
}
