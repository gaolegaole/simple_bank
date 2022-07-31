package main

import (
	"database/sql"
	"log"

	"github.com/gaolegaole/simple_bank/api"
	db "github.com/gaolegaole/simple_bank/db/sqlc"
	"github.com/gaolegaole/simple_bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file :", err)
	}
	log.Printf("%#v", config)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("can't start server:", err)
	}
}
