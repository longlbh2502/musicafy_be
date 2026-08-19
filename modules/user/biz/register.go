package userbiz

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"example.com/musicafy_be/common"
	usermodels "example.com/musicafy_be/modules/user/models"
	"example.com/musicafy_be/utils"
)

type RegisterStore interface {
	FindUser(arg usermodels.AccountQueries) (*usermodels.User, error)
	CreateAccount(data usermodels.User) (*int, error)
	CreateVerify(data usermodels.Verify) (*usermodels.Verify, error)
}

type registerBiz struct {
	store RegisterStore
}

type RegisterReq struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	FullName *string `json:"full_name"`
}

func NewRegisterBiz(store RegisterStore) (*registerBiz, error) {
	return &registerBiz{
		store: store,
	}, nil
}

func (biz *registerBiz) Register(req RegisterReq) (*string, error) {
	user, err := biz.store.FindUser(usermodels.AccountQueries{
		Username: req.Email,
		Email:    req.Email,
	})
	if err != nil {
		if err.Error() != "record not found" {
			return nil, common.NewCustomError(err, "find user with info", "Thông tin đăng ký đã tồn tại", "REGISTER")
		}
	}

	if user != nil && user.IsVerify {
		return nil, common.NewCustomError(nil, "user already verified", "Tài khoản đã được xác thực", "REGISTER")
	}

	if user != nil && !user.IsVerify {

	} else if user != nil && user.IsVerify {
		return nil, common.NewCustomError(nil, "user already verified", "Tài khoản đã được xác thực", "REGISTER")
	} else {
		hassPassword, err := utils.HashedPassword(*req.Password)
		if err != nil {
			return nil, common.NewCustomError(err, "error hash password", "Lỗi hashpassword", "REGISTER")
		}

		_, err = biz.store.CreateAccount(usermodels.User{
			Username:       *req.Email,
			HashedPassword: hassPassword,
			FullName:       *req.FullName,
			Email:          req.Email,
		})
		if err != nil {
			return nil, common.NewCustomError(err, "error create user", "Lỗi tạo bản ghi User", "REGISTER")
		}

	}

	code, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return nil, common.NewCustomError(err, "error generate code", "Lỗi tạo mã xác thực", "REGISTER")
	}

	codeStr := strconv.Itoa(int(code.Int64() + 100000))

	verify := usermodels.Verify{
		Username:   *req.Email,
		Email:      *req.Email,
		SecretCode: codeStr,
		ExpiredAt:  time.Now().Add(time.Minute * 5),
	}

	_, err = biz.store.CreateVerify(verify)
	if err != nil {
		return nil, common.NewCustomError(err, "error create verify", "Lỗi tạo bản ghi Verify", "REGISTER")
	}
	return &codeStr, nil

}
