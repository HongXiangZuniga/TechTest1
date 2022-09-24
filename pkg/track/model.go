package track

import "time"

type Response struct {
	ResultCount int      `json:"resultCount"`
	Results     []Result `json:"results"`
}

type Result struct {
	WrapperType            string    `json:"wrapperType"`
	ArtistID               int       `json:"artistId"`
	CollectionID           int       `json:"collectionId"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	ArtistViewURL          string    `json:"artistViewUrl"`
	CollectionViewURL      string    `json:"collectionViewUrl"`
	ArtworkURL60           string    `json:"artworkUrl60"`
	ArtworkURL100          string    `json:"artworkUrl100"`
	CollectionPrice        float64   `json:"collectionPrice"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackCount             int       `json:"trackCount"`
	Copyright              string    `json:"copyright"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	ReleaseDate            time.Time `json:"releaseDate"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	PreviewURL             string    `json:"previewUrl"`
	Description            string    `json:"description"`
}
