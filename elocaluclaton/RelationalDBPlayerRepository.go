package elocaluclaton

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	ctx *sql.DB
}

type RelationalDBPlayerRepository struct {
	db database
}

func createDatabase(path string) *database {
	dbDir := path
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		os.Mkdir(dbDir, os.ModePerm)
	}
	ctx, err := sql.Open("sqlite3", dbDir+"/repository.db")
	if err != nil {
		log.Fatalln(err)
	}
	initStmts := []string{
		`PRAGMA busy_timeout = 5000`,
		`PRAGMA synchronous = NORMAL`,

		`CREATE TABLE IF NOT EXISTS players (
			tag TEXT PRIMARY KEY,
			elo INTEGER)`,
	}

	for _, stmt := range initStmts {
		_, err = ctx.Exec(stmt)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &database{
		ctx: ctx,
	}
}

/*
	Creates and returns a new relationalDBPlayerRepositoru
*/
func createRelationalDBPlayerRepository(path string) RelationalDBPlayerRepository {
	return RelationalDBPlayerRepository{*createDatabase(path)}
}

func (r RelationalDBPlayerRepository) AddPlayer(p Player) error {
	err := ValidateAddPlayer(p)
	if err != nil {
		return err
	}
	_, err = r.db.ctx.Exec(`
	INSERT INTO players(tag,
						elo)
	values(?,?)`, p.Tag, p.Elo)
	return err
}

func (r RelationalDBPlayerRepository) GetPlayer(tag string) (Player, error) {
	var p Player
	err := ValidateGetPlayer(tag)
	if err != nil {
		return p, nil
	}
	stmt, err := r.db.ctx.Prepare("SELECT tag,elo FROM players WHERE tag=?")
	if err != nil {
		return p, err
	}
	defer stmt.Close()
	var t string
	var e int

	err = stmt.QueryRow(tag).Scan(&t, &e)
	if err != nil {
		return p, err
	}
	p = Player{Tag: t, Elo: e}
	return p, err
}

func (r RelationalDBPlayerRepository) UpdatePlayerElo(tag string, elo int) error {
	p, err := r.GetPlayer(tag)
	if err != nil {
		return err
	}

	err = ValidateUpdatePlayerElo(elo)
	if err != nil {
		return err
	}

	stmt, err := r.db.ctx.Prepare("UPDATE players SET elo = ? WHERE tag = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(elo, p.Tag)
	return err
}
