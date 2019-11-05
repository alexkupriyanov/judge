package Models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"sort"
)

type GameType struct {
	Id      uint
	Name    string
	Matches []Match
	Events  []EventType
	Teams   []Team
}

func NewGameType(name string) (GameType, error) {
	var existingGameType GameType
	err := GetDB().Where(&GameType{Name: name}).First(&existingGameType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return GameType{}, nil
	} else if err == nil {
		return GameType{}, errors.New("Match type with same name already exist! ")
	}
	gameType := GameType{Name: name}
	err = GetDB().Create(&gameType).Error
	if err != nil {
		return GameType{}, err
	}
	return gameType, nil
}

func GetCurrentGameType(gameTypeId int) (GameType, error) {
	var gameType GameType
	err := GetDB().Preload("Matches.Team1").Preload("Matches.Team2").Preload("Events").Preload("Teams").First(&gameType, gameTypeId).Error
	if err != nil {
		return GameType{}, err
	}
	return gameType, nil
}

func GetGameTypeList() ([]GameType, error) {
	var types []GameType
	err := GetDB().Find(&types).Error
	if err != nil {
		return nil, err
	}
	return types, nil
}

type Match struct {
	Id         uint
	GameType   GameType
	GameTypeId uint
	Team1Id    uint
	Team1      Team
	Team2Id    uint
	Team2      Team
	Events     []Event
}

type MatchViewModel struct {
	GameTypeId uint
	Teams      []Team
}

func GetMatchViewModel(gameTypeId uint) (MatchViewModel, error) {
	var matchViewModel MatchViewModel
	var teams []Team
	err := GetDB().Where(&Team{GameTypeId: gameTypeId}).Find(&teams).Error
	if err != nil {
		return MatchViewModel{}, err
	}
	matchViewModel.GameTypeId = gameTypeId
	matchViewModel.Teams = teams
	return matchViewModel, nil
}

func NewMatch(gameTypeId uint, team1Id uint, team2Id uint) (Match, error) {
	game := Match{GameTypeId: gameTypeId, Team1Id: team1Id, Team2Id: team2Id}
	err := GetDB().Create(&game).Error
	if err != nil {
		return Match{}, err
	}
	return game, nil
}

func GetMatch(matchId uint) (Match, error) {
	var match Match
	err := GetDB().Preload("Team1.Players").Preload("Team2.Players").Preload("Events.EventType").Preload("Events.Player1").Preload("Events.Player2").First(&match, matchId).Error
	sort.Slice(match.Events, func(i, j int) bool {
		return match.Events[i].MinutesBeforeStarted < match.Events[j].MinutesBeforeStarted
	})
	if err != nil {
		return Match{}, err
	}
	return match, nil
}
