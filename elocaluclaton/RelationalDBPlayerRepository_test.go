package elocaluclaton

import (
	"os"
	"testing"
)

func Test_CreateRepositorySuccess(t *testing.T) {
	createRelationalDBPlayerRepository("data")
	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}

func Test_AddPlayerSuccess(t *testing.T) {

	r := createRelationalDBPlayerRepository("data")

	err := r.AddPlayer(Player{Tag: "Kirby", Elo: 1000})
	if err != nil {
		println(err.Error())
		t.Fail()
	}

	var tag string
	var elo int

	err = r.db.ctx.QueryRow("SELECT tag,elo FROM players WHERE tag=?", "Kirby").Scan(&tag, &elo)
	if err != nil {
		t.Fail()
	}
	if tag != "Kirby" {
		t.Fail()
	}
	if elo != 1000 {
		t.Fail()
	}

	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}

func Test_GetPlayerSuccess(t *testing.T) {
	r := createRelationalDBPlayerRepository("data")

	_, err := r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Kirb", 1000)
	if err != nil {
		println(err)
		t.Fail()
	}
	p, err := r.GetPlayer("Kirb")

	if err != nil {
		println(err)
		t.Fail()
	}
	if p.Tag != "Kirb" {
		t.Fail()
	}
	if p.Elo != 1000 {
		t.Fail()
	}
	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}

func Test_UpdatePlayerEloSuccess(t *testing.T) {
	r := createRelationalDBPlayerRepository("data")
	_, err := r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Kirby", 1000)
	if err != nil {
		t.Fail()
	}

	err = r.UpdatePlayerElo("Kirby", 700)
	if err != nil {
		t.Fail()
	}
	var elo int
	err = r.db.ctx.QueryRow("SELECT elo FROM players WHERE tag=?", "Kirby").Scan(&elo)
	if elo != 700 {
		t.Fail()
	}

	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}

func Test_GetAllPlayers_Success(t *testing.T) {
	r := createRelationalDBPlayerRepository("data")
	_, err := r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Kirby", 1000)
	if err != nil {
		t.Fail()
	}
	_, err = r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Paul", 6969)
	if err != nil {
		t.Fail()
	}
	_, err = r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Alex", 420)
	if err != nil {
		t.Fail()
	}

	players, err := r.GetAllPlayers()

	if players == nil {
		t.Fail()
	}
	if len(players) != 3 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}
func Test_GetAllPlayers_ReturnsSorted_Success(t *testing.T) {
	r := createRelationalDBPlayerRepository("data")
	_, err := r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Kirby", 1000)
	if err != nil {
		t.Fail()
	}
	_, err = r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Paul", 6969)
	if err != nil {
		t.Fail()
	}
	_, err = r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Alex", 420)
	if err != nil {
		t.Fail()
	}

	players, err := r.GetAllPlayers()

	if players == nil {
		t.Fail()
	}
	if len(players) != 3 {
		t.Fail()
	}
	if players[0].Tag != "Paul" {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}
func Test_GetAllPlayers_WhenEmpty(t *testing.T) {

	r := createRelationalDBPlayerRepository("data")

	players, err := r.GetAllPlayers()

	if players == nil {
		t.Fail()
	}
	if len(players) != 0 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}
func Test_DeletePlayer_Success(t *testing.T) {

	r := createRelationalDBPlayerRepository("data")

	_, err := r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Kirby", 1000)
	if err != nil {
		t.Fail()
	}

	err = r.DeletePlayer("Kirby")
	if err != nil {
		t.Fail()
	}

	var tag string
	err = r.db.ctx.QueryRow("SELECT tag FROM players WHERE tag=?", "Kirby").Scan(&tag)
	if tag != "" {
		t.Fail()
	}

}
func Test_DeletePlayer_PlayerDoesNotExist(t *testing.T) {

	r := createRelationalDBPlayerRepository("data")

	_, err := r.db.ctx.Exec(`INSERT INTO players(tag, elo) values(?,?)`, "Kirby", 1000)
	if err != nil {
		t.Fail()
	}

	err = r.DeletePlayer("John")
	if err != nil {
		t.Fail()
	}

}
