package model

import "time"

type Performance struct {
	ID              int       `json:"id"`
	BandID          int       `json:"band"`
	VenueID         int       `json:"venue"`
	PerformanceDate time.Time `json:"performanceDate"`
	// FIXME: should be time.Duration
	// Duration        time.Duration      `json:"duration"`
	Duration int                `json:"duration"`
	Songs    []*PerformanceSong `json:"songs"`
}

type PerformanceInput struct {
	BandID          int       `json:"bandId"`
	VenueID         int       `json:"venue"`
	PerformanceDate time.Time `json:"performanceDate"`
	// FIXME: should be time.Duration
	// Duration         *time.Duration          `json:"duration,omitempty"`
	Duration         int                     `json:"duration,omitempty"`
	PerformanceSongs []*PerformanceSongInput `json:"performanceSongs,omitempty"`
}

type PerformanceSong struct {
	ID            int           `json:"id"`
	SongID        int           `json:"song"`
	Duration      time.Duration `json:"duration"`
	PerformanceID int           `json:"performance"`
	IsCover       bool          `json:"isCover"`
	Notes         *string       `json:"notes,omitempty"`
}

type PerformanceSongInput struct {
	SongID        int           `json:"songId"`
	Duration      time.Duration `json:"duration"`
	PerformanceID int           `json:"performanceId"`
	Notes         *string       `json:"notes,omitempty"`
	IsCover       bool          `json:"isCover"`
}
