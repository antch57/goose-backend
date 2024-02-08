package albums

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

// Albums tests
func TestAlbumRepo_CreateAlbum(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	layout := "2006-01-02"
	releaseDate, err := time.Parse(layout, "2021-01-01")
	if err != nil {
		t.Errorf("Error parsing release date: %v", err)
	}

	input := &model.AlbumInput{
		Title:       "Test Album",
		BandID:      1,
		ReleaseDate: releaseDate,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `albums` \\(`title`,`band_id`,`release_date`\\) VALUES \\(\\?,\\?,\\?\\)$").WithArgs(input.Title, input.BandID, input.ReleaseDate).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &AlbumRepo{DB: db}

	album, err := repo.CreateAlbum(input)

	assert.NoError(t, err)
	assert.Equal(t, album.Title, input.Title)
	assert.Equal(t, album.BandID, input.BandID)
	assert.Equal(t, album.ReleaseDate, input.ReleaseDate)
}

func TestAlbumRepo_GetAlbums(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	layout := "2006-01-02"
	testReleaseDate, err := time.Parse(layout, "2021-01-02")
	if err != nil {
		t.Errorf("Error parsing release date: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "release_date", "band_id"}).
		AddRow(1, "Test Album 1", testReleaseDate, 1).
		AddRow(2, "Test Album 2", testReleaseDate, 1)

	mock.ExpectQuery("^SELECT \\* FROM `albums`$").WillReturnRows(rows)

	repo := &AlbumRepo{DB: db}

	albums, err := repo.GetAlbums()

	assert.NoError(t, err)
	assert.Equal(t, "Test Album 1", albums[0].Title)
	assert.Equal(t, 1, albums[0].BandID)
	assert.Equal(t, testReleaseDate, albums[0].ReleaseDate)
	assert.Equal(t, "Test Album 2", albums[1].Title)
	assert.Equal(t, 1, albums[1].BandID)
	assert.Equal(t, testReleaseDate, albums[1].ReleaseDate)
}

func TestAlbumRepo_GetAlbumByID(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testRowID := 1
	layout := "2006-01-02"
	testReleaseDate, err := time.Parse(layout, "2021-01-02")
	if err != nil {
		t.Errorf("Error parsing release date: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "release_date", "band_id"}).
		AddRow(1, "Test Album 1", testReleaseDate, 1)

	mock.ExpectQuery("^SELECT \\* FROM `albums` WHERE `albums`.`id` = \\?").WithArgs(testRowID).WillReturnRows(rows)

	repo := &AlbumRepo{DB: db}

	album, err := repo.GetAlbumByID(testRowID)

	assert.NoError(t, err)
	assert.Equal(t, "Test Album 1", album.Title)
	assert.Equal(t, testRowID, album.BandID)
	assert.Equal(t, testReleaseDate, album.ReleaseDate)
}

func TestAlbumRepo_GetAlbumsByBandId(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testBandID := 1
	layout := "2006-01-02"
	testReleaseDate, err := time.Parse(layout, "2021-01-02")
	if err != nil {
		t.Errorf("Error parsing release date: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "release_date", "band_id"}).
		AddRow(1, "Test Album 1", testReleaseDate, testBandID).
		AddRow(3, "Test Album 3", testReleaseDate, testBandID)

	mock.ExpectQuery("^SELECT \\* FROM `albums` WHERE `albums`.`band_id` = \\?$").WithArgs(testBandID).WillReturnRows(rows)

	repo := &AlbumRepo{DB: db}

	albums, err := repo.GetAlbumsByBandId(testBandID)

	assert.NoError(t, err)
	assert.Equal(t, "Test Album 1", albums[0].Title)
	assert.Equal(t, testBandID, albums[0].BandID)
	assert.Equal(t, testReleaseDate, albums[0].ReleaseDate)
	assert.Equal(t, "Test Album 3", albums[1].Title)
	assert.Equal(t, testBandID, albums[1].BandID)
	assert.Equal(t, testReleaseDate, albums[1].ReleaseDate)
}

// AlbumSong tests
func TestAlbumRepo_CreateAlbumSong(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	input := &model.AlbumSongInput{
		AlbumID:     1,
		SongID:      1,
		Duration:    180,
		IsCover:     false,
		TrackNumber: 1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `album_songs` \\(`song_id`,`album_id`,`duration`,`track_number`,`is_cover`\\) VALUES \\(\\?\\,\\?\\,\\?\\,\\?\\,\\?\\)$").
		WithArgs(input.SongID, input.AlbumID, input.Duration, input.TrackNumber, input.IsCover).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	repo := &AlbumRepo{DB: db}

	albumSong, err := repo.CreateAlbumSong(input)

	assert.NoError(t, err)
	assert.Equal(t, input.AlbumID, albumSong.AlbumID)
	assert.Equal(t, input.SongID, input.SongID)
	assert.Equal(t, input.Duration, albumSong.Duration)
	assert.Equal(t, input.IsCover, *albumSong.IsCover)
	assert.Equal(t, input.TrackNumber, albumSong.TrackNumber)
}

func TestAlbumRepo_GetAlbumSongByID(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testRowID := 1
	rows := sqlmock.NewRows([]string{"id", "song_id", "album_id", "duration", "track_number", "is_cover"}).
		AddRow(1, 1, 1, 180, 1, false)

	mock.ExpectQuery("^SELECT \\* FROM `album_songs` WHERE `album_songs`.`id` = \\?").WithArgs(testRowID).WillReturnRows(rows)

	repo := &AlbumRepo{DB: db}

	albumSong, err := repo.GetAlbumSongByID(testRowID)

	assert.NoError(t, err)
	assert.Equal(t, 1, albumSong.AlbumID)
	assert.Equal(t, 1, albumSong.SongID)
	assert.Equal(t, 180, albumSong.Duration)
	assert.Equal(t, 1, albumSong.TrackNumber)
	assert.Equal(t, false, *albumSong.IsCover)
}

func TestAlbumRepo_GetAlbumSongs(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	testAlbumID := 1
	rows := sqlmock.NewRows([]string{"id", "song_id", "album_id", "duration", "track_number", "is_cover"}).
		AddRow(1, 1, testAlbumID, 180, 1, false).
		AddRow(2, 2, testAlbumID, 240, 2, false)

	mock.ExpectQuery("^SELECT \\* FROM `album_songs`$").WillReturnRows(rows)

	repo := &AlbumRepo{DB: db}

	albumSongs, err := repo.GetAlbumSongs()

	assert.NoError(t, err)
	assert.Equal(t, testAlbumID, albumSongs[0].AlbumID)
	assert.Equal(t, 1, albumSongs[0].SongID)
	assert.Equal(t, 180, albumSongs[0].Duration)
	assert.Equal(t, 1, albumSongs[0].TrackNumber)
	assert.Equal(t, false, *albumSongs[0].IsCover)
	assert.Equal(t, testAlbumID, albumSongs[1].AlbumID)
	assert.Equal(t, 2, albumSongs[1].SongID)
	assert.Equal(t, 240, albumSongs[1].Duration)
	assert.Equal(t, 2, albumSongs[1].TrackNumber)
	assert.Equal(t, false, *albumSongs[1].IsCover)
}
