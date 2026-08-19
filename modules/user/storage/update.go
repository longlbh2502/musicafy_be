package userstorage

import usermodels "example.com/musicafy_be/modules/user/models"

func (s *store) UpdateVerify(email string) error {
	results := s.db.Table(usermodels.Verify{}.TableName()).Where("username = ?", email).Update("is_used", true)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
