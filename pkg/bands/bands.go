package bands

import (
	"fmt"

	"github.com/antch57/goose/graph/model"
	"gorm.io/gorm"
)

type BandRepo struct {
	DB *gorm.DB
}

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

func (b *BandRepo) GetBandByID(id int) (*model.Band, error) {
	var band = &model.Band{}
	res := b.DB.First(&band, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return band, nil
}

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
