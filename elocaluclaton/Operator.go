package elocaluclaton

import (
	"log"
)

/*
	This class uses both a player repository and elecalculator to manage functions
*/

type Operator struct {
	repo PlayerRepository
}

func CreateInMemoryOperator() Operator {
	return Operator{repo: createInMemoryRepository()}
}
func CreateRealtionalDBOperator() Operator {
	return Operator{repo: createRelationalDBPlayerRepository("/data")}
}

/*
	Creates a new player by adding it to the repository
*/
func (o Operator) CreateNewPlayer(tag string, elo int) error {
	p := Player{Tag: tag, Elo: elo}

	err := o.repo.AddPlayer(p)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Added Player ", p)
	}
	return err
}

/*
	Retrieves a Player from the repository
*/
func (o Operator) RetrievePlayerInformation(tag string) (Player, error) {
	p, err := o.repo.GetPlayer(tag)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("Retrieved Player ", p)
	}
	return p, err
}

/*
	Looks up players, calculates and returns the new elo for the players
*/
func (o Operator) CalculateEloForPlayers(tagA, tagB string, outcome int) (int, int, error) {
	playerA, err := o.repo.GetPlayer(tagA)
	if err != nil {
		log.Println(err.Error())
		return -1, -1, err
	}
	playerB, err := o.repo.GetPlayer(tagB)
	if err != nil {
		log.Println(err.Error())
		return -1, -1, err
	}
	a, b := CalculateElo(playerA.Elo, playerB.Elo, outcome)
	return a, b, err
}

/*
	Updates the elo for a given player if they exist
*/
func (o Operator) UpdateEloForPlayer(tag string, newElo int) error {
	err := o.repo.UpdatePlayerElo(tag, newElo)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Updated Player:", tag, " | Elo:", newElo)
	return err
}

/*
	Updates Elo for Two given players based on outcome of match
*/
func (o Operator) RunMatch(tagA, tagB string, outcome int) error {
	eloA, eloB, err := o.CalculateEloForPlayers(tagA, tagB, outcome)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = o.UpdateEloForPlayer(tagA, eloA)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = o.UpdateEloForPlayer(tagB, eloB)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
