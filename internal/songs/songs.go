package songs

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/antch57/goose/graph/model"
// 	"github.com/antch57/goose/internal/db"
// )

// func CreateSong(input *model.SongInput, tx *sql.Tx, shouldCommit bool) (*model.Song, error) {
// 	fmt.Println("Creating song...")

// 	// Validate the song input
// 	if input == nil || strings.TrimSpace(input.Title) == "" {
// 		return &model.Song{}, errors.New("songInput or songInput.Title is empty")
// 	}
// 	title := strings.TrimSpace(input.Title)

// 	res, err := tx.Exec("INSERT INTO Songs (title) VALUES (?)", title)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "Duplicate entry") {
// 			tx.Rollback()
// 			return &model.Song{}, fmt.Errorf("song '%s' already exists for bandID '%s' ", title)
// 		}
// 		tx.Rollback()
// 		return &model.Song{}, err
// 	}

// 	songID, err := res.LastInsertId()
// 	if err != nil {
// 		tx.Rollback()
// 		return &model.Song{}, err
// 	}

// 	if shouldCommit {
// 		err = tx.Commit()
// 		if err != nil {
// 			return &model.Song{}, err
// 		}
// 	}

// 	song = model.Song{
// 		ID:    strconv.Itoa(int(songID)),
// 		Title: title,
// 	}

// 	return &song, nil
// }

// func UpdateSong(id string, input *model.SongInput, tx *sql.Tx, shouldCommit bool) (*model.Song, error) {
// 	fmt.Println("Updating song...")
// 	if input == nil {
// 		return &model.Song{}, errors.New("songInput is nil")
// 	}

// 	title := &input.Title
// 	if title == nil {
// 		return &model.Song{}, errors.New("at least one field must be provided to update")
// 	}

// 	query := "UPDATE Songs SET title = ? WHERE id = ?"
// 	_, err := tx.Exec(query, title, id)
// 	if err != nil {
// 		tx.Rollback()
// 		return &model.Song{}, err
// 	}

// 	err = tx.QueryRow("SELECT id, title FROM Songs WHERE id = ?", id).Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.BandID)
// 	if err != nil {
// 		tx.Rollback()
// 		return &model.Song{}, err
// 	}

// 	if shouldCommit {
// 		err = tx.Commit()
// 		if err != nil {
// 			return &model.Song{}, err
// 		}
// 	}

// 	song = model.Song{
// 		ID:       song.ID,
// 		Title:    song.Title,
// 		Duration: song.Duration,
// 		AlbumID:  song.AlbumID,
// 		BandID:   song.BandID,
// 	}
// 	return &song, nil
// }

// func DeleteAlbumSong(songID string) (bool, error) {
// 	fmt.Println("Deleting song...")
// 	_, err := db.Exec("DELETE FROM Songs WHERE id = ?", songID)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// func GetAlbumSong(songId string) (*model.Song, error) {
// 	fmt.Println("Getting Song...")
// 	var song model.Song

// 	row := db.QueryRow("SELECT id, title, duration, band_id, album_id FROM Songs WHERE id = ?", songId)
// 	err := row.Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.BandID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	song = model.Song{
// 		ID:       song.ID,
// 		Title:    song.Title,
// 		Duration: song.Duration,
// 		AlbumID:  song.AlbumID,
// 		BandID:   song.BandID,
// 	}
// 	return &song, nil
// }

// func GetAlbumSongs(albumId string) ([]*model.Song, error) {
// 	songs := []*model.Song{}

// 	rows, err := db.Query("SELECT id, title, duration, album_id, band_id FROM Songs WHERE album_id = ?", albumId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var song model.Song

// 		err := rows.Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.BandID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		song = model.Song{
// 			ID:       song.ID,
// 			Title:    song.Title,
// 			Duration: song.Duration,
// 			AlbumID:  song.AlbumID,
// 			BandID:   song.BandID,
// 		}
// 		songs = append(songs, &song)
// 	}

// 	return songs, nil
// }

// func GetSongsByBandId(bandId string) ([]*model.Song, error) {
// 	var songs = []*model.Song{}
// 	fmt.Println("Getting songs by band id...")
// 	fmt.Println("bandId: ", bandId)

// 	rows, err := db.Db.Query("SELECT id, title, duration, album_id, band_id FROM Songs WHERE band_id = ?", bandId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		fmt.Println("Getting song...")
// 		var song model.Song

// 		err := rows.Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.BandID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		song = model.Song{
// 			ID:       song.ID,
// 			Title:    song.Title,
// 			Duration: song.Duration,
// 			AlbumID:  song.AlbumID,
// 			BandID:   song.BandID,
// 		}
// 		songs = append(songs, &song)
// 	}

// 	return songs, nil
// }

// // Utility function to convert SongInput to Song
// // func ConvertSongInputsToSongs(songInputs []*model.SongInput) []*model.Song {
// // 	songs := []*model.Song{}

// // 	for _, songInput := range songInputs {
// // 		song := &model.Song{
// // 			Title:    songInput.Title,
// // 			Duration: songInput.Duration,
// // 			AlbumID:  songInput.AlbumID,
// // 		}
// // 		songs = append(songs, song)
// // 	}

// // 	return songs
// // }
