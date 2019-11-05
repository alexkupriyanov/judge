package Controllers

import (
	"github.com/gorilla/mux"
	"judge/Models"
	"judge/util"
	"net/http"
	"strconv"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	gameTypeId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = rnd.HTML(w, http.StatusOK, "createteam", gameTypeId)
	return
}

func CreateTeamPost(w http.ResponseWriter, r *http.Request) {
	gameTypeId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = r.ParseForm()
	name := r.Form.Get("name")
	res, err := Models.CreateTeam(name, uint(gameTypeId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	http.Redirect(w, r, "/team/"+strconv.Itoa(int(res.Id)), http.StatusFound)
	return
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	teamId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = rnd.HTML(w, http.StatusOK, "createplayer", teamId)
	return
}

func CreatePlayerPost(w http.ResponseWriter, r *http.Request) {
	teamId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = r.ParseForm()
	name := r.Form.Get("name")
	_, err := Models.CreatePlayer(name, uint(teamId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	http.Redirect(w, r, "/team/"+strconv.Itoa(teamId), http.StatusFound)
	return
}
func GetTeam(w http.ResponseWriter, r *http.Request) {
	teamId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	res, err := Models.GetTeam(uint(teamId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	_ = rnd.HTML(w, http.StatusOK, "team", res)
	return
}
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	playerId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	res, err := Models.GetPlayer(uint(playerId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	_ = rnd.HTML(w, http.StatusOK, "player", res)
	return
}
