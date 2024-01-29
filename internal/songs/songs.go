package songs

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/antch57/goose/graph/model"
	"github.com/antch57/goose/internal/db"
)

func CreateSong(bandID string, albumID string, title string, duration int, tx *sql.Tx, shouldCommit bool) (*model.Song, error) {
	fmt.Println("Creating song...")
	fmt.Println("bandID: ", bandID)
	fmt.Println("albumID: ", albumID)
	fmt.Println("title: ", title)
	fmt.Println("duration: ", duration)
	var song model.Song

	res, err := tx.Exec("INSERT INTO Songs (title, duration, album_id, band_id) VALUES (?, ?, ?, ?)", title, duration, albumID, bandID)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			tx.Rollback()
			return &model.Song{}, fmt.Errorf("song '%s' already exists for bandID '%s' ", title, bandID)
		}
		tx.Rollback()
		return &model.Song{}, err
	}

	songID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return &model.Song{}, err
	}

	if shouldCommit {
		err = tx.Commit()
		if err != nil {
			return &model.Song{}, err
		}
	}

	song = model.Song{
		ID:       strconv.Itoa(int(songID)),
		Title:    title,
		Duration: duration,
		AlbumID:  &albumID,
		BandID:   bandID,
	}

	return &song, nil
}

func DeleteSong(songID string) (bool, error) {
	fmt.Println("Deleting song...")
	_, err := db.Exec("DELETE FROM Songs WHERE id = ?", songID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetSong(id string) (*model.Song, error) {
	fmt.Println("Getting Song...")
	var song model.Song

	row := db.QueryRow("SELECT id, title, duration, band_id, album_id FROM Songs WHERE id = ?", id)
	err := row.Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.BandID)
	if err != nil {
		return nil, err
	}

	song = model.Song{
		ID:       song.ID,
		Title:    song.Title,
		Duration: song.Duration,
		AlbumID:  song.AlbumID,
		BandID:   song.BandID,
	}
	return &song, nil
}

func GetSongsByAlbumId(albumId int) ([]*model.Song, error) {
	songs := []*model.Song{}

	rows, err := db.Query("SELECT id, title, duration, album_id, band_id FROM Songs WHERE album_id = ?", albumId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Song

		err := rows.Scan(&song.ID, &song.Title, &song.Duration, &song.AlbumID, &song.BandID)
		if err != nil {
			return nil, err
		}

		song = model.Song{
			ID:       song.ID,
			Title:    song.Title,
			Duration: song.Duration,
			AlbumID:  song.AlbumID,
			BandID:   song.BandID,
		}
		songs = append(songs, &song)
	}

	return songs, nil
}

// Utility function to convert SongInput to Song
func ConvertSongInputsToSongs(songInputs []*model.SongInput) []*model.Song {
	songs := []*model.Song{}

	for _, songInput := range songInputs {
		song := &model.Song{
			Title:    songInput.Title,
			Duration: songInput.Duration,
			AlbumID:  songInput.AlbumID,
		}
		songs = append(songs, song)
	}

	return songs
}
