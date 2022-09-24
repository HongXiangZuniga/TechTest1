package track

import "time"

type DataApi struct {
	ResultCount int       `json:"resultCount"`
	Results     []Element `json:"results"`
}

type Element struct {
	WrapperType            string    `json:"wrapperType,omitempty"`
	Kind                   string    `json:"kind,omitempty"`
	ArtistID               int       `json:"artistId,omitempty"`
	CollectionID           int       `json:"collectionId,omitempty"`
	TrackID                int       `json:"trackId,omitempty"`
	ArtistName             string    `json:"artistName,omitempty"`
	CollectionName         string    `json:"collectionName,omitempty"`
	TrackName              string    `json:"trackName,omitempty"`
	CollectionCensoredName string    `json:"collectionCensoredName,omitempty"`
	TrackCensoredName      string    `json:"trackCensoredName,omitempty"`
	ArtistViewURL          string    `json:"artistViewUrl,omitempty"`
	CollectionViewURL      string    `json:"collectionViewUrl,omitempty"`
	TrackViewURL           string    `json:"trackViewUrl,omitempty"`
	PreviewURL             string    `json:"previewUrl,omitempty"`
	ArtworkURL30           string    `json:"artworkUrl30,omitempty"`
	ArtworkURL60           string    `json:"artworkUrl60,omitempty"`
	ArtworkURL100          string    `json:"artworkUrl100,omitempty"`
	CollectionPrice        float64   `json:"collectionPrice,omitempty"`
	TrackPrice             float64   `json:"trackPrice,omitempty"`
	ReleaseDate            time.Time `json:"releaseDate,omitempty"`
	CollectionExplicitness string    `json:"collectionExplicitness,omitempty"`
	TrackExplicitness      string    `json:"trackExplicitness,omitempty"`
	DiscCount              int       `json:"discCount,omitempty"`
	DiscNumber             int       `json:"discNumber,omitempty"`
	TrackCount             int       `json:"trackCount,omitempty"`
	TrackNumber            int       `json:"trackNumber,omitempty"`
	TrackTimeMillis        int       `json:"trackTimeMillis,omitempty"`
	Country                string    `json:"country,omitempty"`
	Currency               string    `json:"currency,omitempty"`
	PrimaryGenreName       string    `json:"primaryGenreName,omitempty"`
	IsStreamable           bool      `json:"isStreamable,omitempty"`
}

type Response struct {
	TotalCollections int      `json:"total_albumes"`
	TotalSongs       int      `json:"total_canciones"`
	Collections      []string `json:"albumes"`
	Songs            []Song   `json:"canciones"`
}

type Song struct {
	SongId         int    `json:"cancion_id"`
	CollectionName string `json:"nombre_album"`
	SongName       string `json:"nombre_tema"`
	PreviewUrl     string `json:"preview_url"`
	ReleaseDate    string `json:"fecha_lanzamiento"`
	Price          Price  `json:"precio"`
}

type Price struct {
	Value    int    `json:"valor"`
	Currency string `json:"moneda"`
}
