package model

func (Band) TableName() string {
	return "bands"
}

type Band struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Description *string `json:"description,omitempty"`
}

type BandInput struct {
	Name        string        `json:"name"`
	Genre       string        `json:"genre"`
	Year        int           `json:"year"`
	Description *string       `json:"description,omitempty"`
	Albums      []*AlbumInput `json:"albums,omitempty"`
}
