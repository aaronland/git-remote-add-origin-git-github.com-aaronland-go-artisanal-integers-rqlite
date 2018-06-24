package main

import (
	"github.com/aaronland/go-artisanal-integers-rqlite/engine"
	"github.com/aaronland/go-artisanal-integers/application"
	"log"
	"os"
)

func main() {

	flags := application.NewServerApplicationFlags()

	var dsn string
	flags.StringVar(&dsn, "dsn", "http://localhost:4001", "The data source name (dsn) for connecting to Rqlite.")

	application.ParseFlags(flags)

	eng, err := engine.NewRqliteEngine(dsn)

	if err != nil {
		log.Fatal(err)
	}

	app, err := application.NewServerApplication(eng)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(flags)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
