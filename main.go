package main

import (
	"database/sql"
	"github.com/gaolegaole/simple_bank/api"
	db "github.com/gaolegaole/simple_bank/db/sqlc"
	"github.com/gaolegaole/simple_bank/util"
	"log"
	"runtime"

	_ "github.com/lib/pq"
)

func main() {
	print(">>>>>>>>>>>>>>GOOS:", runtime.GOOS, ",GOARCH:", runtime.GOARCH, "<<<<<<<<<<<<<\n")
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
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatalf("cannot create server %v", err)
	}

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("can't start server:", err)
	}
}
