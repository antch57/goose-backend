package model

func (Song) TableName() string {
	return "songs"
}

type Song struct {
	ID     int    `json:"id"`
	BandID *int   `json:"bandId"`
	Title  string `json:"title"`
}

type SongInput struct {
	Title  string `json:"title"`
	BandID *int   `json:"bandId,omitempty"`
}
