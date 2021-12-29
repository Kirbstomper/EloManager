package elocaluclaton

import "errors"

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
	GetAllPlayers() ([]Player, error)
	DeletePlayer(string) error
}

func ValidateAddPlayer(p Player) error {
	if p.Tag == "" {
		return errors.New("Player Name Cannot Be Empty")
	}
	if p.Elo < 0 {
		return errors.New("Player Elo Cannot be lower than 0")
	}

	return nil
}

func ValidateGetPlayer(tag string) error {
	if tag == "" {
		return errors.New("Player Name Cannot Be Empty")
	}
	return nil
}

func ValidateUpdatePlayerElo(newElo int) error {
	if newElo < 0 {
		return errors.New("New Elo Cannot Be Negative")
	}
	return nil
}
