package albums

import (
	"fmt"
	"strconv"
	"time"

	"github.com/antch57/goose/graph/model"
	"github.com/antch57/goose/internal/db"
	"github.com/antch57/goose/internal/songs"
)

func CreateAlbum(bandID string, title string, releaseDate string) (*model.Album, error) {
	fmt.Println("Creating album...")
	// parsedReleaseDate, err := ConvertReleaseDateToTime(releaseDate)
	// if err != nil {
	// 	return nil, err
	// }

	// res, err := db.Exec("INSERT INTO Albums (title, release_date, band_id) VALUES (?, ?, ?)", title, parsedReleaseDate, bandID)
	// if err != nil {
	// 	return nil, err
	// }

	// albumID, err := res.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// }

	// album := &model.Album{
	// 	ID: strconv.Itoa(int(albumID)),

	// }
	panic("not implemented")
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

func GetAlbumsByBandId(bandId int) ([]*model.Album, error) {
	albums := []*model.Album{}

	rows, err := db.Query("SELECT id, title, release_date FROM Albums WHERE band_id = ?", bandId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album model.Album

		err := rows.Scan(&album.ID, &album.Title, &album.ReleaseDate)
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

		album.Songs = songs

		albums = append(albums, &album)
	}

	return albums, nil
}
