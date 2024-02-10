package performances

import "github.com/antch57/jam-statz/graph/model"

// performance_song CRUD
// Create
func (p *PerformanceRepo) CreatePerformanceSong(input *model.PerformanceSongInput) (*model.PerformanceSong, error) {
	performanceSong := &model.PerformanceSong{
		SongID:        input.SongID,
		Duration:      input.Duration,
		PerformanceID: input.PerformanceID,
		Notes:         input.Notes,
		IsCover:       input.IsCover,
	}

	res := p.DB.Create(performanceSong)
	if res.Error != nil {
		return nil, res.Error
	}

	return performanceSong, nil
}

// READ
func (p *PerformanceRepo) GetPerformanceSong(performanceSongID int) (*model.PerformanceSong, error) {
	performanceSong := &model.PerformanceSong{}

	res := p.DB.First(&performanceSong, performanceSongID)
	if res.Error != nil {
		return nil, res.Error
	}

	return performanceSong, nil
}

func (p *PerformanceRepo) GetPerformanceSongs(performanceID int) ([]*model.PerformanceSong, error) {
	performanceSongs := []*model.PerformanceSong{}

	res := p.DB.Where("performance_id = ?", performanceID).Find(&performanceSongs)
	if res.Error != nil {
		return nil, res.Error
	}

	return performanceSongs, nil
}

func (p *PerformanceRepo) GetPerformanceSongsAll() ([]*model.PerformanceSong, error) {
	performanceSongs := []*model.PerformanceSong{}

	res := p.DB.Find(&performanceSongs)
	if res.Error != nil {
		return nil, res.Error
	}

	return performanceSongs, nil
}

// UPDATE
func (p *PerformanceRepo) UpdatePerformanceSong(perfomanceSongId int, input *model.PerformanceSongInput) (*model.PerformanceSong, error) {
	performanceSong := &model.PerformanceSong{
		ID:            perfomanceSongId,
		SongID:        input.SongID,
		Duration:      input.Duration,
		PerformanceID: input.PerformanceID,
		Notes:         input.Notes,
		IsCover:       input.IsCover,
	}

	res := p.DB.Save(performanceSong)
	if res.Error != nil {
		return nil, res.Error
	}

	return performanceSong, nil
}

// DELETE
func (p *PerformanceRepo) DeletePerformanceSong(performanceSongID int) (bool, error) {
	res := p.DB.Delete(&model.PerformanceSong{}, performanceSongID)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}
