package ginsong

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songsbiz "example.com/musicafy_be/modules/songs/biz"
	songstorage "example.com/musicafy_be/modules/songs/storage"
	"github.com/gin-gonic/gin"
)

func Artist(appctx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appctx.GetMainDBConnection()
		store := songstorage.NewStore(db)

		artistBiz := songsbiz.NewArtistBiz(store)

		artist, err := artistBiz.GetArtist(appctx, ctx.Query("maskId"))
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(artist))
	}
}
