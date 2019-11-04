package Models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type GameType struct {
	Id 		uint
	Name	string
	Games	[]Game
}

func NewGameType(name string) (GameType, error) {
	var existingGameType GameType
	err := GetDB().Where(&GameType{Name:name}).First(&existingGameType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return GameType{}, nil
	} else if err == nil {
		return GameType{}, errors.New("Game type with same name already exist! ")
	}
	gameType := GameType{Name:name}
	err = GetDB().Create(&gameType).Error
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

type Game struct {
	Id         uint
	GameType   GameType
	GameTypeId uint
	Team1Id    uint
	Team1      Team
	Team1Score uint `gorm:"default:0"`
	Team2Id    uint
	Team2      Team
	Team2Score uint `gorm:"default:0"`
}

func NewGame(gameTypeId uint, team1Id uint, team2Id uint) (Game, error) {
	game := Game{GameTypeId:gameTypeId, Team1Id:team1Id, Team2Id:team2Id}
	err := GetDB().Create(&game).Error
	if err != nil {
		return Game{}, err
	}
	return game, nil
}

func GameList(gameTypeId uint) ([]Game, error) {
	var games []Game
	err := GetDB().Where(Game{GameTypeId:gameTypeId}).Find(&games).Error
	if err != nil {
		return nil, err
	}
	return games, nil
}