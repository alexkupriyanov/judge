package Controllers

import (
	"github.com/gorilla/mux"
	"judge/Models"
	"judge/util"
	"net/http"
	"strconv"
)

func GameTypeList(w http.ResponseWriter, r *http.Request) {
	res, _ := Models.GetGameTypeList()
	_ = rnd.HTML(w, http.StatusOK, "gametypelist", res)
}
func CreateGameType(w http.ResponseWriter, r *http.Request) {
	_ = rnd.HTML(w, http.StatusOK, "creategametype", nil)
}

func CreateGameTypePost(w http.ResponseWriter, r *http.Request) {
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

func GetCurrentGameType(w http.ResponseWriter, r *http.Request) {
	gameTypeId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	res, err := Models.GetCurrentGameType(gameTypeId)
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	_ = rnd.HTML(w, http.StatusOK, "gametype", res)
}

func CreateMatch(w http.ResponseWriter, r *http.Request) {
	gameTypeId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	res, err := Models.GetMatchViewModel(uint(gameTypeId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	_ = rnd.HTML(w, http.StatusOK, "createMatch", res)
	return
}

func GetMatch(w http.ResponseWriter, r *http.Request) {
	matchId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	res, err := Models.GetMatch(uint(matchId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	_ = rnd.HTML(w, http.StatusOK, "getMatch", res)

	return
}
func CreateMatchPost(w http.ResponseWriter, r *http.Request) {
	gameTypeId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = r.ParseForm()
	Team1Id, _ := strconv.Atoi(r.Form.Get("Team1"))
	Team2Id, _ := strconv.Atoi(r.Form.Get("Team2"))
	_, err := Models.NewMatch(uint(gameTypeId), uint(Team1Id), uint(Team2Id))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	http.Redirect(w, r, "/game/"+strconv.Itoa(gameTypeId), http.StatusFound)
	return
}
