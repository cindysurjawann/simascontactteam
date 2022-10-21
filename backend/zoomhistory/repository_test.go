package zoomhistory

import (
	"fmt"
	"testing"

	_ "errors"

	"github.com/bagasalim/simas/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
	err = db.AutoMigrate(&model.User{}, &model.Riwayat{})
	assert.NoError(t, err)

	Riwayat := model.Riwayat{
		Nama:       "cindu",
		Email:      "cindu@gmail.com",
		Kategori:   "Perbankan",
		Keterangan: "Isu Perbankan",
		Lokasi:     "Jakarta",
	}
	db.Create(&Riwayat)

	return db
}

func TestCreateUser(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	// repo.
	Riwayat := model.Riwayat{
		Nama:       "cayo",
		Email:      "cayo@gmail.com",
		Kategori:   "Kredit",
		Keterangan: "Gatau",
	}
	// task := "task 1"
	res, err := repo.AddUser(Riwayat)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	Riwayat = model.Riwayat{
		Nama: "cayooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo",
	}
	res, err = repo.AddUser(Riwayat)
	fmt.Println(err, res)
	assert.NotNil(t, err)

}

func TestGetRiwayat(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	res, err := repo.GetRiwayat()
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
