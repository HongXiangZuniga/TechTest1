package track

import "testing"

func TestGetTracks(t *testing.T) {
	testService := NewService(&mockRepo{}, []Favorite{})
	type testCase struct {
		band     string
		response Response
	}
	testCases := []testCase{
		{"test",
			Response{
				TotalCollections: 1,
				TotalSongs:       1,
				Collections:      []string{"test"}, Songs: []Song{
					{
						SongId:         123,
						CollectionName: "test",
						SongName:       "test",
						PreviewUrl:     "test",
						ReleaseDate:    "",
						Price: Price{
							Value:    1,
							Currency: "USD",
						},
					},
				},
			},
		}}
	for _, tc := range testCases {
		resp, err := testService.GetTracks(tc.band)
		if err != nil {
			t.Fatalf(err.Error())
		}
		if tc.response.TotalCollections != resp.TotalCollections {
			t.Fatalf("Error in TotalCollections")
		}
		if tc.response.Songs[0].SongId != resp.Songs[0].SongId {
			t.Fatalf("Error in SongId")
		}
		if tc.response.Collections[0] != resp.Collections[0] {
			t.Fatalf("Error in name Collection")
		}
	}
}
