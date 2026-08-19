package router

import (
	"example.com/musicafy_be/components/appctx"
	"example.com/musicafy_be/middleware"
	ginsong "example.com/musicafy_be/modules/songs/transport"
	ginuser "example.com/musicafy_be/modules/user/transport"
	"github.com/gin-gonic/gin"
)

func SetupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	user := v1.Group("/user")
	{
		user.POST("/register", ginuser.Register(appContext))
		user.POST("/login", ginuser.Login(appContext))
		user.POST("/otp/verify", ginuser.Verify(appContext))
		user.POST("/otp/resend", ginuser.ResendOtp(appContext))
	}

	song := v1.Group("/song")
	{
		song.GET("/search-suggestions", middleware.RequiredAuth(appContext), ginsong.SearchSuggestion(appContext))
		song.GET("/streaming", ginsong.StreamingSong(appContext))
		song.GET("/info-detail", ginsong.DetailSong(appContext))
		song.GET("/lyric", ginsong.Lyric(appContext))
		song.GET("/home", ginsong.HomeApi(appContext))
		song.GET("/hub-home", ginsong.Lyric(appContext))
		song.GET("/hub-detail", ginsong.Lyric(appContext))
		song.GET("/artist", ginsong.Artist(appContext))
	}
}
