package zingmp3

import (
	"strconv"

	songmodels "example.com/musicafy_be/modules/songs/models"
)

type StreamingSongRes struct {
	File128 string `json:"128"`
	File320 string `json:"320"`
}

type resSuggestion struct {
	Err     int    `json:"err"`
	Message string `json:"msg"`
	Data    struct {
		Items []struct {
			Suggestions []SongSearch `json:"suggestions"`
		} `json:"items"`
	} `json:"data"`
}

type Artists struct {
	Type       int       `json:"type"`
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	AliasName  string    `json:"aliasName"`
	Alias      *string   `json:"alias"`
	Cover      *string   `json:"cover"`
	Thumbnail  *string   `json:"thumbnail"`
	ThumbnailM *string   `json:"thumbnailM"`
	Avatar     string    `json:"avatar"`
	PlaylistId string    `json:"playlistId"`
	Followers  int       `json:"follow"`
	Sections   []Section `json:"sections"`
}

func (a *Artists) ToModelDb() songmodels.Artists {
	var sections []songmodels.Section
	for _, v := range a.Sections {
		sections = append(sections, v.ToModelDb())
	}
	thumb := &a.Avatar
	if a.Thumbnail != nil {
		thumb = a.Thumbnail
	}
	alias := a.AliasName
	if a.Alias != nil {
		alias = *a.Alias
	}
	return songmodels.Artists{
		MaskId:      a.Id,
		Name:        a.Name,
		Thumbnail:   thumb,
		ThumbnailM:  a.ThumbnailM,
		PlaylistId:  &a.PlaylistId,
		TotalFollow: a.Followers,
		Sections:    &sections,
		Alias:       alias,
	}
}

type Section struct {
	Items []Song `json:"items"`
	Type  string `json:"sectionType"`
	Title string `json:"title"`
}

func (s *Section) ToModelDb() songmodels.Section {
	var items []songmodels.Songs
	for _, v := range s.Items {
		items = append(items, v.ToModelDb())
	}
	return songmodels.Section{
		Items: items,
		Type:  s.Type,
		Title: s.Title,
	}
}

type ArtistsSearch struct {
	Type       int     `json:"type"`
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	AliasName  string  `json:"aliasName"`
	Thumbnail  *string `json:"thumbnail"`
	ThumbnailM *string `json:"thumbnailM"`
	Avatar     string  `json:"avatar"`
	PlaylistId int     `json:"playlistId"`
	Followers  int     `json:"followers"`
}

func (a *ArtistsSearch) ToModelDb() songmodels.Artists {
	playlistId := strconv.Itoa(a.PlaylistId)
	return songmodels.Artists{
		MaskId:      a.Id,
		Name:        a.Name,
		Thumbnail:   &a.Avatar,
		PlaylistId:  &playlistId,
		TotalFollow: a.Followers,
	}
}

type Genres struct {
	Id               string  `json:"id"`
	EncodeId         *string `json:"encodeId,omitempty"`
	Name             string  `json:"name"`
	Title            *string `json:"title,omitempty"`
	ThumbS           string  `json:"thumbS"`
	Alias            string  `json:"alias"`
	Cover            string  `json:"cover"`
	Thumbnail        *string `json:"thumbnail,omitempty"`
	ThumbnailR       *string `json:"thumbnailR,omitempty"`
	ThumbnailHasText *string `json:"thumbnailHasText,omitempty"`
}

func (g *Genres) ToModelDb() songmodels.Genres {
	name := g.Name
	if g.Title != nil {
		name = *g.Title
	}

	id := g.Id
	if g.EncodeId != nil {
		id = *g.EncodeId
	}

	return songmodels.Genres{
		MaskId:           id,
		Title:            name,
		Name:             &name,
		Thumbnail:        g.Thumbnail,
		ThumbnailR:       g.ThumbnailR,
		ThumbnailHasText: g.ThumbnailHasText,
	}
}

type SongSearch struct {
	Type           int    `json:"type"`
	Title          string `json:"title"`
	ID             string `json:"id"`
	RadioPid       string `json:"radioPid"`
	HasVideo       bool   `json:"hasVideo"`
	Thumb          string `json:"thumb"`
	ThumbVideo     string `json:"thumbVideo"`
	Duration       int    `json:"duration"`
	Link           string `json:"link"`
	ModifiedTime   int64  `json:"modifiedTime"`
	LyricLink      string `json:"lyricLink"`
	LyricID        string `json:"lyricId"`
	DownloadTypes  string `json:"downloadTypes"`
	OrgMD5         string `json:"orgMD5"`
	UserID         int    `json:"userId"`
	EuID           string `json:"euId"`
	Privacy        int    `json:"privacy"`
	HLyricVersion  int64  `json:"hLyricVersion"`
	ReleaseTime    int64  `json:"releaseTime"`
	DownloadStatus int    `json:"downloadStatus"`
	Status         int    `json:"status"`
	PlayStatus     int    `json:"playStatus"`
	// Video          Video     `json:"video"`
	Artists      []ArtistsSearch `json:"artists"`
	Genres       []Genres        `json:"genres"`
	DisSPlatform int             `json:"disSPlatform"`
	DisDPlatform int             `json:"disDPlatform"`
	BoolAtt      int             `json:"boolAtt"`
	Tracking     string          `json:"tracking"`
}

func (s *SongSearch) ToModelDb() songmodels.Songs {
	artists := make([]songmodels.Artists, len(s.Artists))
	for i, artist := range s.Artists {
		artists[i] = artist.ToModelDb()
	}
	genres := make([]songmodels.Genres, len(s.Genres))
	for i, genre := range s.Genres {
		genres[i] = genre.ToModelDb()
	}
	return songmodels.Songs{
		MaskId:    s.ID,
		Title:     s.Title,
		Thumbnail: &s.Thumb,
		Duration:  s.Duration,
		Artists:   artists,
		Genres:    genres,
	}
}

type Song struct {
	Type           int       `json:"type"`
	Title          string    `json:"title"`
	ID             string    `json:"id"`
	EncodeId       *string   `json:"encodeId"`
	RadioPid       string    `json:"radioPid"`
	HasVideo       bool      `json:"hasVideo"`
	Thumb          string    `json:"thumb"`
	Thumbnail      *string   `json:"thumbnail"`
	ThumbnailM     *string   `json:"thumbnailM"`
	ThumbVideo     string    `json:"thumbVideo"`
	Duration       int       `json:"duration"`
	Link           string    `json:"link"`
	ModifiedTime   int64     `json:"modifiedTime"`
	LyricLink      string    `json:"lyricLink"`
	LyricID        string    `json:"lyricId"`
	DownloadTypes  string    `json:"downloadTypes"`
	OrgMD5         string    `json:"orgMD5"`
	UserID         int       `json:"userId"`
	EuID           string    `json:"euId"`
	Privacy        int       `json:"privacy"`
	HLyricVersion  int64     `json:"hLyricVersion"`
	ReleaseTime    int64     `json:"releaseTime"`
	DownloadStatus int       `json:"downloadStatus"`
	Status         int       `json:"status"`
	PlayStatus     int       `json:"playStatus"`
	Artists        []Artists `json:"artists"`
	Genres         []Genres  `json:"genres"`
	DisSPlatform   int       `json:"disSPlatform"`
	DisDPlatform   int       `json:"disDPlatform"`
	BoolAtt        int       `json:"boolAtt"`
	Tracking       string    `json:"tracking"`
}

func (s *Song) ToModelDb() songmodels.Songs {
	artists := make([]songmodels.Artists, len(s.Artists))
	for i, artist := range s.Artists {
		artists[i] = artist.ToModelDb()
	}
	genres := make([]songmodels.Genres, len(s.Genres))
	for i, genre := range s.Genres {
		genres[i] = genre.ToModelDb()
	}
	markId := s.ID
	if s.EncodeId != nil {
		markId = *s.EncodeId
	}
	thumb := &s.Thumb
	if s.Thumbnail != nil {
		thumb = s.Thumbnail
	}
	return songmodels.Songs{
		MaskId:     markId,
		Title:      s.Title,
		Thumbnail:  thumb,
		Duration:   s.Duration,
		Artists:    artists,
		Genres:     genres,
		ThumbnailM: s.ThumbnailM,
	}
}

type SongDetail struct {
	EncodeId     string    `json:"encodeId"`
	Title        string    `json:"title"`
	Alias        string    `json:"alias"`
	IsOffical    bool      `json:"isOffical"`
	Username     string    `json:"username"`
	ArtistsNames string    `json:"artistsNames"`
	Artists      []Artists `json:"artists"`
	IsWorldWide  bool      `json:"isWorldWide"`
	// PreviewInfo      PreviewInfo     `json:"previewInfo"`
	ThumbnailM         string        `json:"thumbnailM"`
	Link               string        `json:"link"`
	Thumbnail          string        `json:"thumbnail"`
	Duration           int           `json:"duration"`
	ZingChoice         bool          `json:"zingChoice"`
	IsPrivate          bool          `json:"isPrivate"`
	PreRelease         bool          `json:"preRelease"`
	ReleaseDate        int64         `json:"releaseDate"`
	GenreIds           []string      `json:"genreIds"`
	Distributor        string        `json:"distributor"`
	Indicators         []interface{} `json:"indicators"`
	IsIndie            bool          `json:"isIndie"`
	StreamingStatus    int           `json:"streamingStatus"`
	StreamPrivileges   []int         `json:"streamPrivileges"`
	DownloadPrivileges []int         `json:"downloadPrivileges"`
	AllowAudioAds      bool          `json:"allowAudioAds"`
	HasLyric           bool          `json:"hasLyric"`
	UserId             int           `json:"userid"`
	Genres             []Genres      `json:"genres"`
	Composers          []Composers   `json:"composers"`
	Album              Albums        `json:"album"`
	IsRBT              bool          `json:"isRBT"`
	Like               int           `json:"like"`
	Listen             int           `json:"listen"`
	Liked              bool          `json:"liked"`
	Comment            int           `json:"comment"`
}

// Composer represents a song's composer
type Composers struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Link      string `json:"link"`
	Spotlight bool   `json:"spotlight"`
	Alias     string `json:"alias"`
	Cover     string `json:"cover"`
	Thumbnail string `json:"thumbnail"`
}

// Album represents information about a song's album
type Albums struct {
	EncodeId        string    `json:"encodeId"`
	Title           string    `json:"title"`
	Thumbnail       string    `json:"thumbnail"`
	IsOffical       bool      `json:"isoffical"`
	Link            string    `json:"link"`
	IsIndie         bool      `json:"isIndie"`
	ReleaseDate     string    `json:"releaseDate"`
	SortDescription string    `json:"sortDescription"`
	ReleasedAt      int64     `json:"releasedAt"`
	GenreIds        []string  `json:"genreIds"`
	PR              bool      `json:"PR"`
	Artists         []Artists `json:"artists"`
	ArtistsNames    string    `json:"artistsNames"`
}

type Lyric struct {
	Sentences      []Sentence `json:"sentences"`
	File           string     `json:"file"`
	EnabledVideoBG bool       `json:"enabledVideoBG"`
	DefaultIBGUrls []string   `json:"defaultIBGUrls"`
	BGMode         int        `json:"BGMode"`
}

type Sentence struct {
	Words []Word `json:"words"`
}

type Word struct {
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	Data      string `json:"data"`
}

type HubHomeResponse struct {
	Genres []Genres `json:"genre"`
}
