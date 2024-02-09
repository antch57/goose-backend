package venues

import (
	"fmt"

	"github.com/antch57/jam-statz/graph/model"
	"gorm.io/gorm"
)

type VenueRepo struct {
	DB *gorm.DB
}

// Venue CRUD
func (v *VenueRepo) CreateVenue(input *model.VenueInput) (*model.Venue, error) {
	// Set location to "Unknown" if it's empty
	location := new(string)
	if input.Location == "" {
		input.Location = "Unknown"
	} else {
		*location = input.Location
	}

	venue := &model.Venue{
		Name:     input.Name,
		Location: location,
	}

	res := v.DB.Create(venue)
	if res.Error != nil {
		return nil, res.Error
	}

	return venue, nil
}

func (v *VenueRepo) GetVenue(venueId int) (*model.Venue, error) {
	venue := &model.Venue{}

	res := v.DB.First(&venue, venueId)
	if res.Error != nil {
		return nil, res.Error
	}

	return venue, nil
}

func (v *VenueRepo) GetVenues() ([]*model.Venue, error) {
	venues := []*model.Venue{}

	res := v.DB.Find(&venues)
	if res.Error != nil {
		return nil, res.Error
	}

	return venues, nil
}

func (v *VenueRepo) UpdateVenue(venueId int, input *model.VenueInput) (*model.Venue, error) {
	venue := &model.Venue{
		ID:       venueId,
		Name:     input.Name,
		Location: &input.Location,
	}

	res := v.DB.Model(&model.Venue{}).Where("id = ?", venueId).Updates(venue)
	if res.Error != nil {
		return nil, res.Error
	}
	fmt.Println("i think this is ID: ?", venue.ID)
	return venue, nil
}

func (v *VenueRepo) DeleteVenue(venueId int) (bool, error) {
	res := v.DB.Delete(&model.Venue{}, venueId)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}
