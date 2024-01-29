package bands

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/antch57/goose/graph/model"
	"github.com/antch57/goose/internal/albums"
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

	albumList := albums.ConvertAlbumInputsToAlbums(albumsInput)

	for _, album := range albumList {
		// FIXME: create album in albums pkg?
		releaseDate, err := albums.ConvertReleaseDateToTime(album.ReleaseDate)
		if err != nil {
			tx.Rollback()
			return model.Band{}, err
		}

		res, err := tx.Exec("INSERT INTO Albums (title, release_date, band_id) VALUES (?, ?, ?)", album.Title, releaseDate, bandID)
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
			// FIXME: create song in songs pkg?
			_, err = tx.Exec("INSERT INTO Songs (title, duration, album_id, band_id) VALUES (?, ?, ?, ?)", song.Title, song.Duration, albumID, bandID)
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
		Albums:      albumList,
		Description: bandDescription,
	}

	return band, nil
}

// Grab all bands from the database
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

		albums, err := albums.GetAlbumsByBandId(int(bandID))
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

// Grab a single band from the database based off of ID
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

	albums, err := albums.GetAlbumsByBandId(int(bandID))
	if err != nil {
		return nil, err
	}

	band.Albums = albums

	return &band, nil
}
