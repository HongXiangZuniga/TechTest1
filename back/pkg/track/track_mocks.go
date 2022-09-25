package track

import (
	"time"
)

type mockRepo struct {
}

func NewMockRepo() Repository {
	return &mockRepo{}
}

func (mockRepo *mockRepo) GetTracks(term string) (*DataApi, error) {
	var response DataApi
	if term == "test" {
		response.ResultCount = 1
		response.Results = append(response.Results, Element{
			WrapperType:            "track",
			Kind:                   "song",
			ArtistID:               321243033,
			CollectionID:           1010499557,
			TrackID:                123,
			ArtistName:             "test",
			CollectionName:         "test",
			TrackName:              "test",
			CollectionCensoredName: "Night Sessions",
			TrackCensoredName:      "Radiohead",
			ArtistViewURL:          "https://music.apple.com/us/artist/nicholas-cole/321243033?uo=4",
			CollectionViewURL:      "https://music.apple.com/us/album/radiohead/1010499557?i=1010499559&uo=4",
			TrackViewURL:           "https://music.apple.com/us/album/radiohead/1010499557?i=1010499559&uo=4",
			PreviewURL:             "test",
			ArtworkURL30:           "https://is4-ssl.mzstatic.com/image/thumb/Music5/v4/a9/06/4f/a9064fbf-01c2-1e02-c377-ed47ef914eb3/886445346596.jpg/30x30bb.jpg",
			ArtworkURL60:           "https://is4-ssl.mzstatic.com/image/thumb/Music5/v4/a9/06/4f/a9064fbf-01c2-1e02-c377-ed47ef914eb3/886445346596.jpg/60x60bb.jpg",
			ArtworkURL100:          "https://is4-ssl.mzstatic.com/image/thumb/Music5/v4/a9/06/4f/a9064fbf-01c2-1e02-c377-ed47ef914eb3/886445346596.jpg/100x100bb.jpg",
			CollectionPrice:        5.99,
			TrackPrice:             1.29,
			ReleaseDate:            time.Now(),
			CollectionExplicitness: "notExplicit",
			TrackExplicitness:      "notExplicit",
			DiscCount:              1,
			DiscNumber:             1,
			TrackCount:             10,
			TrackNumber:            2,
			TrackTimeMillis:        220817,
			Country:                "USA",
			Currency:               "USD",
			PrimaryGenreName:       "Jazz",
			IsStreamable:           true,
		})
	}
	return &response, nil
}
