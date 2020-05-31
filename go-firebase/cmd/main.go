package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error verifying ID token %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token\n"))
			return
		}
		log.Printf("Verified ID token: %v\n", token)
		next.ServeHTTP(w, r)
	}
}

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello private!\n"))
}

func main() {
	err := dbInit()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	allowedOrigins := handlers.AllowedOrigins([]string{"http://vue-webapp/"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	r.HandleFunc("/private", authMiddleware(private))
	r.HandleFunc("/books/register", bookRegister)
	r.HandleFunc("/books/list", bookList)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}