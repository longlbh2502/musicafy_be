package ginuser

import (
	"net/http"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	userbiz "example.com/musicafy_be/modules/user/biz"
	userstorage "example.com/musicafy_be/modules/user/storage"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ResendOtpRequest struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

func ResendOtp(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		store := userstorage.NewStore(db)

		var data ResendOtpRequest
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		resentOtpBiz := userbiz.NewResentOtpBiz(store)
		verify, err := resentOtpBiz.ResentOtp(data.Email)
		if err != nil {
			panic(err)
		}

		// Gửi email với nội dung HTML
		emailSender := appctx.GetMailer()
		err = emailSender.SendOTPMail(data.Email, verify.SecretCode)
		if err != nil {
			log.Error().Err(err).Msgf("Không thể gửi email: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể gửi email"})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
