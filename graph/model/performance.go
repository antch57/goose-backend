package model

import "time"

// FIXME: Duration fields should be "time.Duration" type
type Performance struct {
	ID              int       `json:"id"`
	BandID          int       `json:"band"`
	VenueID         int       `json:"venue"`
	PerformanceDate time.Time `json:"performanceDate"`
	Duration        int       `json:"duration"`
}

type PerformanceInput struct {
	BandID          int       `json:"bandId"`
	VenueID         int       `json:"venue"`
	PerformanceDate time.Time `json:"performanceDate"`
	Duration        int       `json:"duration,omitempty"`
}

type PerformanceSong struct {
	ID            int     `json:"id"`
	SongID        int     `json:"song"`
	Duration      int     `json:"duration"`
	PerformanceID int     `json:"performance"`
	IsCover       bool    `json:"isCover"`
	Notes         *string `json:"notes,omitempty"`
}

type PerformanceSongInput struct {
	SongID        int     `json:"songId"`
	Duration      int     `json:"duration"`
	PerformanceID int     `json:"performanceId"`
	Notes         *string `json:"notes,omitempty"`
	IsCover       bool    `json:"isCover"`
}
