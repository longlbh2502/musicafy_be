package songmodels

type Songs struct {
	ID          int         `json:"-" gorm:"primaryKey;column:id"`
	MaskId      string      `json:"mask_id" gorm:"column:mask_id"`
	Title       string      `json:"title" gorm:"column:title"`
	Alias       *string     `json:"alias,omitempty" gorm:"column:alias"`
	IsOffical   bool        `json:"is_offical" gorm:"column:is_offical"`
	Thumbnail   *string     `json:"thumbnail" gorm:"column:thumbnail"`
	ThumbnailM  *string     `json:"thumbnailM" gorm:"column:thumbnailM"`
	Duration    int         `json:"duration" gorm:"column:duration"`
	AlbumID     *int        `json:"album_id,omitempty" gorm:"column:album"`
	Album       *Albums     `json:"album,omitempty" gorm:"foreignKey:AlbumID;references:ID"`
	Artists     []Artists   `json:"artists,omitempty" gorm:"many2many:songs_artists"`
	Composers   []Composers `json:"composers,omitempty" gorm:"many2many:songs_composers"`
	Genres      []Genres    `json:"genres,omitempty" gorm:"many2many:songs_genres"`
	ReleaseDate int         `json:"release_date" gorm:"column:releaseDate"`
	Distributor *string     `json:"distributor" gorm:"column:distributor"`
	HasLyric    bool        `json:"has_lyric" gorm:"column:hasLyric"`
	Like        int         `json:"like" gorm:"column:like"`
	Listen      int         `json:"listen" gorm:"column:listen"`
	Comment     int         `json:"comment" gorm:"column:comment"`
	AudioFile   string      `json:"audio_file" gorm:"column:audio_file"`
}

func (Songs) TableName() string {
	return "songs"
}

type Albums struct {
	ID          int       `json:"-" gorm:"primaryKey;column:id"`
	MaskId      string    `json:"mask_id" gorm:"column:mask_id"`
	Title       string    `json:"title" gorm:"column:title"`
	IsOffical   bool      `json:"is_offical" gorm:"column:is_offical"`
	Thumbnail   *string   `json:"thumbnail" gorm:"column:thumbnail"`
	Description string    `json:"description" gorm:"column:sortDescription"`
	ReleaseAt   int       `json:"release_at" gorm:"column:release_at"`
	Artists     []Artists `json:"artists,omitempty" gorm:"many2many:albums_artists"`
	Genres      []Genres  `json:"genres,omitempty" gorm:"many2many:songs_genres"`
}

func (Albums) TableName() string {
	return "albums"
}

type Genres struct {
	ID               int     `json:"-" gorm:"primaryKey;column:id"`
	MaskId           string  `json:"mask_id" gorm:"column:mask_id"`
	Title            string  `json:"title" gorm:"column:title"`
	Name             *string `json:"name" gorm:"column:name"`
	Alias            *string `json:"alias" gorm:"column:alias"`
	Thumbnail        *string `json:"thumbnail" gorm:"-"`
	ThumbnailHasText *string `json:"thumbnailHasText" gorm:"-"`
	ThumbnailR       *string `json:"thumbnailR" gorm:"-"`
}

func (Genres) TableName() string {
	return "genres"
}
