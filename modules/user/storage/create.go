package userstorage

import (
	"context"

	"example.com/musicafy_be/common"
	usermodels "example.com/musicafy_be/modules/user/models"
)

func (s *store) CreateAccount(data usermodels.User) (*int, error) {
	var id int
	results := s.db.Table(usermodels.User{}.TableName()).Select("Username", "HashedPassword", "FullName", "Email").Create(&data)
	if results.Error != nil {
		return nil, common.ErrDB(results.Error)
	}
	id = data.ID

	return &id, nil
}

func (s *store) CreateVerify(data usermodels.Verify) (*usermodels.Verify, error) {
	results := s.db.Table(usermodels.Verify{}.TableName()).Create(&data)
	if results.Error != nil {
		return nil, common.ErrDB(results.Error)
	}
	return &data, nil
}

func (s *store) CreateSession(ctx context.Context, data usermodels.Session) (*usermodels.Session, error) {
	results := s.db.Table(usermodels.Session{}.TableName()).Create(&data)
	if results.Error != nil {
		return nil, common.ErrDB(results.Error)
	}
	return &data, nil
}
