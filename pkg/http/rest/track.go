package rest

import (
	"net/http"

	"github.com/HongXiangZuniga/TestYapo/pkg/track"
	"github.com/gin-gonic/gin"
)

type SongsHandler interface {
	GetSongsByBand(*gin.Context)
	AddFavorite(ctx *gin.Context)
}

type port struct {
	Track track.Service
}

func NewUsersHandler(Itunes track.Service) SongsHandler {
	return &port{Itunes}
}

func (port *port) GetSongsByBand(ctx *gin.Context) {
	band := ctx.Query("name")
	result, err := port.Track.GetTracks(band)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": result})
	return
}

func (port *port) AddFavorite(ctx *gin.Context) {
	var favorite track.Favorite
	err := ctx.Bind(&favorite)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = port.Track.AddFavorite(favorite)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
