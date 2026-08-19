package songsbiz

import (
	"example.com/musicafy_be/components/appctx"
	songmodels "example.com/musicafy_be/modules/songs/models"
)

type DetailSongReq struct {
	ID string `json:"id"`
}

type detailSongStore interface {
	FindSong(id string) (*songmodels.Songs, error)
	CreateSong(data songmodels.Songs) (*songmodels.Songs, error)
	FindAlbum(id string) (*songmodels.Albums, error)
	CreateAlbum(data songmodels.Albums) (*songmodels.Albums, error)
	CreateArtist(data songmodels.Artists) (*songmodels.Artists, error)
	CreateComposer(data songmodels.Composers) (*songmodels.Composers, error)
	CreateGenre(data songmodels.Genres) (*songmodels.Genres, error)
}

type detailSongBiz struct {
	store detailSongStore
}

func NewDetailSongBiz(store detailSongStore) *detailSongBiz {
	return &detailSongBiz{
		store: store,
	}
}

func (biz *detailSongBiz) DetailSong(appContext appctx.AppContext, params DetailSongReq) (*songmodels.Songs, error) {
	zingApi := appContext.GetZingmp3Api()
	// call api zing mp3 get song detail
	song, err := zingApi.DetailSong(params.ID)
	if err != nil {
		return nil, err
	}

	// Bắt đầu transaction
	tx := appContext.GetMainDBConnection().Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// find album
	albumDb, err := biz.store.CreateAlbum(songmodels.Albums{
		MaskId:      song.Album.EncodeId,
		Title:       song.Album.Title,
		Thumbnail:   &song.Album.Thumbnail,
		Description: song.Album.SortDescription,
		ReleaseAt:   int(song.Album.ReleasedAt) / 1000,
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var artists []songmodels.Artists
	for _, artist := range song.Artists {
		artistDb, err := biz.store.CreateArtist(songmodels.Artists{
			MaskId:      artist.Id,
			Name:        artist.Name,
			Thumbnail:   artist.Thumbnail,
			ThumbnailM:  artist.ThumbnailM,
			PlaylistId:  &artist.PlaylistId,
			TotalFollow: artist.Followers,
			Alias:       artist.AliasName,
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		artists = append(artists, *artistDb)
	}

	var composers []songmodels.Composers
	for _, composer := range song.Composers {
		composerDb, err := biz.store.CreateComposer(songmodels.Composers{
			MaskId:    composer.ID,
			Name:      composer.Name,
			Spotlight: composer.Spotlight,
			Alias:     composer.Alias,
			Cover:     &composer.Cover,
			Thumbnail: &composer.Thumbnail,
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		composers = append(composers, *composerDb)
	}

	var genres []songmodels.Genres
	for _, genre := range song.Genres {
		genreDb, err := biz.store.CreateGenre(songmodels.Genres{
			MaskId: genre.Id,
			Title:  genre.Name,
			Name:   &genre.Name,
			Alias:  &genre.Alias,
		})
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		genres = append(genres, *genreDb)
	}

	// write song to db
	songDb, err := biz.store.CreateSong(songmodels.Songs{
		MaskId:      params.ID,
		Title:       song.Title,
		Duration:    song.Duration,
		AlbumID:     &albumDb.ID,
		Artists:     artists,
		Composers:   composers,
		Genres:      genres,
		ReleaseDate: int(song.ReleaseDate),
		Distributor: &song.Distributor,
		HasLyric:    song.HasLyric,
		Like:        song.Like,
		Listen:      song.Listen,
		Comment:     song.Comment,
		AudioFile:   song.Link,
		IsOffical:   song.IsOffical,
		Alias:       &song.Alias,
		Thumbnail:   &song.Thumbnail,
		ThumbnailM:  &song.ThumbnailM,
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return songDb, nil
}
