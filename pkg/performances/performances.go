package performances

import (
	"github.com/antch57/jam-statz/graph/model"
	"gorm.io/gorm"
)

type PerformanceRepo struct {
	DB *gorm.DB
}

// performance CRUD
// Create
func (p *PerformanceRepo) CreatePerformance(input *model.PerformanceInput) (*model.Performance, error) {
	performance := &model.Performance{
		BandID:          input.BandID,
		VenueID:         input.VenueID,
		PerformanceDate: input.PerformanceDate,
		Duration:        input.Duration,
	}

	res := p.DB.Create(performance)
	if res.Error != nil {
		return nil, res.Error
	}

	return performance, nil
}

// READ
func (p *PerformanceRepo) GetPerformance(performanceId int) (*model.Performance, error) {
	performance := &model.Performance{}

	res := p.DB.First(&performance, performanceId)
	if res.Error != nil {
		return nil, res.Error
	}

	return performance, nil
}

func (p *PerformanceRepo) GetPerformances() ([]*model.Performance, error) {
	performances := []*model.Performance{}

	res := p.DB.Find(&performances)
	if res.Error != nil {
		return nil, res.Error
	}

	return performances, nil
}

func (p *PerformanceRepo) GetPerformancesByVenueID(venueID int) ([]*model.Performance, error) {
	performances := []*model.Performance{}

	res := p.DB.Where("venue_id = ?", venueID).Find(&performances)
	if res.Error != nil {
		return nil, res.Error
	}

	return performances, nil
}

// UPDATE
func (p *PerformanceRepo) UpdatePerformance(performanceId int, input *model.PerformanceInput) (*model.Performance, error) {
	performance := &model.Performance{
		ID:              performanceId,
		BandID:          input.BandID,
		VenueID:         input.VenueID,
		PerformanceDate: input.PerformanceDate,
		Duration:        input.Duration,
	}

	res := p.DB.Save(performance)
	if res.Error != nil {
		return nil, res.Error
	}

	return performance, nil
}

// DELETE
func (p *PerformanceRepo) DeletePerformance(performanceId int) (bool, error) {
	res := p.DB.Delete(&model.Performance{}, performanceId)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

// TODO: tests for all 3 of these
// TODO: Should i just pull these from the song and band repos?
// utility
func (p *PerformanceRepo) GetBand(bandID int) (*model.Band, error) {
	band := &model.Band{}

	res := p.DB.First(&band, bandID)
	if res.Error != nil {
		return nil, res.Error
	}

	return band, nil
}

func (p *PerformanceRepo) GetVenue(venueID int) (*model.Venue, error) {
	venue := &model.Venue{}

	res := p.DB.First(&venue, venueID)
	if res.Error != nil {
		return nil, res.Error
	}

	return venue, nil
}

func (p *PerformanceRepo) GetSong(songID int) (*model.Song, error) {
	song := &model.Song{}

	res := p.DB.First(&song, songID)
	if res.Error != nil {
		return nil, res.Error
	}

	return song, nil
}
