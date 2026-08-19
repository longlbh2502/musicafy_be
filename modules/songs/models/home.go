package songmodels

type HomeSection struct {
	Banner      string `json:"banner"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	SectionType string `json:"sectionType"`
	SectionId   string `json:"sectionId"`
	Items       any    `json:"items"`
}
