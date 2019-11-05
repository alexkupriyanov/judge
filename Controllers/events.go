package Controllers

import (
	"github.com/gorilla/mux"
	"judge/Models"
	"judge/util"
	"net/http"
	"strconv"
)

func CreateEventType(w http.ResponseWriter, r *http.Request) {
	gameTypeId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = rnd.HTML(w, http.StatusOK, "createEventType", gameTypeId)
	return
}

func CreateEventTypePost(w http.ResponseWriter, r *http.Request) {
	gameTypeId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = r.ParseForm()
	name := r.Form.Get("name")
	playerCount, _ := strconv.Atoi(r.Form.Get("playerCount"))
	_, err := Models.NewEventType(name, playerCount, uint(gameTypeId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	http.Redirect(w, r, "/game/"+strconv.Itoa(gameTypeId), http.StatusFound)
	return
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	matchId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	res, err := Models.GetEventViewModel(uint(matchId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	_ = rnd.HTML(w, http.StatusOK, "selectEvent", res)
	return
}

func CreateEventPost(w http.ResponseWriter, r *http.Request) {
	matchId, _ := strconv.Atoi(mux.Vars(r)["Id"])
	_ = r.ParseForm()
	eventId, _ := strconv.Atoi(r.Form.Get("eventType"))
	http.Redirect(w, r, "/match/"+strconv.Itoa(matchId)+"/event/"+strconv.Itoa(eventId), http.StatusFound)
	return
}

func CreateCurrentEvent(w http.ResponseWriter, r *http.Request) {
	matchId, _ := strconv.Atoi(mux.Vars(r)["MatchId"])
	eventId, _ := strconv.Atoi(mux.Vars(r)["EventId"])
	res, err := Models.GetEventData(uint(matchId), uint(eventId))
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	_ = rnd.HTML(w, http.StatusOK, "createEvent", res)
	return
}

func CreateCurrentEventPost(w http.ResponseWriter, r *http.Request) {
	matchId, _ := strconv.Atoi(mux.Vars(r)["MatchId"])
	eventId, _ := strconv.Atoi(mux.Vars(r)["EventId"])
	creatorId, _ := r.Context().Value("UserId").(uint)
	_ = r.ParseForm()
	playersCount, _ := strconv.Atoi(r.Form.Get("playersCount"))
	player1Id, _ := strconv.Atoi(r.Form.Get("player1"))
	player2Id, err := strconv.Atoi(r.Form.Get("player2"))
	timeBefore, err := strconv.Atoi(r.Form.Get("timeBefore"))
	if err != nil {
		_, err = Models.NewEvent(uint(eventId), creatorId, uint(matchId), uint(playersCount), uint(player1Id), nil, timeBefore)
	} else {
		val := uint(player2Id)
		_, err = Models.NewEvent(uint(eventId), creatorId, uint(matchId), uint(playersCount), uint(player1Id), &val, timeBefore)
	}
	if err != nil {
		util.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	http.Redirect(w, r, "/match/"+strconv.Itoa(matchId), http.StatusFound)
	return
}
