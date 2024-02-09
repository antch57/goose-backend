package model

import "time"

func (Album) TableName() string {
	return "albums"
}

type Album struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	BandID      int       `json:"bandId"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type AlbumInput struct {
	Title       string    `json:"title"`
	BandID      int       `json:"bandId"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type AlbumSong struct {
	ID          int   `json:"id"`
	SongID      int   `json:"songId"`
	AlbumID     int   `json:"albumId"`
	Duration    int   `json:"duration"` // FIXME: should be time.Duration
	TrackNumber int   `json:"track_number"`
	IsCover     *bool `json:"isCover,omitempty"`
}

type AlbumSongInput struct {
	SongID      int  `json:"songId"`
	AlbumID     int  `json:"albumId"`
	Duration    int  `json:"duration"` // FIXME: should be time.Duration
	TrackNumber int  `json:"trackNumber"`
	IsCover     bool `json:"isCover"`
}
