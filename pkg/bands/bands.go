package bands

import (
	"fmt"

	"github.com/antch57/jam-statz/graph/model"
	"gorm.io/gorm"
)

type BandRepo struct {
	DB *gorm.DB
}

// Band CRUD Operations
func (b *BandRepo) CreateBand(input *model.BandInput) (*model.Band, error) {
	bandInsert := &model.Band{
		Name:        input.Name,
		Genre:       input.Genre,
		Year:        input.Year,
		Description: input.Description,
	}

	// Create the band
	res := b.DB.Create(bandInsert)
	if res.Error != nil {
		return nil, res.Error
	}

	return bandInsert, nil
}

func (b *BandRepo) GetBands() ([]*model.Band, error) {
	var bandsList []*model.Band

	res := b.DB.Find(&bandsList)
	if res.Error != nil {
		return nil, res.Error
	}

	return bandsList, nil
}

func (b *BandRepo) GetBandById(id int) (*model.Band, error) {
	var band = &model.Band{}
	res := b.DB.First(&band, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return band, nil
}

// TODO: add test for this in song package
func (b *BandRepo) GetBandBySongId(id int) (*model.Band, error) {
	var band = &model.Band{}
	res := b.DB.Table("songs").Select("bands.id, bands.name, bands.genre, bands.year, bands.description").Joins("left join bands on bands.id = songs.band_id").Where("songs.id = ?", id).Scan(band)
	if res.Error != nil {
		return nil, res.Error
	}

	if band.ID == 0 {
		// TODO: Return a custom error
		fmt.Println("Band not found")
		return nil, nil
	}

	return band, nil
}

func (b *BandRepo) UpdateBand(id int, input *model.BandInput) (*model.Band, error) {
	var bandUpdate = &model.Band{
		ID:          id,
		Name:        input.Name,
		Genre:       input.Genre,
		Year:        input.Year,
		Description: input.Description,
	}

	res := b.DB.Model(&model.Band{}).Where("id = ?", id).Updates(bandUpdate)
	if res.Error != nil {
		return nil, res.Error
	}

	return bandUpdate, nil
}

func (b *BandRepo) DeleteBand(id int) (bool, error) {
	res := b.DB.Delete(&model.Band{}, id)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}
