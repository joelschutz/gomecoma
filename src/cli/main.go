package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/joelschutz/gomecoma/src/ent"
	"github.com/manifoldco/promptui"
)

type OptionAction func(ctx context.Context, client *ent.Client) error

type Option struct {
	Name   string
	Action OptionAction
	Active bool
	Icon   string
}

var optionTemplate promptui.SelectTemplates = promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "{{.Icon}} {{ .Name | cyan }} ({{ .Active | red }})",
	Inactive: "  {{ .Name | cyan }} ({{ .Active | red }})",
	Selected: "{{.Icon}} {{ .Name | red | cyan }}",
	Details: `
--------- Option ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Active:" | faint }}	{{ .Active }}`,
}

func LoopCLI(ctx context.Context, client *ent.Client) {
	// for {
	// 	PromptMovieUpdate(ctx, client)
	// }
	options := []Option{
		Option{Name: "ScanFolder", Action: promptScanFolder, Active: true, Icon: "\U0001F5C1"},
		Option{Name: "Files", Action: promptFilesMenu, Active: true, Icon: "\U0001F5C1"},
	}
	MainMenu(ctx, client, options)
}

func MainMenu(ctx context.Context, client *ent.Client, options []Option) {
	options = append(options, Option{
		Name: "Exit",
		Action: func(ctx context.Context, client *ent.Client) error {
			return errors.New("Exit")
		},
		Active: true,
		Icon:   "\U00002716",
	})
	for {

		fmt.Println("GoMeCoMa v 0.1.0")
		prompt := promptui.Select{
			Label:     "Pick and Option:",
			Items:     options,
			Templates: &optionTemplate,
		}

		i, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		err = options[i].Action(ctx, client)
		if err != nil {
			return
		}
	}

}
