package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/joelschutz/gomecoma/src/ent"
	"github.com/joelschutz/gomecoma/src/services"
	"github.com/manifoldco/promptui"
)

type FileAction func(ctx context.Context, client *ent.Client, f *ent.File) error
type FileOptionAction func(ctx context.Context, client *ent.Client, action FileAction) error

type FileOption struct {
	Name   string
	Active bool
	Icon   string
	Action FileOptionAction
}

var fileTemplate promptui.SelectTemplates = promptui.SelectTemplates{
	Label:    "{{ . }}?",
	Active:   "\U00002192 {{ .Name | cyan }} ({{ .Synced | red }})",
	Inactive: "  {{ .Name | cyan }} ({{ .Synced | red }})",
	Selected: "\U00002192 {{ .Name | red | cyan }}",
	Details: `
--------- Option ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Synced:" | faint }}	{{ .Synced }}
{{ "Type:" | faint }}	{{ .Type }}
{{ "Path:" | faint }}	{{ .Path }}`,
}

func promptFilesMenu(ctx context.Context, client *ent.Client) error {
	fmt.Println("Select Action")

	options := []FileOption{
		FileOption{
			Name: "Update",
			// Action: PromptMovieUpdate,
			Active: true,
			Icon:   "\U00002192",
		},
	}
	prompt := promptui.Select{
		Label:     "Pick and Option:",
		Items:     options,
		Templates: &optionTemplate,
	}

	_, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	err = promptSelectFileMethod(ctx, client, PromptMovieUpdate)
	if err != nil {
		return err
	}
	return nil
}

func promptSelectFileMethod(ctx context.Context, client *ent.Client, action FileAction) error {
	fmt.Println("Navegate or enter Id?")

	options := []FileOption{
		FileOption{
			Name:   "Navegate files",
			Action: promptNavegateFiles,
			Active: true,
			Icon:   "\U00002192",
		},
		FileOption{
			Name:   "Enter file Id",
			Action: promptEnterFileId,
			Active: true,
			Icon:   "\U00002192",
		},
	}
	prompt := promptui.Select{
		Label:     "Pick and Option:",
		Items:     options,
		Templates: &optionTemplate,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	err = options[i].Action(ctx, client, action)
	if err != nil {
		return err
	}
	return nil
}

func promptScanFolder(ctx context.Context, client *ent.Client) error {
	validate := func(input string) error {
		if strings.Contains(input, "/") {
			return nil
		}
		return errors.New("Invalid path")
	}

	fileSelect := promptui.Prompt{
		Label:    "What is the path to the folder?",
		Validate: validate,
	}

	result, err := fileSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	services.ScanFolder(ctx, client, result)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}
	return nil
}

func promptNavegateFiles(ctx context.Context, client *ent.Client, action FileAction) error {

	fmt.Println("Navegate trought scanned files")

	options, err := client.File.Query().All(ctx)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}
	prompt := promptui.Select{
		Label:     "Pick and Option:",
		Items:     options,
		Templates: &fileTemplate,
		Size:      20,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	err = action(ctx, client, options[i])
	if err != nil {
		return err
	}
	return nil
}

func promptEnterFileId(ctx context.Context, client *ent.Client, action FileAction) error {
	validate := func(input string) error {
		_, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	fileSelect := promptui.Prompt{
		Label:    "What is the file ID?",
		Validate: validate,
	}

	result, err := fileSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	cr, err := strconv.Atoi(result)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	f, err := client.File.Get(ctx, cr)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	err = action(ctx, client, f)
	if err != nil {
		return err
	}
	return nil
}
