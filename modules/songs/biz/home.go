package songsbiz

import (
	"example.com/musicafy_be/components/appctx"
	songmodels "example.com/musicafy_be/modules/songs/models"
)

type HomeStorage interface {
}

type HomeBiz struct {
	appctx appctx.AppContext
}

func NewHomeBiz(appctx appctx.AppContext) *HomeBiz {
	return &HomeBiz{
		appctx: appctx,
	}
}

func (h *HomeBiz) GetHomeSession() (*HomeResponse, error) {
	zingmp3 := h.appctx.GetZingmp3Api()
	sections, err := zingmp3.Home()
	if err != nil {
		return nil, err
	}

	var sectionsFilter []songmodels.HomeSection
	for _, v := range sections {
		if v.SectionType == "newReleaseChart" || v.SectionType == "new-release" {
			sectionsFilter = append(sectionsFilter, v)
		}
	}

	hubHome, err := zingmp3.HubHome()
	if err != nil {
		return nil, err
	}

	var genres []songmodels.Genres
	for _, v := range hubHome.Genres {
		genres = append(genres, v.ToModelDb())
	}

	return &HomeResponse{
		Sections: sectionsFilter,
		Genres:   genres,
	}, nil
}

type HomeResponse struct {
	Sections []songmodels.HomeSection `json:"sections"`
	Genres   []songmodels.Genres      `json:"genres"`
}
