package albums

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/antch57/goose/graph/model"
	"github.com/antch57/goose/internal/db"
	"github.com/antch57/goose/internal/songs"
)

func CreateAlbum(bandID string, title string, releaseDateString string, songsList []*model.SongInput, tx *sql.Tx, shouldCommit bool) (*model.Album, error) {
	fmt.Println("Creating album...")
	var album model.Album
	releaseDate, err := ConvertReleaseDateToTime(releaseDateString)
	if err != nil {
		return &model.Album{}, err
	}

	res, err := tx.Exec("INSERT INTO Albums (title, release_date, band_id) VALUES (?, ?, ?)", title, releaseDate, bandID)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			tx.Rollback()
			return &model.Album{}, fmt.Errorf("album '%s' already exists for bandId '%s' ", title, bandID)
		}
		tx.Rollback()
		return &model.Album{}, err
	}

	albumID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return &model.Album{}, err
	}

	albumIDString := strconv.Itoa(int(albumID))
	var songArray []*model.Song

	for _, song := range songsList {
		commitSongTransaction := false

		res, err := songs.CreateSong(bandID, albumIDString, song.Title, song.Duration, tx, commitSongTransaction)
		if err != nil {
			tx.Rollback()
			return &model.Album{}, err
		}
		songArray = append(songArray, res)
		fmt.Println("songArray: ", songArray)

	}
	if shouldCommit {
		err = tx.Commit()
		if err != nil {
			return &model.Album{}, err
		}
	}

	album = model.Album{
		ID:          strconv.Itoa(int(albumID)),
		Title:       title,
		ReleaseDate: releaseDateString,
		Songs:       songArray,
		BandID:      bandID,
	}
	return &album, nil
}

func GetAlbumsByBandId(bandId int) ([]*model.Album, error) {
	albums := []*model.Album{}

	rows, err := db.Query("SELECT id, title, release_date, band_id FROM Albums WHERE band_id = ?", bandId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album model.Album

		err := rows.Scan(&album.ID, &album.Title, &album.ReleaseDate, &album.BandID)
		if err != nil {
			return nil, err
		}

		albumID, err := strconv.Atoi(album.ID)
		if err != nil {
			return nil, err
		}

		songs, err := songs.GetSongsByAlbumId(albumID)
		if err != nil {
			return nil, err
		}

		album = model.Album{
			ID:          album.ID,
			Title:       album.Title,
			ReleaseDate: album.ReleaseDate,
			Songs:       songs,
			BandID:      album.BandID,
		}

		albums = append(albums, &album)
	}

	return albums, nil
}

func DeleteAlbum(albumID string) (bool, error) {
	fmt.Println("Deleting album...")
	_, err := db.Exec("DELETE FROM Albums WHERE id = ?", albumID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetAlbum(id string) (*model.Album, error) {
	fmt.Println("Getting Album...")
	var album model.Album

	row := db.QueryRow("SELECT id, title, release_date, band_id FROM Albums WHERE id = ?", id)
	err := row.Scan(&album.ID, &album.Title, &album.ReleaseDate, &album.BandID)
	if err != nil {
		return nil, err
	}

	album = model.Album{
		ID:          album.ID,
		Title:       album.Title,
		ReleaseDate: album.ReleaseDate,
		BandID:      album.BandID,
	}
	return &album, nil
}

// Utility function to convert AlbumInput to Album
func ConvertAlbumInputsToAlbums(albumInputs []*model.AlbumInput) []*model.Album {
	albums := []*model.Album{}

	for _, albumInput := range albumInputs {
		album := &model.Album{
			Title:       albumInput.Title,
			ReleaseDate: albumInput.ReleaseDate,
			Songs:       songs.ConvertSongInputsToSongs(albumInput.Songs),
		}
		albums = append(albums, album)
	}

	return albums
}

// Utility function to convert ReleaseDate string to time.Time for albums
func ConvertReleaseDateToTime(dateString string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
