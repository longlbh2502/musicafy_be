package songstorage

import songmodels "example.com/musicafy_be/modules/songs/models"

func (s *store) FindSong(id string) (*songmodels.Songs, error) {
	var songs songmodels.Songs

	results := s.db.Table(songmodels.Songs{}.TableName()).Where("mask_id = ?", id).First(&songs)
	if results.Error != nil {
		return nil, results.Error
	}

	return &songs, nil
}

func (s *store) FindAlbum(id string) (*songmodels.Albums, error) {
	var albums songmodels.Albums
	results := s.db.Table(songmodels.Albums{}.TableName()).Where("mask_id = ?", id).First(&albums)
	return &albums, results.Error
}

func (s *store) FindArtist(id string) (*songmodels.Artists, error) {
	var artist songmodels.Artists
	results := s.db.Table(songmodels.Artists{}.TableName()).Where("mask_id = ?", id).First(&artist)
	return &artist, results.Error
}

func (s *store) FindLyric(song int) (*songmodels.Lyric, error) {
	var lyric songmodels.Lyric
	results := s.db.Table(songmodels.Lyric{}.TableName()).Where("song = ?", song).First(&lyric)
	return &lyric, results.Error
}
