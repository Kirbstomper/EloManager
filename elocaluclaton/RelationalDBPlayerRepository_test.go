package elocaluclaton

import (
	"os"
	"testing"
)

func Test_CreateRepositorySuccess(t *testing.T) {
	createRelationalDBPlayerRepository()
	t.Cleanup(func() {
		os.RemoveAll("data")
	})
}

func Test_AddPlayerSuccess(t *testing.T) {

	r := createRelationalDBPlayerRepository()

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