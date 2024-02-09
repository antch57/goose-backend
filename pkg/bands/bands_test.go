package bands

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/antch57/jam-statz/graph/model"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockDb() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating mock database: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("error opening Gorm database: %v", err)
	}

	return db, mock, nil
}

// Bands tests
// CREATE
func TestBandRepo_CreateBand(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	description := new(string)
	*description = "Test band description"

	input := &model.BandInput{
		Name:        "Test Band",
		Genre:       "Rock",
		Year:        2021,
		Description: description,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `bands` \\(`name`,`genre`,`year`,`description`\\) VALUES \\(\\?,\\?,\\?,\\?\\)$").WithArgs(input.Name, input.Genre, input.Year, input.Description).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &BandRepo{DB: db}
	res, err := repo.CreateBand(input)
	if err != nil {
		t.Errorf("Error creating band: %v", err)
	}

	assert.Equal(t, input.Name, res.Name)
	assert.Equal(t, input.Genre, res.Genre)
	assert.Equal(t, input.Year, res.Year)
	assert.Equal(t, *input.Description, *res.Description)
}

// READ
func TestBandRepo_GetBands(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "genre", "year", "description"}).
		AddRow(1, "Test Band", "Rock", 2021, "Test band description").
		AddRow(2, "Test Band 2", "Pop", 2020, "Test band 2 description")

	mock.ExpectQuery("^SELECT \\* FROM `bands`$").WillReturnRows(rows)

	repo := &BandRepo{DB: db}
	res, err := repo.GetBands()
	if err != nil {
		t.Errorf("Error getting bands: %v", err)
	}

	assert.Equal(t, 2, len(res))
	assert.Equal(t, "Test Band", res[0].Name)
	assert.Equal(t, "Rock", res[0].Genre)
	assert.Equal(t, 2021, res[0].Year)
	assert.Equal(t, "Test band description", *res[0].Description)
	assert.Equal(t, "Test Band 2", res[1].Name)
	assert.Equal(t, "Pop", res[1].Genre)
	assert.Equal(t, 2020, res[1].Year)
	assert.Equal(t, "Test band 2 description", *res[1].Description)
}

func TestBandRepo_GetBandById(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "genre", "year", "description"}).
		AddRow(1, "Test Band", "Rock", 2021, "Test band description")

	mock.ExpectQuery("^SELECT \\* FROM `bands` WHERE `bands`.`id` = \\? ORDER BY `bands`.`id` LIMIT 1$").WithArgs(1).WillReturnRows(rows)

	repo := &BandRepo{DB: db}
	res, err := repo.GetBandById(1)
	if err != nil {
		t.Errorf("Error getting band by ID: %v", err)
	}

	assert.Equal(t, 1, res.ID)
	assert.Equal(t, "Test Band", res.Name)
	assert.Equal(t, "Rock", res.Genre)
	assert.Equal(t, 2021, res.Year)
	assert.Equal(t, "Test band description", *res.Description)
}

// UPDATE
func TestBandRepo_UpdateBand(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	description := new(string)
	*description = "Test band description"

	input := &model.BandInput{
		Name:        "Test Band",
		Genre:       "Rock",
		Year:        2021,
		Description: description,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE `bands` SET `id`=\\?,`name`=\\?,`genre`=\\?,`year`=\\?,`description`=\\? WHERE id = \\?$").WithArgs(1, input.Name, input.Genre, input.Year, *input.Description, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &BandRepo{DB: db}
	band, err := repo.UpdateBand(1, input)
	if err != nil {
		t.Errorf("Error updating band: %v", err)
	}

	assert.Equal(t, 1, band.ID)
	assert.Equal(t, input.Name, band.Name)
	assert.Equal(t, input.Genre, band.Genre)
	assert.Equal(t, input.Year, band.Year)
	assert.Equal(t, *input.Description, *band.Description)
}

// DELETE
func TestBandRepo_DeleteBand(t *testing.T) {
	db, mock, err := mockDb()
	if err != nil {
		t.Errorf("Error creating mock database: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM `bands` WHERE `bands`.`id` = \\?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &BandRepo{DB: db}
	res, err := repo.DeleteBand(1)
	if err != nil {
		t.Errorf("Error deleting band: %v", err)
	}

	assert.Equal(t, true, res)
}
