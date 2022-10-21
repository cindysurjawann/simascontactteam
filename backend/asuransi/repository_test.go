package asuransi

import (
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
	err = db.AutoMigrate(&model.Asuransi{})
	assert.NoError(t, err)

	asuransi := []model.Asuransi{
		{
			Judul:             "Asuransi Masa Depan",
			Premi:             200000,
			UangPertanggungan: 100000000,
			Deskripsi:         "Asuransi Masa Depan setiap tahun Anda hanya membayar premi sebesar Rp200.000 dan mendapat uang pertanggunan Rp100.000.000",
			Syarat:            "Minimal 17 tahun dan maksimal 62 tahun, WNI",
			Foto:              "test123",
		},
	}
	err = db.Create(&asuransi).Error
	assert.NoError(t, err)
	return db
}

func TestGetAsuransi(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	res, err := repo.GetAsuransi()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res[0].Judul, "Asuransi Masa Depan")

	db.Exec("delete from asuransis where ID = 1")

	res, _ = repo.GetAsuransi()
	assert.Equal(t, res, []model.Asuransi{})
}

func TestAddAsuransi(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	data := AsuransiRequest{
		Judul:             "Asuransi Surga",
		Premi:             200000,
		UangPertanggungan: 100000000,
		Deskripsi:         "Asuransi Surga setiap tahun Anda hanya membayar premi sebesar Rp200.000 dan mendapat uang pertanggunan Rp100.000.000",
		Syarat:            "Minimal 17 tahun dan maksimal 62 tahun, WNI",
		Foto:              "cobadong",
	}
	res, _, err := service.CreateAsuransi(data)
	//fmt.Println("Asuransi Surga", data.Judul, res, err)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
