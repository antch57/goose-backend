package bands

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/antch57/goose/graph/model"
	"github.com/antch57/goose/internal/db"
)

func CreateBand(name string, genre string, year int, albumsInput []*model.AlbumInput, description *string) (model.Band, error) {
	fmt.Println("Creating band...")

	// Being db transaction
	tx, err := db.Transacntion()
	if err != nil {
		return model.Band{}, err
	}

	var res sql.Result
	if description != nil {
		res, err = tx.Exec("INSERT INTO Bands (name, genre, year, description) VALUES (?, ?, ?, ?)", name, genre, year, *description)
	} else {
		res, err = tx.Exec("INSERT INTO Bands (name, genre, year) VALUES (?, ?, ?)", name, genre, year)
	}
	if err != nil {
		// Handle unique constraint error
		if strings.Contains(err.Error(), "Duplicate entry") {
			tx.Rollback()
			return model.Band{}, fmt.Errorf("band with name %s and genre %s already exists", name, genre)
		}

		tx.Rollback()
		return model.Band{}, err
	}

	bandID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return model.Band{}, err
	}

	albums := convertAlbumInputsToAlbums(albumsInput)

	for _, album := range albums {
		releaseDate, err := convertReleaseDateToTime(album.ReleaseDate)
		if err != nil {
			tx.Rollback()
			return model.Band{}, err
		}

		res, err := tx.Exec("INSERT INTO Albums (title, releaseDate, bandId) VALUES (?, ?, ?)", album.Title, releaseDate, bandID)
		if err != nil {
			tx.Rollback()
			return model.Band{}, err
		}

		albumID, err := res.LastInsertId()
		if err != nil {
			tx.Rollback()
			return model.Band{}, err
		}

		for _, song := range album.Songs {
			_, err = tx.Exec("INSERT INTO Songs (title, duration, albumId) VALUES (?, ?, ?)", song.Title, song.Duration, albumID)
			if err != nil {
				tx.Rollback()
				return model.Band{}, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return model.Band{}, err
	}

	var bandDescription *string
	if description != nil {
		// Create new string so that the pointer is not pointing to the same memory address
		bandDescription = new(string)
		*bandDescription = *description
	}
	band := model.Band{
		ID:          strconv.Itoa(int(bandID)),
		Name:        name,
		Genre:       genre,
		Year:        year,
		Albums:      albums,
		Description: bandDescription,
	}

	return band, nil
}

func GetBands() ([]*model.Band, error) {
	fmt.Println("Getting bands...")
	bands := []*model.Band{}

	rows, err := db.Query("SELECT id, name, genre, year, description FROM Bands")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var band model.Band

		err := rows.Scan(&band.ID, &band.Name, &band.Genre, &band.Year, &band.Description)
		if err != nil {
			return nil, err
		}

		bandID, err := strconv.Atoi(band.ID)
		if err != nil {
			return nil, err
		}

		albums, err := getAlbumsByBandId(int(bandID))
		if err != nil {
			return nil, err
		}

		band.Albums = albums

		bands = append(bands, &band)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bands, nil
}

func GetBand(id string) (*model.Band, error) {
	fmt.Println("Getting band...")
	var band model.Band

	row := db.QueryRow("SELECT id, name, genre, year, description FROM Bands WHERE id = ?", id)
	err := row.Scan(&band.ID, &band.Name, &band.Genre, &band.Year, &band.Description)
	if err != nil {
		return nil, err
	}

	bandID, err := strconv.Atoi(band.ID)
	if err != nil {
		return nil, err
	}

	albums, err := getAlbumsByBandId(int(bandID))
	if err != nil {
		return nil, err
	}

	band.Albums = albums

	return &band, nil
}

// Utility function to convert AlbumInput to Album
func convertAlbumInputsToAlbums(albumInputs []*model.AlbumInput) []*model.Album {
	albums := []*model.Album{}

	for _, albumInput := range albumInputs {
		album := &model.Album{
			Title:       albumInput.Title,
			ReleaseDate: albumInput.ReleaseDate,
			Songs:       convertSongInputsToSongs(albumInput.Songs),
		}
		albums = append(albums, album)
	}

	return albums
}

// Utility function to convert SongInput to Song
func convertSongInputsToSongs(songInputs []*model.SongInput) []*model.Song {
	songs := []*model.Song{}

	for _, songInput := range songInputs {
		song := &model.Song{
			Title:    songInput.Title,
			Duration: songInput.Duration,
		}
		songs = append(songs, song)
	}

	return songs
}

func convertReleaseDateToTime(dateString string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func getAlbumsByBandId(bandId int) ([]*model.Album, error) {
	albums := []*model.Album{}

	rows, err := db.Query("SELECT id, title, releaseDate FROM Albums WHERE bandId = ?", bandId)
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

		songs, err := getSongsByAlbumId(albumID)
		if err != nil {
			return nil, err
		}

		album.Songs = songs

		albums = append(albums, &album)
	}

	return albums, nil
}

func getSongsByAlbumId(albumId int) ([]*model.Song, error) {
	songs := []*model.Song{}

	rows, err := db.Query("SELECT id, title, duration FROM Songs WHERE albumId = ?", albumId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Song

		err := rows.Scan(&song.ID, &song.Title, &song.Duration)
		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)
	}

	return songs, nil
}
