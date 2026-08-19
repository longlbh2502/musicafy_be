package songmodels

type Artists struct {
	ID     int    `json:"-" gorm:"primaryKey;column:id"`
	MaskId string `json:"mask_id" gorm:"column:mask_id"`
	Name   string `json:"name" gorm:"column:name"`
	// Sportlight  bool    `json:"sportlight" gorm:"column:sportlight"`
	Alias       string     `json:"alias" gorm:"column:alias"`
	Thumbnail   *string    `json:"thumbnail" gorm:"column:thumbnail"`
	ThumbnailM  *string    `json:"thumbnailM" gorm:"column:thumbnailM"`
	PlaylistId  *string    `json:"playlist_id" gorm:"column:playlistId"`
	TotalFollow int        `json:"total_follow" gorm:"column:totalFollow"`
	Sections    *[]Section `json:"sections" gorm:"-"`
}

func (Artists) TableName() string {
	return "artists"
}

type Section struct {
	Items []Songs `json:"items"`
	Type  string  `json:"sectionType"`
	Title string  `json:"title"`
}

type Composers struct {
	ID        int     `json:"-" gorm:"primaryKey;column:id"`
	MaskId    string  `json:"mask_id" gorm:"column:mask_id"`
	Name      string  `json:"name" gorm:"column:name"`
	Spotlight bool    `json:"spotlight" gorm:"column:spotlight"`
	Alias     string  `json:"alias" gorm:"column:alias"`
	Thumbnail *string `json:"thumbnail" gorm:"column:thumbnail"`
	Cover     *string `json:"cover" gorm:"column:cover"`
}

func (Composers) TableName() string {
	return "composers"
}
