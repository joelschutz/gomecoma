package services

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joelschutz/gomecoma/src/ent"
)

func ScanFolder(ctx context.Context, client *ent.Client, path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			parts := strings.Split(info.Name(), ".")
			switch parts[len(parts)-1] {
			case
				"mp4",
				"avi",
				"mkv":
				f, err := client.File.Create().
					SetName(info.Name()).
					SetPath(path).
					SetType("video").Save(ctx)
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(f)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
