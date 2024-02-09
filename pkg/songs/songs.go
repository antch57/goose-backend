package songs

import (
	"fmt"

	"github.com/antch57/jam-statz/graph/model"
	"gorm.io/gorm"
)

type SongRepo struct {
	DB *gorm.DB
}

// Song CRUD Operations
// Create
func (s *SongRepo) CreateSong(input *model.SongInput) (*model.Song, error) {
	var songInsert = &model.Song{}
	songInsert.Title = input.Title

	if input.BandID != nil {
		songInsert.BandID = input.BandID
	}

	res := s.DB.Create(songInsert)
	if res.Error != nil {
		return nil, res.Error
	}

	return songInsert, nil
}

// Read
func (s *SongRepo) GetSongs() ([]*model.Song, error) {
	var songsList []*model.Song
	res := s.DB.Find(&songsList)
	if res.Error != nil {
		return nil, res.Error
	}

	fmt.Println("songsList ", songsList)
	return songsList, nil
}

func (s *SongRepo) GetSongById(songId int) (*model.Song, error) {
	var song = &model.Song{}

	res := s.DB.First(&song, songId)
	if res.Error != nil {
		return nil, res.Error
	}

	return song, nil
}

func (s *SongRepo) GetSongsByBandId(bandId int) ([]*model.Song, error) {
	var songsList []*model.Song

	res := s.DB.Where("band_id = ?", bandId).Find(&songsList)
	if res.Error != nil {
		return nil, res.Error
	}

	return songsList, nil
}

// Update
func (s *SongRepo) UpdateSong(songId int, input *model.SongInput) (*model.Song, error) {
	song := &model.Song{
		ID:     songId,
		Title:  input.Title,
		BandID: input.BandID,
	}

	res := s.DB.Model(&model.Song{}).Where("id = ?", songId).Updates(song)
	if res.Error != nil {
		return nil, res.Error
	}

	return song, nil
}

// Delete
func (s *SongRepo) DeleteSong(songId int) (bool, error) {
	res := s.DB.Delete(&model.Song{}, songId)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}
