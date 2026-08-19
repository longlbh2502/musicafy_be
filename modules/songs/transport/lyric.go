package ginsong

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songsbiz "example.com/musicafy_be/modules/songs/biz"
	songstorage "example.com/musicafy_be/modules/songs/storage"
	"github.com/gin-gonic/gin"
)

func Lyric(appContext appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appContext.GetMainDBConnection()
		store := songstorage.NewStore(db)
		biz := songsbiz.NewLyricBiz(store)

		songId := c.Query("song_id")
		lyric, err := biz.Lyric(appContext, songId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(lyric))
	}
}
