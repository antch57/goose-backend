package songs

import (
	"fmt"

	"github.com/antch57/goose/graph/model"
	"github.com/antch57/goose/internal/db"
)

func CreateSong() {
	fmt.Println("Creating album...")
	panic("CreateSong not implemented")
}

func GetSongsByAlbumId(albumId int) ([]*model.Song, error) {
	songs := []*model.Song{}

	rows, err := db.Query("SELECT id, title, duration FROM Songs WHERE album_id = ?", albumId)
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

// Utility function to convert SongInput to Song
func ConvertSongInputsToSongs(songInputs []*model.SongInput) []*model.Song {
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
