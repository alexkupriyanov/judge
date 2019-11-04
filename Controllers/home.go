package Controllers

import (
	"github.com/thedevsaddam/renderer"
	"net/http"
)

var rnd *renderer.Render
func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./static/*.html",
	}

	rnd = renderer.New(opts)
}

func Home(w http.ResponseWriter, r *http.Request) {
	_ = rnd.HTML(w, http.StatusOK, "home", nil)
}
