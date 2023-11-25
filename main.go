package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/amupxm/todo_api/api"
	db "github.com/amupxm/todo_api/db/sqlc"
	"github.com/amupxm/todo_api/util"
	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()
	util.LoadConfig()
	conn, err := sql.Open("postgres", util.GetDatabaseConnectionString())
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(util.Config.ServerAddress)
}
