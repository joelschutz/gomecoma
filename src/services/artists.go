package services

import (
	"context"
	"fmt"

	"github.com/joelschutz/gomecoma/src/ent"
	"github.com/joelschutz/gomecoma/src/ent/artist"
)

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
