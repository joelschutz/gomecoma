package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joelschutz/gomecoma/src/ent"
	"github.com/joelschutz/gomecoma/src/ent/artist"

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

	CreateMovie(context.Background(), client)
}

func CreateMovie(ctx context.Context, client *ent.Client) (m *ent.Movie, e error) {
	cast, err := CreateArtists(ctx, client, []string{"Timothée Chalamet", "Rebecca Ferguson", "Zendaya"})
	if err != nil {
		log.Fatal(err)
	}
	directors, err := CreateArtists(ctx, client, []string{"Denis Villeneuve"})
	if err != nil {
		log.Fatal(err)
	}
	writers, err := CreateArtists(ctx, client, []string{"Jon Spaihts", "Denis Villeneuve", "Eric Roth"})
	if err != nil {
		log.Fatal(err)
	}

	m, err = client.Movie.Create().
		SetTitle("Dune").
		SetDuration(155).
		AddCast(cast...).
		AddDirectors(directors...).
		AddWriters(writers...).Save(ctx)
	if err != nil {
		log.Fatalf("failed creating Movie: %v", err)
	}

	return
}

func CreateArtists(ctx context.Context, client *ent.Client, artists []string) (r []*ent.Artist, e error) {
	for _, v := range artists {
		ar, err := client.Artist.Query().Where(artist.Name(v)).First(ctx)
		if err != nil {
			if err == err.(*ent.NotFoundError) {
				ar, err = client.Artist.
					Create().
					SetName(v).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("failed creating Artist: %w", err)
				}
			} else {
				return nil, fmt.Errorf("failed querying user: %w", err)
			}
		}
		r = append(r, ar)
	}

	return
}
