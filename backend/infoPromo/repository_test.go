package infoPromo

import (
	"fmt"
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
	err = db.AutoMigrate(&model.InfoPromo{})
	assert.NoError(t, err)

	info := []model.InfoPromo{
		{
			Judul: "Gebyar Sinarmas",
			Kategori: "Promo Simobiplus",
			Startdate: "2022-10-20",
			Enddate: "2022-10-30",
			Kodepromo: "202223",
			Foto: "test123",	
			Deskripsi: "Gebyar sinarmas hadir untuk memeriahkan hari kemerdekaan indonesia, ayo join dan gebyarkan indonesia bersama sinarmas dan nikmati keunggulan diskon pembayaran melalui simobiplus",
			Syarat: "1. Satu rekening hanya bisa melakukan pembayaran satu kali; 2. Satu nomor hp hanya bisa melakukan pembayaran satu kali; 3. Nasabah dapat membuka rekening melalui simobiplus",
		},
	}
	err = db.Create(&info).Error
	assert.NoError(t, err)
	return db
}

func TestGetPromos(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)

	res, err := repo.GetInfos()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res[0].Judul, "Gebyar Sinarmas")

	db.Exec("delete from info_promos where ID = 1")

	res, _ = repo.GetInfos()
	//assert.Equal(t, err.Error(), errors.New("information not found").Error())
	assert.Equal(t, res, []model.InfoPromo{})
}

func TestGetRecentPromos(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)

	res, err := repo.GetRecentInfos()
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res[0].Judul, "Gebyar Sinarmas")

	db.Exec("delete from info_promos where id = 1")

	res, _ = repo.GetRecentInfos()
	//assert.Equal(t, err.Error(), errors.New("information not found").Error())
	assert.Equal(t, res, []model.InfoPromo{})
}

func TestAddInfo(t *testing.T)()  {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	data := InfoRequest{
		Judul: "Gebyar Sinarmas",
		Kategori: "Promo Simobiplus",
		Startdate: "2022-10-20",
		Enddate: "2022-10-30",
		Kodepromo: "20225	",
		Foto: "test123",	
		Deskripsi: "Gebyar sinarmas hadir untuk memeriahkan hari kemerdekaan indonesia, ayo join dan gebyarkan indonesia bersama sinarmas dan nikmati keunggulan diskon pembayaran melalui simobiplus",
		Syarat: "1. Satu rekening hanya bisa melakukan pembayaran satu kali; 2. Satu nomor hp hanya bisa melakukan pembayaran satu kali; 3. Nasabah dapat membuka rekening melalui simobiplus",
	}
	res, _, err := service.AddInfo(data)
	fmt.Println("Gebyar Sinarmas", data.Judul, res, err)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}