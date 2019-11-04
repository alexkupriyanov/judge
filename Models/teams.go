package Models

type Player struct {
	Id     uint
	Name   string
	TeamId uint
	Team   Team
}

type Team struct {
	Id      uint
	Players []Player
}
