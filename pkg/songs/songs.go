package songs

import (
	"fmt"

	"github.com/antch57/goose/graph/model"
	"gorm.io/gorm"
)

type SongRepo struct {
	DB *gorm.DB
}

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

// func (s *SongRepo) GetSongsByAlbumId(albumId int)

func (s *SongRepo) GetSongsByBandId(bandId int) ([]*model.Song, error) {
	var songsList []*model.Song

	res := s.DB.Where("band_id = ?", bandId).Find(&songsList)
	if res.Error != nil {
		return nil, res.Error
	}

	return songsList, nil
}

// func (s *SongRepo) UpdateSong(input *model.SongInput) (*model.Song, error) {
// 	song := &model.Song{
// 		Title:  input.Title,
// 		BandID: input.BandID,
// 	}

// 	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")

// 	res := s.DB.Model(&model.Song{}).Where("id = ?", input.ID).Updates(song)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}

// 	return song, nil
// }
