package songsbiz

import (
	"encoding/json"
	"errors"
	"fmt"

	"example.com/musicafy_be/common"
	"example.com/musicafy_be/components/appctx"
	songmodels "example.com/musicafy_be/modules/songs/models"
	"gorm.io/gorm"
)

type LyricStore interface {
	FindSong(id string) (*songmodels.Songs, error)
	FindLyric(song int) (*songmodels.Lyric, error)
	CreateLyric(data songmodels.Lyric) (*songmodels.Lyric, error)
}

type LyricBiz struct {
	store LyricStore
}

func NewLyricBiz(store LyricStore) *LyricBiz {
	return &LyricBiz{
		store: store,
	}
}

func (a *LyricBiz) Lyric(appctx appctx.AppContext, songId string) (*songmodels.Lyric, error) {

	song, err := a.store.FindSong(songId)
	if err != nil {
		return nil, common.ErrDBWithMsg(err, "Không có bài hát")
	}

	lyric, err := a.store.FindLyric(song.ID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		zingmp3 := appctx.GetZingmp3Api()
		lyricZing, err := zingmp3.LyricSong(songId)
		if err != nil {
			return nil, common.ErrInternalWithMsg(err, "Lỗi lấy file Lyric Zingmp3")
		}
		jsonBytes, err := json.Marshal(lyricZing.Sentences)
		if err != nil {
			return nil, fmt.Errorf("lỗi khi chuyển đổi thành JSON: %v", err)
		}
		dataLyric := songmodels.Lyric{
			Song: song.ID,
			File: lyricZing.File,
			Data: jsonBytes,
		}
		lyric, err = a.store.CreateLyric(dataLyric)
		if err != nil {
			return nil, common.ErrDBWithMsg(err, "lỗi tạo lyric từ zingmp3")
		}
	}

	return lyric, nil
}
