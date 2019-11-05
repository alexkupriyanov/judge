package Models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Player struct {
	Id         uint
	Name       string
	TeamId     uint
	Team       Team
	Statistics []Statistic
}

type Team struct {
	Id         uint
	Name       string
	Players    []Player
	GameType   GameType
	Matches    []Match
	GameTypeId uint
}

type Statistic struct {
	EventType EventType
	Value     int
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
	var eventTypes []EventType
	err = GetDB().Where(&EventType{GameTypeId: player.Team.GameTypeId}).Find(&eventTypes).Error
	if err != nil {
		return player, nil
	}
	for _, eventType := range eventTypes {
		player.AddStatistic(eventType.Name)
	}
	return player, err
}

func (player *Player) AddStatistic(eventTypeName string) {
	var eventType EventType
	err := GetDB().Where(EventType{Name: eventTypeName}).First(&eventType).Error
	if err != nil {
		return
	}
	var events []Event
	var eventCount int
	err = GetDB().Preload("EventType").Where(&Event{EventTypeId: eventType.Id, Player1Id: player.Id}).Find(&events).Count(&eventCount).Error
	if err != nil || eventCount == 0 {
		return
	}
	player.Statistics = append(player.Statistics, Statistic{EventType: eventType, Value: eventCount})
}
