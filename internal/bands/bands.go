package bands

// import (
// 	"database/sql"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/antch57/goose/graph/model"
// 	"github.com/antch57/goose/internal/albums"
// 	"github.com/antch57/goose/internal/db"
// 	"github.com/antch57/goose/internal/songs"
// )

// func CreateBand(bandName string, genre string, year int, albumsInput []*model.AlbumInput, description *string, tx *sql.Tx) (model.Band, error) {
// 	fmt.Println("Creating band...")

// 	var res sql.Result
// 	var err error
// 	if description != nil {
// 		res, err = tx.Exec("INSERT INTO Bands (name, genre, year, description) VALUES (?, ?, ?, ?)", bandName, genre, year, *description)
// 	} else {
// 		res, err = tx.Exec("INSERT INTO Bands (name, genre, year) VALUES (?, ?, ?)", bandName, genre, year)
// 	}
// 	if err != nil {
// 		// Handle unique constraint error
// 		if strings.Contains(err.Error(), "Duplicate entry") {
// 			tx.Rollback()
// 			return model.Band{}, fmt.Errorf("band with name %s and genre %s already exists", bandName, genre)
// 		}

// 		tx.Rollback()
// 		return model.Band{}, err
// 	}

// 	bandID, err := res.LastInsertId()
// 	if err != nil {
// 		tx.Rollback()
// 		return model.Band{}, err
// 	}

// 	bandIDString := strconv.Itoa(int(bandID))

// 	var albumList = []*model.Album{}
// 	var songArray = []*model.Song{}
// 	for _, album := range albumsInput {
// 		commitAlbumTransaction := false
// 		res, err := albums.CreateAlbum(bandIDString, album.Title, album.ReleaseDate, album.Songs, tx, commitAlbumTransaction)
// 		if err != nil {
// 			tx.Rollback()
// 			return model.Band{}, err
// 		}

// 		albumList = append(albumList, res)
// 		songArray = append(albumSongList, res.Songs...)
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return model.Band{}, err
// 	}

// 	var bandDescription *string
// 	if description != nil {
// 		// Create new string so that the pointer is not pointing to the same memory address
// 		bandDescription = new(string)
// 		*bandDescription = *description
// 	}

// 	band := model.Band{
// 		ID:          bandIDString,
// 		Name:        bandName,
// 		Genre:       genre,
// 		Year:        year,
// 		Albums:      albumList,
// 		Description: bandDescription,
// 		Songs:       songArray,
// 	}

// 	return band, nil
// }

// func DeleteBand(bandID string) (bool, error) {
// 	fmt.Println("Deleting Band...")
// 	_, err := db.Exec("DELETE FROM Bands WHERE id = ?", bandID)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// // Grab all bands from the database
// func GetBands() ([]*model.Band, error) {
// 	fmt.Println("Getting bands...")
// 	bands := []*model.Band{}

// 	rows, err := db.Query("SELECT id, name, genre, year, description FROM Bands")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var band model.Band

// 		err := rows.Scan(&band.ID, &band.Name, &band.Genre, &band.Year, &band.Description)
// 		if err != nil {
// 			return nil, err
// 		}

// 		albums, err := albums.GetAlbumsByBandId(band.ID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		bandSongs, err := songs.GetSongsByBandId(band.ID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		band = model.Band{
// 			ID:          band.ID,
// 			Name:        band.Name,
// 			Genre:       band.Genre,
// 			Albums:      albums,
// 			Year:        band.Year,
// 			Description: band.Description,
// 			Songs:       bandSongs,
// 		}

// 		bands = append(bands, &band)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return bands, nil
// }

// // Grab a single band from the database based off of ID
// func GetBand(bandId string) (*model.Band, error) {
// 	fmt.Println("Getting band...")
// 	var band model.Band

// 	row := db.QueryRow("SELECT id, name, genre, year, description FROM Bands WHERE id = ?", bandId)
// 	err := row.Scan(&band.ID, &band.Name, &band.Genre, &band.Year, &band.Description)
// 	if err != nil {
// 		return nil, err
// 	}

// 	albums, err := albums.GetAlbumsByBandId(bandId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	bandSongs, err := songs.GetSongsByBandId(bandId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	band = model.Band{
// 		ID:          band.ID,
// 		Name:        band.Name,
// 		Genre:       band.Genre,
// 		Albums:      albums,
// 		Year:        band.Year,
// 		Description: band.Description,
// 		Songs:       bandSongs,
// 	}

// 	return &band, nil
// }

// func UpdateBand(bandID string, name *string, genre *string, year *int, description *string, tx *sql.Tx, shouldCommit bool) (*model.Band, error) {
// 	fmt.Println("Updating Band...")
// 	var band model.Band
// 	if name == nil && genre == nil && year == nil && description == nil {
// 		return nil, fmt.Errorf("no fields to update")
// 	}

// 	// Start building the SQL query
// 	query := "UPDATE Bands SET "
// 	args := []interface{}{}

// 	if name != nil {
// 		query += "name = ?, "
// 		args = append(args, name)
// 	}
// 	if genre != nil {
// 		query += "genre = ?, "
// 		args = append(args, genre)
// 	}
// 	if year != nil {
// 		query += "year = ?, "
// 		args = append(args, year)
// 	}
// 	if description != nil {
// 		query += "description = ?, "
// 		args = append(args, description)
// 	}

// 	query = strings.TrimSuffix(query, ", ")

// 	query += " WHERE id = ?"
// 	args = append(args, bandID)

// 	_, err := tx.Exec(query, args...)
// 	if err != nil {
// 		tx.Rollback()
// 		return &model.Band{}, err
// 	}

// 	err = tx.QueryRow("SELECT id, name, genre, year, description FROM Bands WHERE id = ?", bandID).Scan(&band.ID, &band.Name, &band.Genre, &band.Year, &band.Description)
// 	if err != nil {
// 		tx.Rollback()
// 		return &model.Band{}, err
// 	}

// 	albumList, err := albums.GetAlbumsByBandId(band.ID)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	for _, album := range albumList {
// 		songList, err := songs.GetSongsByAlbumId(album.ID)
// 		if err != nil {
// 			tx.Rollback()
// 			return nil, err
// 		}

// 		// Add the songs to the album
// 		album.Songs = songList
// 		band.Albums = append(band.Albums, album)
// 	}

// 	allSongs, err := songs.GetSongsByBandId(bandID)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	if shouldCommit {
// 		err = tx.Commit()
// 		if err != nil {
// 			return &model.Band{}, err
// 		}
// 	}

// 	band = model.Band{
// 		ID:          band.ID,
// 		Name:        band.Name,
// 		Genre:       band.Genre,
// 		Year:        band.Year,
// 		Description: band.Description,
// 		Albums:      band.Albums,
// 		Songs:       allSongs,
// 	}
// 	return &band, nil
// }
