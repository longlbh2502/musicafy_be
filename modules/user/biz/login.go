package userbiz

import (
	"context"
	"errors"
	"time"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/token"
	usermodels "example.com/musicafy_be/modules/user/models"
	"example.com/musicafy_be/utils"
)

type LoginReq struct {
	Username  *string `json:"username,omitempty"`
	Password  *string `json:"password,omitempty"`
	Email     *string `json:"email,omitempty"`
	ID        *int    `json:"id,omitempty"`
	UserAgent *string `json:"user_agent,omitempty"`
	ClientIp  *string `json:"client_ip,omitempty"`
}

type LoginRes struct {
	Session               *usermodels.Session `json:"session,omitempty"`
	AccessToken           string              `json:"access_token"`
	RefreshToken          string              `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time           `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time           `json:"refresh_token_expires_at"`
	User                  usermodels.User     `json:"data"`
}

type LoginStore interface {
	FindUser(arg usermodels.AccountQueries) (*usermodels.User, error)
	CreateSession(ctx context.Context, arg usermodels.Session) (*usermodels.Session, error)
}

type loginBiz struct {
	store LoginStore
}

func NewLoginBiz(store LoginStore) *loginBiz {
	return &loginBiz{
		store: store,
	}
}

func (biz *loginBiz) LoginBiz(ctx context.Context, arg LoginReq, token token.TokenMaker) (*LoginRes, error) {
	user, err := biz.store.FindUser(usermodels.AccountQueries{
		Username: arg.Username,
		Email:    arg.Email,
		ID:       arg.ID,
	})
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, common.ErrInvalidRequest(errors.New("account inactive"))
	}
	if password := *arg.Password; arg.Password != nil {
		if err = utils.CheckPassword(password, user.HashedPassword); err != nil {
			return nil, usermodels.ErrUsernameOrPasswordInvalid
		}
	}

	accressToken, accressTokenPayload, _ := token.CreateToken(*user)
	refreshToken, refreshTokenPayload, _ := token.CreateToken(*user)

	session, err := biz.store.CreateSession(ctx, usermodels.Session{
		Username:     user.Username,
		RefreshToken: refreshToken,
		// UserAgent:    metadata.UserAgent,
		ClientIp:  *arg.ClientIp,
		IsBlocked: false,
		ExpiresAt: refreshTokenPayload.ExpireAt,
	})
	if err != nil {
		return nil, common.NewCustomError(err, "failed to create session: %e", "Lỗi tạo session đăng nhập", "SESSION_CREATE")
		// return nil, status.Errorf(codes.Internal, "failed to create session: %e", err)
	}

	res := LoginRes{
		User:                  *user,
		Session:               session,
		AccessToken:           accressToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accressTokenPayload.ExpireAt,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpireAt,
	}

	return &res, nil
}
