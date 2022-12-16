package db

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"
)

type BleMeshConfigurationHTTP struct {
	ID     int             `json:"id"`
	Config json.RawMessage `json:"config"`
}
type BleMeshConfigurationDB struct {
	ID     int            `db:"id" json:"id"`
	Config types.JSONText `db:"config" json:"config"`
}

var db *sqlx.DB

func Connect() {
	var err error
	db, err = sqlx.Connect("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateConfiguration(config *BleMeshConfigurationDB) {
	res, err := db.NamedExec(`UPDATE mesh_configuration SET config=:config WHERE id=:id`,
		map[string]interface{}{
			"config": config.Config,
			"id":     config.ID,
		})

	if err != nil {
		log.Fatal("configuration update error: ", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal("get rows affected error: ", err)
	}

	if n == 0 {
		log.Fatal("configuration update error: ", err)
	}
}

func GetConfiguration(id int, resultConfig *BleMeshConfigurationDB) {
	err := db.Get(resultConfig, "SELECT * FROM mesh_configuration WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
	}
}
