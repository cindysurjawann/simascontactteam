package infoPromo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceGetPromos(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	info, status, err := service.GetInfos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, info)
}

func TestServiceGetRecentPromos(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	info, status, err := service.GetRecentInfos()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, info)
}

func TestServiceAddInfo(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	data := InfoRequest{
		Judul: "Gebyar Sinarmas",
			Kategori: "Promo Simobiplus",
			Startdate: "2022-10-20",
			Enddate: "2022-10-30",
			Kodepromo: "202313",
			Foto: "test123",	
			Deskripsi: "Gebyar sinarmas hadir untuk memeriahkan hari kemerdekaan indonesia, ayo join dan gebyarkan indonesia bersama sinarmas dan nikmati keunggulan diskon pembayaran melalui simobiplus",
			Syarat: "1. Satu rekening hanya bisa melakukan pembayaran satu kali; 2. Satu nomor hp hanya bisa melakukan pembayaran satu kali; 3. Nasabah dapat membuka rekening melalui simobiplus",
	}
	res, _, err := service.AddInfo(data)
	fmt.Println("Gebyar Sinarmas", data.Judul, res, err)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}