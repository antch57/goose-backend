package albums

import (
	"fmt"

	"github.com/antch57/goose/graph/model"
)

func CreateAlbum(bandID string, title string, releaseDate string) (*model.Album, error) {
	fmt.Println("Creating album...")
	panic("create albums not implemented")
}
