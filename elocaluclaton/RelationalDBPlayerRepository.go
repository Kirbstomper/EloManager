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

func createDatabase() *database {
	dbDir := "data"
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
func createRelationalDBPlayerRepository() RelationalDBPlayerRepository {
	return RelationalDBPlayerRepository{*createDatabase()}
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

//func (RelationalDBPlayerRepository) GetPlayer(string) (Player, error)
//func (RelationalDBPlayerRepository) UpdatePlayerElo(string, int) error
