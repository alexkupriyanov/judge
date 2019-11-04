package Models

type EventType struct {
	Id     			uint
	Name   			string
	PlayerNumber	int
	GameType		GameType
	GameTypeId		uint
}

func NewEventType(name string, playerNumber int, gameTypeId uint) (EventType, error) {
	eventType := EventType{
		Name:         name,
		PlayerNumber: playerNumber,
		GameTypeId:   gameTypeId,
	}
	err := GetDB().Create(&eventType).Error
	if err != nil {
		return EventType{}, err
	}
	return eventType, nil
}

func GetEventTypes(gameTypeId uint) ([]EventType, error) {
	var types []EventType
	err := GetDB().Where(EventType{GameTypeId:gameTypeId}).Find(&types).Error
	if err != nil {
		return nil, err
	}
	return types, nil
}


type Event struct {
	Id 			uint
	EventTypeId uint
	EventType 	EventType
	CreatorId 	uint
	Creator 	User
	GameId		*uint
	Game		Game
	PlayerCount uint
	Player1Id	uint
	Player1		Player
	Player2Id	*uint
	Player2		*Player
}

func NewEvent(eventTypeId uint, creatorId uint, gameId *uint, playerCount uint, player1Id uint, player2Id *uint) *Event {
	return &Event{EventTypeId: eventTypeId, CreatorId: creatorId, GameId: gameId, PlayerCount: playerCount, Player1Id: player1Id, Player2Id: player2Id}
}