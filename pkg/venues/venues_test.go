package venues

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/antch57/jam-statz/graph/model"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockDb() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating mock database: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("error opening Gorm database: %v", err)
	}

	return db, mock, nil
}

func TestVenueRepo_CreateVenue(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	input := &model.VenueInput{Name: "Test Venue", Location: "Denver, CO"}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `venues` \\(`name`,`location`\\) VALUES \\(\\?,\\?\\)$").WithArgs(input.Name, input.Location).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &VenueRepo{DB: db}

	venue, err := repo.CreateVenue(input)
	if err != nil {
		t.Errorf("Error creating venue: %v", err)
	}

	assert.Equal(t, input.Name, venue.Name)
	assert.Equal(t, input.Location, *venue.Location)

}

func TestVenueRepo_GetVenue(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "location"}).AddRow(1, "Test Venue", "Denver, CO")

	mock.ExpectQuery("^SELECT \\* FROM `venues` WHERE `venues`.`id` = \\? ORDER BY `venues`.`id` LIMIT 1$").WithArgs(1).WillReturnRows(rows)

	repo := &VenueRepo{DB: db}

	venue, err := repo.GetVenue(1)
	if err != nil {
		t.Errorf("Error getting venue: %v", err)
	}

	assert.Equal(t, 1, venue.ID)
	assert.Equal(t, "Test Venue", venue.Name)
	assert.Equal(t, "Denver, CO", *venue.Location)
}

func TestVenueRepo_GetVenues(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "location"}).AddRow(1, "Test Venue", "Denver, CO")

	mock.ExpectQuery("^SELECT \\* FROM `venues`$").WillReturnRows(rows)

	repo := &VenueRepo{DB: db}

	venues, err := repo.GetVenues()
	if err != nil {
		t.Errorf("Error getting venues: %v", err)
	}

	assert.Equal(t, 1, len(venues))
	assert.Equal(t, "Test Venue", venues[0].Name)
	assert.Equal(t, "Denver, CO", *venues[0].Location)
}

// 	assert.NoError(t, err)
// 	assert.NotNil(t, song)
// 	assert.Equal(t, "New Song!", song.Title)
// 	assert.Equal(t, testBandId, *song.BandID)
// }

func TestVenueRepo_UpdateVenue(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	venueId := 1
	input := &model.VenueInput{Name: "Test Venue", Location: "Denver, CO"}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `venues` SET `id`=\\?,`name`=\\?,`location`=\\? WHERE id = \\?$").WithArgs(venueId, input.Name, input.Location, venueId).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &VenueRepo{DB: db}

	res, err := repo.UpdateVenue(venueId, input)
	if err != nil {
		t.Errorf("Error updating venue: %v", err)
	}

	assert.Equal(t, venueId, res.ID)
	assert.Equal(t, input.Name, res.Name)
	assert.Equal(t, input.Location, *res.Location)
}
