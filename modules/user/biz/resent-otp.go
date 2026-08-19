package userbiz

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	usermodels "example.com/musicafy_be/modules/user/models"
)

type ResentOtpStore interface {
	CreateVerify(data usermodels.Verify) (*usermodels.Verify, error)
}

type ResentOtpBiz struct {
	store ResentOtpStore
}

func NewResentOtpBiz(store ResentOtpStore) *ResentOtpBiz {
	return &ResentOtpBiz{store: store}
}

func (biz *ResentOtpBiz) ResentOtp(email string) (*usermodels.Verify, error) {

	code, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return nil, err
	}

	verify := usermodels.Verify{
		Username:   email,
		Email:      email,
		SecretCode: strconv.Itoa(int(code.Int64() + 100000)),
		ExpiredAt:  time.Now().Add(time.Minute * 5),
	}
	_, err = biz.store.CreateVerify(verify)
	if err != nil {
		return nil, err
	}

	return &verify, nil
}
