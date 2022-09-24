package track

import "fmt"

type Service interface {
	GetTracks(band string) (*Response, error)
}

type port struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &port{repository}
}

func (port *port) GetTracks(band string) (*Response, error) {
	response, err := port.repository.GetTracks(band)
	if err != nil {
		return nil, err
	}
	for _, element := range response.Results {
		fmt.Println(element.ArtistName)
	}
	return response, nil

}
