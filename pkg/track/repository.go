package track

import (
	"encoding/json"
	"net/http"
)

type Repository interface {
	GetTracks(term string) (*Response, error)
}

type api struct {
	httpClient http.Client
}

func NewRepository(httpClient http.Client) Repository {
	return &api{httpClient}
}

func (api *api) GetTracks(term string) (*Response, error) {
	var response Response
	url := "https://itunes.apple.com/search?term=" + term
	resp, err := api.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
