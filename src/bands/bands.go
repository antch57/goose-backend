package bands

import (
	"fmt"

	"github.com/antch57/goose/graph/model"
)

func CreateBand() (model.Band, error) {
	fmt.Println("Creating band...")
	fmt.Println("save to db later..")
	year := 2024

	band := model.Band{Name: "ant", Genre: "jamz", Year: &year, Albums: []*model.Album{}}
	return band, nil
}

func GetBands() ([]*model.Band, error) {
	fmt.Println("Getting band...")
	fmt.Println("get from db later..")
	year := 2024
	return []*model.Band{{Name: "ant", Genre: "jamz", Year: &year, Albums: []*model.Album{}}}, nil
}
