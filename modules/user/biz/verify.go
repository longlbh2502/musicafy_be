package userbiz

import (
	"context"
	"errors"
	"time"

	"example.com/musicafy_be/common"
	usermodels "example.com/musicafy_be/modules/user/models"
)

type VerifyStore interface {
	FindVerify(email string) (*usermodels.Verify, error)
	UpdateVerify(email string) error
}

type VerifyBiz struct {
	store VerifyStore
}

func NewVerifyBiz(store VerifyStore) *VerifyBiz {
	return &VerifyBiz{store: store}
}

func (biz *VerifyBiz) Verify(ctx context.Context, email string, code string) error {
	verify, err := biz.store.FindVerify(email)
	if err != nil {
		return common.NewCustomError(err, "find verify", "Lỗi tìm kiếm mã xác thực", "VERIFY")
	}

	if verify.ExpiredAt.Before(time.Now()) {
		return common.NewCustomError(errors.New("code expired"), "code expired", "Mã xác thực đã hết hạn", "VERIFY")
	}
	print(verify.SecretCode)
	if verify.SecretCode != code {
		return common.NewCustomError(errors.New("invalid code"), "invalid code", "Mã xác thực không hợp lệ", "VERIFY")
	}

	if err := biz.store.UpdateVerify(email); err != nil {
		return common.NewCustomError(err, "update verify", "Lỗi cập nhật mã xác thực", "VERIFY")
	}

	return nil
}
