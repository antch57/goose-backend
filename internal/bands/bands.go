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

func CreateBand(bandName string, genre string, year int, albumsInput []*model.AlbumInput, description *string, tx *sql.Tx) (model.Band, error) {
	fmt.Println("Creating band...")

	var res sql.Result
	var err error
	if description != nil {
		res, err = tx.Exec("INSERT INTO Bands (name, genre, year, description) VALUES (?, ?, ?, ?)", bandName, genre, year, *description)
	} else {
		res, err = tx.Exec("INSERT INTO Bands (name, genre, year) VALUES (?, ?, ?)", bandName, genre, year)
	}
	if err != nil {
		// Handle unique constraint error
		if strings.Contains(err.Error(), "Duplicate entry") {
			tx.Rollback()
			return model.Band{}, fmt.Errorf("band with name %s and genre %s already exists", bandName, genre)
		}

		tx.Rollback()
		return model.Band{}, err
	}

	bandID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return model.Band{}, err
	}

	bandIDString := strconv.Itoa(int(bandID))

	var albumList = []*model.Album{}
	var songArray = []*model.Song{}
	for _, album := range albumsInput {
		commitAlbumTransaction := false
		res, err := albums.CreateAlbum(bandIDString, album.Title, album.ReleaseDate, album.Songs, tx, commitAlbumTransaction)
		if err != nil {
			tx.Rollback()
			return model.Band{}, err
		}

		albumList = append(albumList, res)
		songArray = append(songArray, res.Songs...)
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

	// TODO: add songs array in return value
	band := model.Band{
		ID:          bandIDString,
		Name:        bandName,
		Genre:       genre,
		Year:        year,
		Albums:      albumList,
		Description: bandDescription,
		Songs:       songArray,
	}

	return band, nil
}

func DeleteBand(bandID string) (bool, error) {
	fmt.Println("Deleting Band...")
	_, err := db.Exec("DELETE FROM Bands WHERE id = ?", bandID)
	if err != nil {
		return false, err
	}

	return true, nil
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

		band = model.Band{
			ID:          band.ID,
			Name:        band.Name,
			Genre:       band.Genre,
			Albums:      albums,
			Year:        band.Year,
			Description: band.Description,
		}

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

	band = model.Band{
		ID:          band.ID,
		Name:        band.Name,
		Genre:       band.Genre,
		Albums:      albums,
		Year:        band.Year,
		Description: band.Description,
	}

	return &band, nil
}
