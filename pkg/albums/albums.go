package albums

import (
	"github.com/antch57/jam-statz/graph/model"
	"gorm.io/gorm"
)

type AlbumRepo struct {
	DB *gorm.DB
}

// Album CRUD Operations
// Create
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

// Read
func (a *AlbumRepo) GetAlbums() ([]*model.Album, error) {
	var albumsList []*model.Album
	res := a.DB.Find(&albumsList)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumsList, nil
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

	res := a.DB.Where("`albums`.`band_id` = ?", bandId).Find(&albumsList)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumsList, nil
}

// AlbumSong CRUD Operations
func (a *AlbumRepo) CreateAlbumSong(input *model.AlbumSongInput) (*model.AlbumSong, error) {
	var albumSongInsert = &model.AlbumSong{
		SongID:      input.SongID,
		AlbumID:     input.AlbumID,
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

func (a *AlbumRepo) GetAlbumSongByID(albumSongId int) (*model.AlbumSong, error) {
	var albumSong *model.AlbumSong

	res := a.DB.First(&albumSong, albumSongId)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumSong, nil
}

func (a *AlbumRepo) GetAlbumSongs() ([]*model.AlbumSong, error) {
	var albumSongsList []*model.AlbumSong
	res := a.DB.Find(&albumSongsList)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumSongsList, nil
}

func (a *AlbumRepo) GetAlbumSongsByAlbumId(albumId int) ([]*model.AlbumSong, error) {
	var albumSongsList []*model.AlbumSong

	res := a.DB.Where("`album_songs`.`album_id` = ?", albumId).Find(&albumSongsList)
	if res.Error != nil {
		return nil, res.Error
	}

	return albumSongsList, nil
}
