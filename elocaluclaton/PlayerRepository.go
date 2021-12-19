package elocaluclaton

type Player struct {
	Tag string
	Elo int
}

//
// - A game is a player repository
// Add player to Game
// - A player is simply a Tag and Elo
// Query a player for a game given a Tag
// Update Elo for player in given game
// Delete Player from Game

type PlayerRepository interface {
	AddPlayer(Player) error
	GetPlayer(string) (Player, error)
	UpdatePlayerElo(string, int) error
	DeletePlayer(string) error
}
