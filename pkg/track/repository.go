package track

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/patrickmn/go-cache"
)

type Repository interface {
	GetTracks(term string) (*DataApi, error)
}

type repo struct {
	httpClient http.Client
	cache      *cache.Cache
}

func NewRepository(httpClient http.Client, cache *cache.Cache) Repository {
	return &repo{httpClient, cache}
}

func (repo *repo) GetTracks(term string) (*DataApi, error) {
	var response DataApi
	resultCached, found := repo.cache.Get(term)
	if found {
		result, _ := resultCached.(DataApi)
		log.Println(fmt.Sprintf("Result for term [%s] in memory cache!", term))
		return &result, nil
	}
	url := "https://itunes.apple.com/search?term=" + url.QueryEscape(term) + "&limit=200"
	resp, err := repo.httpClient.Get(url)
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
	repo.cache.Set(term, response, cache.DefaultExpiration)
	return &response, nil
}
