package songstorage

import (
	"example.com/musicafy_be/common"
	songmodels "example.com/musicafy_be/modules/songs/models"
)

func (s *store) CreateSong(data songmodels.Songs) (*songmodels.Songs, error) {
	results := s.db.Table(songmodels.Songs{}.TableName()).FirstOrCreate(
		&data,
		songmodels.Songs{MaskId: data.MaskId},
	)
	if results.Error != nil {
		return nil, common.ErrDB(results.Error)
	}
	return &data, nil
}

func (s *store) CreateAlbum(data songmodels.Albums) (*songmodels.Albums, error) {
	results := s.db.Table(songmodels.Albums{}.TableName()).FirstOrCreate(
		&data,
		songmodels.Albums{MaskId: data.MaskId},
	)
	if results.Error != nil {
		return nil, common.ErrDBWithMsg(results.Error, "Lỗi tạo album")
	}
	return &data, nil
}

func (s *store) CreateArtist(data songmodels.Artists) (*songmodels.Artists, error) {
	results := s.db.Table(songmodels.Artists{}.TableName()).FirstOrCreate(
		&data,
		songmodels.Artists{MaskId: data.MaskId},
	)
	if results.Error != nil {
		return nil, common.ErrDBWithMsg(results.Error, "Lỗi tạo nghệ sĩ")
	}
	return &data, nil
}

func (s *store) CreateComposer(data songmodels.Composers) (*songmodels.Composers, error) {
	results := s.db.Table(songmodels.Composers{}.TableName()).FirstOrCreate(
		&data,
		songmodels.Composers{MaskId: data.MaskId},
	)
	if results.Error != nil {
		return nil, common.ErrDBWithMsg(results.Error, "Lỗi tạo tác giả")
	}
	return &data, nil
}

func (s *store) CreateGenre(data songmodels.Genres) (*songmodels.Genres, error) {
	results := s.db.Table(songmodels.Genres{}.TableName()).FirstOrCreate(
		&data,
		songmodels.Genres{MaskId: data.MaskId},
	)
	if results.Error != nil {
		return nil, common.ErrDBWithMsg(results.Error, "Lỗi tạo thể loại")
	}
	return &data, nil
}

func (s *store) CreateLyric(data songmodels.Lyric) (*songmodels.Lyric, error) {
	results := s.db.Table(songmodels.Lyric{}.TableName()).Where("song = ?", data.Song).FirstOrCreate(
		&data,
	)
	if results.Error != nil {
		return nil, common.ErrDBWithMsg(results.Error, "Lỗi tạo lời bài hát")
	}
	return &data, nil
}
