package songmodels

import "encoding/json"

type Lyric struct {
	Id   int             `json:"-" gorm:"column:id"`
	Song int             `json:"song" gorm:"column:song"`
	File string          `json:"file" gorm:"column:file"`
	Data json.RawMessage `json:"data" gorm:"column:data;type:json"`
}

func (Lyric) TableName() string {
	return "song_lyric"
}
