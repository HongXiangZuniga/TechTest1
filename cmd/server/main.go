package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HongXiangZuniga/TestYapo/pkg/http/rest"
	"github.com/HongXiangZuniga/TestYapo/pkg/track"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

func main() {
	log.Println("Init API")
	cache := cache.New(60*time.Minute, 0*time.Minute)
	trackRepo := track.NewRepository(http.Client{}, cache)
	trackService := track.NewService(trackRepo)
	songsHandler := rest.NewUsersHandler(trackService)

	r := rest.NewHandler(songsHandler)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
	}
	if os.Getenv("PORT") == "" {
		log.Println("PORT is necesary for run")
	} else {
		err := r.Run(":" + os.Getenv("PORT"))
		if err != nil {
			log.Println("Api not run: " + err.Error())
		} else {
			log.Println("Api ok")
		}
	}
}
