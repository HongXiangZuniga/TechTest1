package track

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type Service interface {
	GetTracks(band string) (*Response, error)
	AddFavorite(Favorite) error
}

type port struct {
	repository Repository
	favorite   []Favorite
}

func NewService(repository Repository, favorite []Favorite) Service {
	return &port{repository, favorite}
}

func (port *port) GetTracks(band string) (*Response, error) {
	var response Response
	trackSelected := make([]Element, 0)
	resultApi, err := port.repository.GetTracks(band)
	if err != nil {
		return nil, err
	}
	for _, element := range resultApi.Results {
		if element.ArtistName == band && element.WrapperType == "track" && element.Kind == "song" {
			trackSelected = append(trackSelected, element)
		}
	}
	for i, track := range trackSelected {
		if i < 25 {
			response.Songs = append(response.Songs, Song{
				SongId:         track.TrackID,
				CollectionName: track.CollectionName,
				SongName:       track.TrackName,
				PreviewUrl:     track.PreviewURL,
				ReleaseDate:    formatDate(track.ReleaseDate),
				Price: Price{
					Value:    int(math.Round(track.TrackPrice)),
					Currency: track.Currency,
				},
			})
			if !stringInArray(track.CollectionName, response.Collections) {
				response.Collections = append(response.Collections, track.CollectionName)
			}
		}
	}
	response.TotalSongs = len(response.Songs)
	response.TotalCollections = len(response.Collections)
	return &response, nil

}

func (port *port) AddFavorite(favorite Favorite) error {
	resultApi, err := port.repository.GetTracks(favorite.BandName)
	if err != nil {
		return err
	}
	for _, element := range resultApi.Results {
		if element.ArtistName == favorite.BandName && element.WrapperType == "track" && element.Kind == "song" {
			if element.TrackID == favorite.SongId {
				port.favorite = append(port.favorite, favorite)
				return nil
			}
		}
	}
	return errors.New("Favorite is not valid")
}

func stringInArray(a string, array []string) bool {
	for _, b := range array {
		if b == a {
			return true
		}
	}
	return false
}

func formatDate(t time.Time) string {
	return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
}
