package Controllers

import (
	"github.com/gorilla/sessions"
	"judge/Models"
	"judge/util"
	"net/http"
	"os"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = os.Getenv("Secret")
	store = sessions.NewCookieStore([]byte(key))
)

var Register = func(w http.ResponseWriter, r *http.Request) {
	_ = rnd.HTML(w, http.StatusOK, "register", nil)
}
var RegisterPost = func(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "token")
	_ = r.ParseForm()
	login := r.Form.Get("email")
	password := r.Form.Get("password")
	name := r.Form.Get("username")
	res, err := Models.NewUser(login, password, name)
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	session.Values["token"] = res.Token
	_ = session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	_ = rnd.HTML(w, http.StatusOK, "login", nil)
}
var LoginPost = func(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "token")
	_ = r.ParseForm()
	login := r.Form.Get("email")
	password := r.Form.Get("password")
	res, err := Models.Login(login, password)
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	session.Values["token"] = res.Token
	_ = session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
