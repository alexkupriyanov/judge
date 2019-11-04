package Controllers

import (
	"judge/Models"
	"judge/util"
	"net/http"
)

func GameTypeList(w http.ResponseWriter, r *http.Request) {
	res, _ := Models.GetGameTypeList()
	_ = rnd.HTML(w, http.StatusOK, "gametypelist", res)
}
func CreateGameType(w http.ResponseWriter, r *http.Request) {
	_ = rnd.HTML(w, http.StatusOK, "creategametype", nil)
}
func CreateGameTypePost(w http.ResponseWriter, r *http.Request)  {
	_ = r.ParseForm()
	name := r.Form.Get("name")
	_, err := Models.NewGameType(name)
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	http.Redirect(w, r, "/games", http.StatusFound)
	return
}