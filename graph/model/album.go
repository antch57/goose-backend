package model

import "time"

type Album struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	BandID      string    `json:"bandId"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type AlbumInput struct {
	Title       string    `json:"title"`
	BandID      string    `json:"bandId"`
	ReleaseDate time.Time `json:"releaseDate"`
}

type AlbumSong struct {
	ID          string        `json:"id"`
	SongID      string        `json:"songId"`
	AlbumID     string        `json:"albumId"`
	BandID      string        `json:"bandId"`
	Duration    time.Duration `json:"duration"`
	TrackNumber int           `json:"track_number"`
	IsCover     *bool         `json:"isCover,omitempty"`
}

type AlbumSongInput struct {
	SongID      string        `json:"songId"`
	AlbumID     string        `json:"albumId"`
	BandID      string        `json:"bandId"`
	Duration    time.Duration `json:"duration"`
	TrackNumber int           `json:"trackNumber"`
	IsCover     bool          `json:"isCover"`
}
