package songsbiz

import (
	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songmodels "example.com/musicafy_be/modules/songs/models"
)

type ArtistStore interface {
	FindArtist(id string) (*songmodels.Artists, error)
}

type ArtistBiz struct {
	store ArtistStore
}

func NewArtistBiz(store ArtistStore) *ArtistBiz {
	return &ArtistBiz{
		store: store,
	}
}

func (biz *ArtistBiz) GetArtist(appContext appctx.AppContext, maskId string) (*songmodels.Artists, error) {
	zingApi := appContext.GetZingmp3Api()
	artist, err := biz.store.FindArtist(maskId)
	if err != nil {
		return nil, common.ErrDBWithMsg(err, "Không tìm thấy nghệ sĩ")
	}

	artistDetail, err := zingApi.Artist(artist.Alias)
	if err != nil {
		return nil, err
	}

	res := artistDetail.ToModelDb()

	return &res, nil
}
