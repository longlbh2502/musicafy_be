package usermodels

import (
	"time"

	"example.com/musicafy_be/common"
)

type Verify struct {
	common.SQLModel `json:",inline" gorm:"-"`
	Username        string    `json:"username" gorm:"column:username;"`
	Email           string    `json:"email" gorm:"column:email;"`
	SecretCode      string    `json:"secret_code" gorm:"column:secret_code;"`
	ExpiredAt       time.Time `json:"expired_at" gorm:"column:expired_at;"`
}

func (Verify) TableName() string {
	return "verifies"
}
