package asuransi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceGetAsuransi(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	info, status, err := service.GetAsuransi()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, info)
}

func TestServiceCreateAsuransi(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	data := AsuransiRequest{
		Judul:             "Asuransi Kesehatan",
		Premi:             200000,
		UangPertanggungan: 100000000,
		Deskripsi:         "Asuransi Kesehatan setiap tahun Anda hanya membayar premi sebesar Rp200.000 dan mendapat uang pertanggunan Rp100.000.000",
		Syarat:            "Minimal 17 tahun dan maksimal 62 tahun, WNI",
		Foto:              "test123",
	}
	res, _, err := service.CreateAsuransi(data)
	//fmt.Println("Asuransi Kesehatan", data.Judul, res, err)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}