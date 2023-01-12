package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joelschutz/gomecoma/src/ent"
	"github.com/joelschutz/gomecoma/src/tmdb"
)

func CreateMovieFromFile(ctx context.Context, client *ent.Client, f *ent.File, r tmdb.ResultFull) (m *ent.Movie, e error) {
	releaseDate, err := time.Parse("2006-01-02", r.ReleaseDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	m, _ = f.QueryMovie().First(ctx)
	if m == nil {
		m, err = client.Movie.Create().
			SetTitle(r.Title).
			SetOriginalTitle(r.OriginalTitle).
			SetPlot(r.Overview).
			SetDuration(int(r.Runtime)).
			SetReleaseDate(releaseDate).
			SetFile(f).Save(ctx)
		if err != nil {
			log.Fatalf("failed creating Movie: %v", err)
		}
	} else {
		m.Update().
			SetTitle(r.Title).
			SetOriginalTitle(r.OriginalTitle).
			SetPlot(r.Overview).
			SetDuration(int(r.Runtime)).
			SetReleaseDate(releaseDate).
			SetFile(f).Save(ctx)
		if err != nil {
			log.Fatalf("failed creating Movie: %v", err)
		}
	}

	fmt.Println(m)

	fu, err := client.File.UpdateOne(f).
		SetExternalID(strconv.Itoa(int(r.Id))).
		SetExternalInfoProvider("TMDB").
		SetMovie(m).
		SetSynced(true).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(fu)
	return
}

func SyncMovies(ctx context.Context, client *ent.Client, f *ent.File) (*ent.File, error) {
	a := regexp.MustCompile(`\.\d{4}\.`)
	ms := a.Split(f.Name, -1)
	r := strings.Split(ms[0], ".")
	q := strings.Join(r, "%20")
	fmt.Println(strings.ReplaceAll(q, "%20", " "))

	responseObject, err := tmdb.SearchMovie(&http.Client{}, q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	results := ""
	for _, result := range responseObject.Results {
		results += strconv.Itoa(int(result.Id)) + ","
	}
	results = strings.TrimSuffix(results, ",")

	f, err = client.File.UpdateOne(f).SetResults(results).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(f)
	return f, nil
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
