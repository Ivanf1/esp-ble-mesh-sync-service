package db

import (
	"log"
	"os"

	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"
)

type BleMeshConfigurationHTTP struct {
	ID     string
	Config json.RawMessage
}
type BleMeshConfigurationDB struct {
	ID     int            `db:"id"`
	Config types.JSONText `db:"config"`
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
	_, err := db.NamedExec(`UPDATE mesh_configuration SET config=:config WHERE id=:id`,
		map[string]interface{}{
			"config": config.Config,
			"id":     config.ID,
		})

	if err != nil {
		log.Fatal(err)
	}
}

func GetConfiguration(id int, resultConfig *BleMeshConfigurationDB) {
	err := db.Get(&resultConfig.Config, "SELECT config FROM mesh_configuration WHERE id=$1", id)

	if err != nil {
		log.Fatal(err)
	}
}
