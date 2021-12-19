package elocaluclaton

import "errors"

/*
	Implementing the Player repository using a simple map<String, Player>
*/
type InMemoryPlayerRepository struct {
	players map[string]Player
}

func (r InMemoryPlayerRepository) AddPlayer(p Player) error {
	if p.Tag == "" {
		return errors.New("Player Name Cannot Be Empty")
	}
	if p.Elo < 0 {
		return errors.New("Player Elo Cannot be lower than 0")
	}
	if r.players[p.Tag].Tag != "" {
		return errors.New("Player with this tag[" + p.Tag + "] already exists")
	}
	r.players[p.Tag] = p
	return nil
}

func (r InMemoryPlayerRepository) GetPlayer(tag string) (Player, error) {
	p := r.players[tag]

	if tag == "" {
		return p, errors.New("Player Name Cannot Be Empty")
	}
	if p.Tag == "" {
		return p, errors.New("Player with tag[" + tag + "] Does Not Exist")
	}
	return p, nil
}