package usermodels

import (
	"errors"
	"time"

	"example.com/musicafy_be/common"
)

type User struct {
	ID                int        `json:"id" gorm:"primaryKey;column:id"`
	Username          string     `json:"username" gorm:"column:username"`
	HashedPassword    string     `json:"-" gorm:"column:hashed_password"`
	FullName          string     `json:"full_name" gorm:"column:full_name"`
	Email             *string    `json:"email,omitempty" gorm:"column:email"`
	IsVerify          bool       `json:"is_verify" gorm:"column:is_verify"`
	IsActive          bool       `json:"is_active" gorm:"column:active"`
	Role              *int       `json:"role,omitempty" gorm:"column:role"`
	Gender            *string    `json:"gender,omitempty"`
	Licence           *string    `json:"licence,omitempty"`
	Dob               *time.Time `json:"dob,omitempty"`
	Avatar            string     `json:"avatar,omitempty" gorm:"column:active"`
	PackageID         *int       `json:"-,omitempty" gorm:"column:package"`
	Package           *Package   `json:"package,omitempty" gorm:"foreignKey:PackageID"`
	PackageExpire     *time.Time `json:"package_expire,omitempty" gorm:"column:package_expire"`
	PasswordChangedAt time.Time  `json:"password_changed_at"`
}

func (User) TableName() string {
	return "users"
}

type AccountQueries struct {
	ID       *int    `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"email hoặc mật khẩu không đúng",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"email đã tồn tại",
		"ErrEmailExisted",
	)
)
