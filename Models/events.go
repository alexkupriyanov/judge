package Models

type EventType struct {
	Id          uint
	Name        string
	PlayerCount int
	GameType    GameType
	GameTypeId  uint
}

func NewEventType(name string, playerCount int, gameTypeId uint) (EventType, error) {
	eventType := EventType{
		Name:        name,
		PlayerCount: playerCount,
		GameTypeId:  gameTypeId,
	}
	err := GetDB().Create(&eventType).Error
	if err != nil {
		return EventType{}, err
	}
	return eventType, nil
}

type Event struct {
	Id                   uint
	EventTypeId          uint
	EventType            EventType
	CreatorId            uint
	Creator              User
	MatchId              uint
	Match                Match
	PlayerCount          uint
	Player1Id            uint
	Player1              Player
	Player2Id            *uint
	Player2              *Player
	MinutesBeforeStarted int
}

type EventViewModel struct {
	MatchId    uint
	Team1Name  string
	Team2Name  string
	EventTypes []EventType
}
type CurrentEventViewModel struct {
	MatchId      uint
	EventName    string
	EventId      uint
	Team1Name    string
	Team2Name    string
	PlayersCount int
	Players      []Player
}

func GetEventViewModel(matchId uint) (EventViewModel, error) {
	var match Match
	err := GetDB().First(&match, matchId).Error
	if err != nil {
		return EventViewModel{}, err
	}
	var eventViewModel EventViewModel
	err = GetDB().Where(&EventType{GameTypeId: match.GameTypeId}).Find(&eventViewModel.EventTypes).Error
	if err != nil {
		return EventViewModel{}, err
	}
	eventViewModel.Team1Name = match.Team1.Name
	eventViewModel.Team2Name = match.Team2.Name
	eventViewModel.MatchId = match.Id
	return eventViewModel, nil
}
func GetEventData(matchId uint, eventTypeId uint) (CurrentEventViewModel, error) {
	var match Match
	err := GetDB().Preload("Team1.Players.Team").Preload("Team2.Players.Team").First(&match, matchId).Error
	if err != nil {
		return CurrentEventViewModel{}, err
	}
	var event EventType
	err = GetDB().First(&event, eventTypeId).Error
	if err != nil {
		return CurrentEventViewModel{}, err
	}
	res := CurrentEventViewModel{
		MatchId:      match.Id,
		EventName:    event.Name,
		EventId:      event.Id,
		Team1Name:    match.Team1.Name,
		Team2Name:    match.Team2.Name,
		PlayersCount: event.PlayerCount,
		Players:      append(match.Team1.Players, match.Team2.Players...),
	}
	return res, nil
}

func NewEvent(eventTypeId uint, creatorId uint, matchId uint, playerCount uint, player1Id uint, player2Id *uint, timeBefore int) (Event, error) {
	event := Event{
		EventTypeId:          eventTypeId,
		CreatorId:            creatorId,
		MatchId:              matchId,
		PlayerCount:          playerCount,
		Player1Id:            player1Id,
		Player2Id:            player2Id,
		MinutesBeforeStarted: timeBefore,
	}
	err := GetDB().Create(&event).Error
	if err != nil {
		return Event{}, err
	}
	return event, nil
}
