package main

import (
	"context"
	"log"

	"github.com/joelschutz/gomecoma/src/cli"
	"github.com/joelschutz/gomecoma/src/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// es, _ := client.Movie.Query().All(context.Background())
	// if len(es) < 1 {
	// 	services.CreateMovie(context.Background(), client)
	// }

	cli.LoopCLI(context.Background(), client)
	// fs, _ := client.File.Query().All(context.Background())
	// if len(fs) < 1 {
	// 	services.ScanFolder(context.Background(), client, "/media/joelschutz/HDD_EXT/Filmes")
	// 	for _, f := range fs {
	// 		services.SyncMovies(context.Background(), client, f)
	// 	}
	// } else {
	// 	go cli.LoopCLI(context.Background(), client)
	// }

	// // Start listening.
	// srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Starting Server")
	// if err := http.ListenAndServe(":8080", srv); err != nil {
	// 	log.Fatal(err)
	// }
}
