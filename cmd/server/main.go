package main

import (
	"log"
	"net/http"
	"os"

	"github.com/HongXiangZuniga/TestYapo/pkg/http/rest"
	"github.com/HongXiangZuniga/TestYapo/pkg/track"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Init API")
	trackRepo := track.NewRepository(http.Client{})
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
