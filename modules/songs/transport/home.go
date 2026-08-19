package ginsong

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songsbiz "example.com/musicafy_be/modules/songs/biz"
	"github.com/gin-gonic/gin"
)

func HomeApi(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		// db := appctx.GetMainDBConnection()
		// store := songstorage.NewStore(db)
		biz := songsbiz.NewHomeBiz(appctx)
		homeSections, err := biz.GetHomeSession()
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(homeSections))
	}
}
