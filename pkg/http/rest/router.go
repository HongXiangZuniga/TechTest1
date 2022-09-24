package rest

import "github.com/gin-gonic/gin"

func NewHandler(songsHandler SongsHandler) *gin.Engine {
	r := gin.Default()
	_ = r.Group("/")
	{
		r.GET("/", songsHandler.GetSongsByBand)
	}
	return r
}
