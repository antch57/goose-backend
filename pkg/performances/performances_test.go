package performances

import (
	"fmt"
	"testing"
	"time"

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

// performance tests
func TestPerformanceRepo_CreatePerformance(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	parsedTime, err := time.Parse("2006-01-02", "2021-01-01")
	if err != nil {
		t.Errorf("Error parsing time: %v", err)
	}

	input := &model.PerformanceInput{
		BandID:          1,
		VenueID:         1,
		Duration:        60,
		PerformanceDate: parsedTime,
	}

	insertSQL := "^INSERT INTO `performances` \\(`band_id`,`venue_id`,`performance_date`,`duration`\\) VALUES \\(\\?,\\?,\\?,\\?\\)$"
	repo := &PerformanceRepo{DB: db}
	mock.ExpectBegin()
	mock.ExpectExec(insertSQL).WithArgs(input.BandID, input.VenueID, input.PerformanceDate, input.Duration).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res, err := repo.CreatePerformance(input)
	if err != nil {
		t.Errorf("Error creating performance: %v", err)
	}

	assert.NotEmpty(t, res.ID)
	assert.Equal(t, input.BandID, res.BandID)
	assert.Equal(t, input.VenueID, res.VenueID)
	assert.Equal(t, input.Duration, res.Duration)
	assert.Equal(t, input.PerformanceDate, res.PerformanceDate)
}

func TestPerformanceRepo_GetPerformance(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	input := &model.Performance{
		ID:              1,
		BandID:          1,
		VenueID:         1,
		Duration:        60,
		PerformanceDate: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "band_id", "venue_id", "performance_date", "duration"}).
		AddRow(input.ID, input.BandID, input.VenueID, input.PerformanceDate, input.Duration)

	selectSQL := "^SELECT \\* FROM `performances` WHERE `performances`.`id` = \\? ORDER BY `performances`.`id` LIMIT 1$"
	repo := &PerformanceRepo{DB: db}

	mock.ExpectQuery(selectSQL).WithArgs(input.ID).WillReturnRows(rows)

	res, err := repo.GetPerformance(input.ID)
	if err != nil {
		t.Errorf("Error getting performance: %v", err)
	}

	assert.NotEmpty(t, res.ID)
	assert.Equal(t, input.BandID, res.BandID)
	assert.Equal(t, input.VenueID, res.VenueID)
	assert.Equal(t, input.Duration, res.Duration)
	assert.Equal(t, input.PerformanceDate, res.PerformanceDate)
}

func TestPerformanceRepo_GetPerformances(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	input := []*model.Performance{
		{
			ID:              1,
			BandID:          1,
			VenueID:         1,
			Duration:        60,
			PerformanceDate: time.Now(),
		},
		{
			ID:              2,
			BandID:          1,
			VenueID:         2,
			Duration:        500,
			PerformanceDate: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "band_id", "venue_id", "performance_date", "duration"}).
		AddRow(input[0].ID, input[0].BandID, input[0].VenueID, input[0].PerformanceDate, input[0].Duration).
		AddRow(input[1].ID, input[1].BandID, input[1].VenueID, input[1].PerformanceDate, input[1].Duration)

	selectSQL := "^SELECT \\* FROM `performances`$"
	repo := &PerformanceRepo{DB: db}

	mock.ExpectQuery(selectSQL).WillReturnRows(rows)

	res, err := repo.GetPerformances()
	if err != nil {
		t.Errorf("Error getting performances: %v", err)
	}

	assert.NotEmpty(t, res)
	assert.Len(t, res, 2)
	assert.Equal(t, input[0].BandID, res[0].BandID)
	assert.Equal(t, input[0].VenueID, res[0].VenueID)
	assert.Equal(t, input[0].Duration, res[0].Duration)
	assert.Equal(t, input[0].PerformanceDate, res[0].PerformanceDate)
	assert.Equal(t, input[1].BandID, res[1].BandID)
	assert.Equal(t, input[1].VenueID, res[1].VenueID)
	assert.Equal(t, input[1].Duration, res[1].Duration)
	assert.Equal(t, input[1].PerformanceDate, res[1].PerformanceDate)
}

func TestPerformanceRepo_GetPerformancesByVenueID(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	venueId := 1
	input := []*model.Performance{
		{
			ID:              1,
			BandID:          1,
			VenueID:         venueId,
			Duration:        60,
			PerformanceDate: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "band_id", "venue_id", "performance_date", "duration"}).
		AddRow(input[0].ID, input[0].BandID, input[0].VenueID, input[0].PerformanceDate, input[0].Duration)

	selectSQL := "^SELECT \\* FROM `performances` WHERE venue_id = \\?$"
	repo := &PerformanceRepo{DB: db}

	mock.ExpectQuery(selectSQL).WithArgs(venueId).WillReturnRows(rows)

	res, err := repo.GetPerformancesByVenueID(venueId)
	if err != nil {
		t.Errorf("Error getting performances: %v", err)
	}

	assert.NotEmpty(t, res)
	assert.Len(t, res, 1)
	assert.Equal(t, input[0].BandID, res[0].BandID)
	assert.Equal(t, input[0].VenueID, res[0].VenueID)
	assert.Equal(t, input[0].Duration, res[0].Duration)
	assert.Equal(t, input[0].PerformanceDate, res[0].PerformanceDate)
}

func TestPerformanceRepo_UpdatePerformance(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testId := 1
	input := &model.PerformanceInput{
		BandID:          1,
		VenueID:         1,
		Duration:        60,
		PerformanceDate: time.Now(),
	}

	updateSQL := "^UPDATE `performances` SET `band_id`=\\?,`venue_id`=\\?,`performance_date`=\\?,`duration`=\\? WHERE `id` = \\?$"
	repo := &PerformanceRepo{DB: db}
	mock.ExpectBegin()
	mock.ExpectExec(updateSQL).WithArgs(input.BandID, input.VenueID, input.PerformanceDate, input.Duration, testId).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res, err := repo.UpdatePerformance(testId, input)
	if err != nil {
		t.Errorf("Error updating performance: %v", err)
	}

	assert.Equal(t, testId, res.ID)
	assert.Equal(t, input.BandID, res.BandID)
	assert.Equal(t, input.VenueID, res.VenueID)
	assert.Equal(t, input.Duration, res.Duration)
	assert.Equal(t, input.PerformanceDate, res.PerformanceDate)
}

func TestPerformanceRepo_DeletePerformance(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testId := 1
	deleteSQL := "^DELETE FROM `performances` WHERE `performances`.`id` = \\?$"
	repo := &PerformanceRepo{DB: db}
	mock.ExpectBegin()
	mock.ExpectExec(deleteSQL).WithArgs(testId).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res, err := repo.DeletePerformance(testId)
	if err != nil {
		t.Errorf("Error deleting performance: %v", err)
	}

	assert.True(t, res)
}
