package elocaluclaton

import (
	"strings"
	"testing"
)

func TestInMemoryAddPlayerSucess(t *testing.T) {
	repo := InMemoryPlayerRepository{make(map[string]Player)}
	p := Player{Tag: "Kirbstomper", Elo: 1000}
	err := repo.AddPlayer(p)
	if err != nil {
		t.Fail()
	}
	if repo.players[p.Tag] != p {
		t.Fail()
	}
}

func TestInMemoryAddPlayerDuplicateError(t *testing.T) {
	repo := InMemoryPlayerRepository{make(map[string]Player)}
	p := Player{Tag: "Kirbstomper", Elo: 1000}
	err := repo.AddPlayer(p)
	err = repo.AddPlayer(p)

	if err == nil {
		t.Fail()
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Fail()
	}
}
func TestInMemoryAddPlayerEmptyTagError(t *testing.T) {
	repo := InMemoryPlayerRepository{make(map[string]Player)}
	p := Player{Tag: "", Elo: 1000}
	err := repo.AddPlayer(p)
	if err == nil {
		t.Fail()
	}
	if !strings.Contains(err.Error(), "Cannot Be Empty") {
		t.Fail()
	}
}
func TestInMemoryAddPlayerNegativeEloError(t *testing.T) {
	repo := InMemoryPlayerRepository{make(map[string]Player)}
	p := Player{Tag: "Kirb", Elo: -1000}
	err := repo.AddPlayer(p)

	if err == nil {
		t.Fail()
	}
	if !strings.Contains(err.Error(), "Elo Cannot be lower than 0") {
		t.Fail()
	}
}

func TestGetPlayerSuccess(t *testing.T) {
	repo := InMemoryPlayerRepository{make(map[string]Player)}
	p := Player{Tag: "Kirb", Elo: 1000}
	repo.players[p.Tag] = p

	player, err := repo.GetPlayer("Kirb")

	if err != nil {
		t.Fail()
	}
	if player.Tag != "Kirb" {
		t.Fail()
	}
	if player.Elo != 1000 {
		t.Fail()
	}

}

func TestGetPlayerPlayerDoesNotExistError(t *testing.T) {
	repo := InMemoryPlayerRepository{make(map[string]Player)}

	_, err := repo.GetPlayer("Kirb")
	if err == nil {
		t.Fail()
	}
	if !strings.Contains(err.Error(), "Does Not Exist") {
		t.Fail()
	}
}

func TestGetPlayerEmptyStringError(t *testing.T) {
	repo := InMemoryPlayerRepository{make(map[string]Player)}
	_, err := repo.GetPlayer("")

	if err == nil {
		t.Fail()
	}
	if !strings.Contains(err.Error(), "Cannot Be Empty") {
		t.Fail()
	}
}
