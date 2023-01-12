package cli

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/joelschutz/gomecoma/src/ent"
	"github.com/joelschutz/gomecoma/src/services"
	"github.com/joelschutz/gomecoma/src/tmdb"
	"github.com/manifoldco/promptui"
)

func PromptMovieUpdate(ctx context.Context, client *ent.Client, f *ent.File) error {
	f, err := services.SyncMovies(ctx, client, f)
	if err != nil {
		fmt.Println(err)
		return err
	}

	Options := []tmdb.ResultFull{}
	for _, r := range strings.Split(f.Results, ",") {
		r, err := tmdb.RetrieveMovieInfo(&http.Client{}, r)
		if err != nil {
			fmt.Println(err)
			return err
		}
		Options = append(Options, r)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U00002192 {{ .Title | cyan }} ({{ .ReleaseDate | red }})",
		Inactive: "  {{ .Title | cyan }} ({{ .ReleaseDate | red }})",
		Selected: "\U00002192 {{ .Title | red | cyan }}",
		Details: `
--------- Option ----------
{{ "Title:" | faint }}	{{ .Title }}
{{ "ReleaseDate:" | faint }}	{{ .ReleaseDate }}
{{ "OriginalTitle:" | faint }}	{{ .OriginalTitle }}
{{ "Overview:" | faint }}	{{ .Overview }}`,
	}

	prompt := promptui.Select{
		Label:     "Select Result",
		Items:     Options,
		Templates: templates,
	}

	SelectResult, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	fmt.Printf("You choose %d - %q\n", SelectResult, Options[SelectResult].Title)

	services.CreateMovieFromFile(ctx, client, f, Options[SelectResult])
	fmt.Println("Movie synced")
	return nil
}
