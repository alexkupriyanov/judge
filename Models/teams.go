package Models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Player struct {
	Id     uint
	Name   string
	TeamId uint
	Team   Team
}

type Team struct {
	Id         uint
	Name       string
	Players    []Player
	GameType   GameType
	Matches    []Match
	GameTypeId uint
}

func CreateTeam(name string, gameTypeId uint) (Team, error) {
	var team Team
	err := GetDB().Where(&Team{Name: name}).First(&team).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Team{}, err
	} else if err == nil {
		return Team{}, errors.New("Team with same name already exist ")
	}
	team = Team{Name: name, GameTypeId: gameTypeId}
	err = GetDB().Create(&team).Error
	if err != nil {
		return Team{}, err
	}
	return team, nil
}
func GetTeam(teamId uint) (Team, error) {
	team := Team{Id: teamId}
	err := GetDB().Preload("Players").Preload("GameType").First(&team).Error
	if err != nil {
		return Team{}, err
	}
	return team, nil
}

func CreatePlayer(name string, teamId uint) (Player, error) {
	player := Player{Name: name, TeamId: teamId}
	err := GetDB().Create(&player).Error
	if err != nil {
		return Player{}, err
	}
	return player, nil
}

func GetPlayer(playerId uint) (Player, error) {
	var player Player
	err := GetDB().Preload("Team").First(&player, playerId).Error
	if err != nil {
		return Player{}, err
	}
	return player, err
}
