package model

func (Venue) TableName() string {
	return "venues"
}

type Venue struct {
	ID           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Location     *string        `json:"location,omitempty"`
	Performances []*Performance `json:"performances,omitempty"`
}

type VenueInput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}
