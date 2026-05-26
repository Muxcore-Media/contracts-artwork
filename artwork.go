package artwork

import (
	"context"
	"io"
)

type ArtworkType string

const (
	ArtworkPoster   ArtworkType = "poster"
	ArtworkBackdrop ArtworkType = "backdrop"
	ArtworkLogo     ArtworkType = "logo"
	ArtworkBanner   ArtworkType = "banner"
	ArtworkThumb    ArtworkType = "thumb"
)

type ArtworkCandidate struct {
	ID       string
	Source   string
	ArtType  ArtworkType
	URL      string
	MIMEType string
	Width    int
	Height   int
	Language string
	Score    int
}

type ArtworkProvider interface {
	Search(ctx context.Context, externalID, source string, artType ArtworkType, language string) ([]ArtworkCandidate, error)
	Fetch(ctx context.Context, candidate ArtworkCandidate) (io.ReadCloser, error)
}

const EventArtworkFetched = "artwork.fetched"

type ArtworkFetchedPayload struct {
	MediaID     string `json:"media_id"`
	ArtworkType string `json:"artwork_type"`
	StorageKey  string `json:"storage_key"`
	Source      string `json:"source"`
}
