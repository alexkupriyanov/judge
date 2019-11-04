package util

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	_ "fmt"
	"log"
	"net/http"
)

type HttpError struct {
	Error interface{}
}

func RandomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func ThrowError(err error, httpStatusCode int, w http.ResponseWriter) {
	log.Println("HTTP ERROR: " + err.Error())
	var result HttpError
	result.Error = err.Error()
	w.WriteHeader(httpStatusCode)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Printf("System error: Can't create object for: %o", result)
	}
	return
}
