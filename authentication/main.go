package main

import (
	"flag"
	"fmt"
	"log"
	"microservice/db"

	"github.com/joho/godotenv"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "run authentication service local")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Connection: ", conn.DB().Name)
}
