package ginuser

import (
	"errors"
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	userbiz "example.com/musicafy_be/modules/user/biz"
	userstorage "example.com/musicafy_be/modules/user/storage"
	"github.com/gin-gonic/gin"
)

func Register(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		store := userstorage.NewStore(db)
		biz, err := userbiz.NewRegisterBiz(store)
		if err != nil {
			panic(err)
		}

		var arg userbiz.RegisterReq
		if err := c.ShouldBind(&arg); err != nil {
			panic(err)
		}

		if arg.Email == nil || arg.FullName == nil || arg.Password == nil {
			panic(common.ErrInvalidRequest(errors.New("Invalid body")))
		}

		code, err := biz.Register(arg)
		if err != nil {
			panic(err)
		}

		mailer := appctx.GetMailer()
		err = mailer.SendOTPMail(*arg.Email, *code)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
