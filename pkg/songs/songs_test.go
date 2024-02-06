package songs

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/antch57/goose/graph/model"
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

// Create
func TestSongRepo_CreateSong(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `songs`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &SongRepo{DB: db}

	input := &model.SongInput{
		Title:  "Test Song",
		BandID: nil,
	}

	song, err := repo.CreateSong(input)

	assert.NoError(t, err)
	assert.Equal(t, song.Title, input.Title)
	assert.Equal(t, song.BandID, input.BandID)
}

// Read
func TestSongRepo_GetSongs(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "band_id"}).
		AddRow(1, "Test Song 1", nil).
		AddRow(2, "Test Song 2", nil)

	mock.ExpectQuery("^SELECT (.+) FROM `songs`$").WillReturnRows(rows)
	mock.ExpectCommit()

	repo := &SongRepo{DB: db}
	songs, err := repo.GetSongs()

	assert.NoError(t, err)
	assert.NotNil(t, songs)
	assert.Equal(t, 2, len(songs))
	assert.Equal(t, "Test Song 1", songs[0].Title)
	assert.Equal(t, "Test Song 2", songs[1].Title)
}

func TestSongRepo_GetSongById(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "band_id"}).
		AddRow(1, "Test Song 1", 9)

	mock.ExpectQuery("^SELECT (.+) FROM `songs` WHERE `songs`.`id` = ?").WithArgs(1).WillReturnRows(rows)
	mock.ExpectCommit()

	repo := &SongRepo{DB: db}
	song, err := repo.GetSongById(1)

	var bandIdAssertion = new(int)
	*bandIdAssertion = 9

	assert.NoError(t, err)
	assert.NotNil(t, song)
	assert.Equal(t, "Test Song 1", song.Title)
	assert.Equal(t, bandIdAssertion, song.BandID)
}

// TODO: implement GetSongsByAlbumId

func TestSongRepo_GetSongsByBandId(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testBandId := 9

	rows := sqlmock.NewRows([]string{"id", "title", "band_id"}).
		AddRow(1, "Test Song 1", testBandId).
		AddRow(2, "Test Song 2", testBandId)

	mock.ExpectQuery("^SELECT (.+) FROM `songs` WHERE band_id = ?").WithArgs(testBandId).WillReturnRows(rows)
	mock.ExpectCommit()

	repo := &SongRepo{DB: db}
	songs, err := repo.GetSongsByBandId(testBandId)

	assert.NoError(t, err)
	assert.NotNil(t, songs)
	assert.Equal(t, 2, len(songs))
	assert.Equal(t, "Test Song 1", songs[0].Title)
	assert.Equal(t, "Test Song 2", songs[1].Title)
}

// Update
// func TestSongRepo_UpdateSong(t *testing.T) {
// 	db, mock, err := mockDb()
// 	if err != nil {
// 		t.Errorf("Error creating mock database: %v", err)
// 	}

// 	// Test data
// 	testBandId := 9
// 	sqlmock.NewRows([]string{"id", "title", "band_id"}).
// 		AddRow(1, "Test Song 1", testBandId).
// 		AddRow(2, "Test Song 2", testBandId)

// 	repo := &SongRepo{DB: db}

// 	input := &model.SongInput{
// 		Title: "New Song!",
// 	}

// 	// Test the update
// 	mock.ExpectBegin()
// 	mock.ExpectExec("^UPDATE `songs`").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	song, err := repo.UpdateSong(input)

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.NotNil(t, song)
// 	assert.Equal(t, "New Song!", song.Title)
// 	assert.Equal(t, testBandId, *song.BandID)
// }

// Delete
// TODO: implement DeleteSong
