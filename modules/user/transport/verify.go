package ginuser

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	userbiz "example.com/musicafy_be/modules/user/biz"
	userstorage "example.com/musicafy_be/modules/user/storage"
	"github.com/gin-gonic/gin"
)

type VerifyRequest struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}

func Verify(appContext appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req VerifyRequest
		if err := c.ShouldBind(&req); err != nil {
			panic(err)
		}

		db := appContext.GetMainDBConnection()
		store := userstorage.NewStore(db)
		biz := userbiz.NewVerifyBiz(store)

		if err := biz.Verify(c.Request.Context(), req.Email, req.Code); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
