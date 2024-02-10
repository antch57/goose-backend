package performances

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/antch57/jam-statz/graph/model"
	"github.com/stretchr/testify/assert"
)

// performance_song tests
func TestPerformanceRepo_CreatePerformanceSong(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	input := &model.PerformanceSongInput{
		PerformanceID: 1,
		SongID:        1,
	}

	insertSQL := "^INSERT INTO `performance_songs` \\(`song_id`,`duration`,`performance_id`,`is_cover`,`notes`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?\\)$"
	repo := &PerformanceRepo{DB: db}

	mock.ExpectBegin()
	mock.ExpectExec(insertSQL).WithArgs(input.SongID, input.Duration, input.PerformanceID, input.IsCover, input.Notes).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res, err := repo.CreatePerformanceSong(input)
	if err != nil {
		t.Errorf("Error creating performance song: %v", err)
	}

	assert.NotEmpty(t, res.ID)
	assert.Equal(t, input.PerformanceID, res.PerformanceID)
	assert.Equal(t, input.SongID, res.SongID)
	assert.Equal(t, input.Duration, res.Duration)
	assert.Equal(t, input.IsCover, res.IsCover)
	assert.Equal(t, input.Notes, res.Notes)
}

func TestPerformanceRepo_GetPerformanceSong(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	inputNotes := new(string)
	*inputNotes = "This is a note"

	input := &model.PerformanceSong{
		ID:            1,
		PerformanceID: 1,
		SongID:        1,
		Duration:      300,
		IsCover:       false,
		Notes:         inputNotes,
	}

	repo := &PerformanceRepo{DB: db}

	rows := sqlmock.NewRows([]string{"id", "performance_id", "song_id", "duration", "is_cover", "notes"}).
		AddRow(input.ID, input.PerformanceID, input.SongID, input.Duration, input.IsCover, input.Notes)

	mock.ExpectQuery("^SELECT \\* FROM `performance_songs` WHERE `performance_songs`.`id` = \\? ORDER BY `performance_songs`.`id` LIMIT 1$").WithArgs(input.ID).WillReturnRows(rows)

	res, err := repo.GetPerformanceSong(input.ID)
	if err != nil {
		t.Errorf("Error getting performance song: %v", err)
	}

	assert.NotEmpty(t, res.ID)
	assert.Equal(t, input.PerformanceID, res.PerformanceID)
	assert.Equal(t, input.SongID, res.SongID)
	assert.Equal(t, input.Duration, res.Duration)
	assert.Equal(t, input.IsCover, res.IsCover)
	assert.Equal(t, input.Notes, res.Notes)
}

func TestPerformanceRepo_GetPerformanceSongs(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	performanceID := 1
	repo := &PerformanceRepo{DB: db}

	rows := sqlmock.NewRows([]string{"id", "performance_id", "song_id", "duration", "is_cover", "notes"}).
		AddRow(1, performanceID, 1, 300, false, nil).
		AddRow(2, performanceID, 2, 300, false, nil).
		AddRow(3, performanceID, 3, 300, false, nil)

	mock.ExpectQuery("^SELECT \\* FROM `performance_songs` WHERE performance_id = \\?$").WithArgs(performanceID).WillReturnRows(rows)

	res, err := repo.GetPerformanceSongs(performanceID)
	if err != nil {
		t.Errorf("Error getting performance songs: %v", err)
	}

	assert.Len(t, res, 3)
}

func TestPerformanceRepo_GetPerformanceSongsAll(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	repo := &PerformanceRepo{DB: db}

	rows := sqlmock.NewRows([]string{"id", "performance_id", "song_id", "duration", "is_cover", "notes"}).
		AddRow(1, 1, 1, 300, false, nil).
		AddRow(2, 1, 2, 300, false, nil).
		AddRow(3, 1, 3, 300, false, nil)

	mock.ExpectQuery("^SELECT \\* FROM `performance_songs`").WillReturnRows(rows)

	res, err := repo.GetPerformanceSongsAll()
	if err != nil {
		t.Errorf("Error getting performance songs: %v", err)
	}

	assert.Len(t, res, 3)
}

func TestPerformanceRepo_UpdatePerformanceSong(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testId := 1
	input := &model.PerformanceSongInput{
		PerformanceID: 1,
		SongID:        1,
	}

	updateSQL := "^UPDATE `performance_songs` SET `song_id`=\\?,`duration`=\\?,`performance_id`=\\?,`is_cover`=\\?,`notes`=\\? WHERE `id` = \\?$"
	repo := &PerformanceRepo{DB: db}

	mock.ExpectBegin()
	mock.ExpectExec(updateSQL).
		WithArgs(input.SongID, input.Duration, input.PerformanceID, input.IsCover, input.Notes, testId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res, err := repo.UpdatePerformanceSong(1, input)
	if err != nil {
		t.Errorf("Error updating performance song: %v", err)
	}

	assert.NotEmpty(t, res.ID)
	assert.Equal(t, input.PerformanceID, res.PerformanceID)
	assert.Equal(t, input.SongID, res.SongID)
	assert.Equal(t, input.Duration, res.Duration)
	assert.Equal(t, input.IsCover, res.IsCover)
	assert.Equal(t, input.Notes, res.Notes)
}

func TestPerformanceRepo_DeletePerformanceSong(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testId := 1
	repo := &PerformanceRepo{DB: db}

	deleteSQL := "^DELETE FROM `performance_songs` WHERE `performance_songs`.`id` = \\?$"
	mock.ExpectBegin()
	mock.ExpectExec(deleteSQL).WithArgs(testId).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	res, err := repo.DeletePerformanceSong(testId)
	if err != nil {
		t.Errorf("Error deleting performance song: %v", err)
	}

	assert.True(t, res)
}
