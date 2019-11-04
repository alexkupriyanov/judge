package main

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"judge/Models"
	"judge/util"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = os.Getenv("Secret")
	store = sessions.NewCookieStore([]byte(key))
)

var Authentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		//check if request does not need authentication, serve the request if it doesn't need it
		session, _ := store.Get(r, "token")
		tokenPart := session.Values["token"].(string) //Grab the token part, what we are truly interested in
		var token Models.Token
		err := Models.GetDB().Where(Models.Token{Token: tokenPart}).First(&token).Error
		if err != nil {
			log.Println(err) //Malformed token, returns with http code 403 as usual
			util.ThrowError(errors.New("Malformed authentication token"), http.StatusUnauthorized, w)
			return
		}

		log.Printf("Requested by user %d", token.UserId)
		if token.ExpiredAt.UTC().Unix() < time.Now().UTC().Unix() {
				log.Printf("Token for user %s is expired ", token.User.Name)
				util.ThrowError(errors.New("Expired token"), http.StatusUnauthorized, w)
				return
		}
		ctx := context.WithValue(r.Context(), "UserId", token.UserId)
		ctx = context.WithValue(ctx, "Token", token.Token)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

var Logger = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		start := time.Now()
		if r.URL.Path == "/api/queue" {
			next.ServeHTTP(w, r)
			return
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

	})
}
