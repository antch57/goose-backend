package albums

import (
	"github.com/antch57/goose/graph/model"
	"gorm.io/gorm"
)

type AlbumRepo struct {
	DB *gorm.DB
}

func (a *AlbumRepo) CreateAlbum(input *model.AlbumInput) (*model.Album, error) {
	var albumInsert = &model.Album{
		Title:       input.Title,
		ReleaseDate: input.ReleaseDate,
		BandID:      input.BandID,
	}

	res := a.DB.Create(albumInsert)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumInsert, nil
}

func (a *AlbumRepo) CreateAlbumSong(input *model.AlbumSongInput) (*model.AlbumSong, error) {
	var albumSongInsert = &model.AlbumSong{
		SongID:      input.SongID,
		AlbumID:     input.AlbumID,
		BandID:      input.BandID,
		Duration:    input.Duration,
		TrackNumber: input.TrackNumber,
		IsCover:     &input.IsCover,
	}

	res := a.DB.Create(albumSongInsert)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumSongInsert, nil
}

func (a *AlbumRepo) GetAlbumByID(albumId int) (*model.Album, error) {
	var album *model.Album

	res := a.DB.First(&album, albumId)
	if res.Error != nil {
		return nil, res.Error
	}

	return album, nil
}

func (a *AlbumRepo) GetAlbumsByBandId(bandId int) ([]*model.Album, error) {
	var albumsList []*model.Album

	// TODO: fix this next
	// FIXME: only selects first row. should be all rows
	res := a.DB.First(&albumsList, bandId)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumsList, nil
}
