package asuransi

import (
	"log"
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	GetAsuransi() ([]model.Asuransi, int, error)
	CreateAsuransi(data AsuransiRequest) (model.Asuransi, int, error)
}

type service struct {
	repo AsuransiRepository
}

func NewService(repo AsuransiRepository) *service {
	return &service{repo}
}

func (s *service) GetAsuransi() ([]model.Asuransi, int, error) {
	asuransi, err := s.repo.GetAsuransi()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return asuransi, http.StatusOK, nil
}

func (s *service) CreateAsuransi(data AsuransiRequest) (model.Asuransi, int, error) {
	Insurance := model.Asuransi{
		Judul: data.Judul,
		Premi: data.Premi,
		UangPertanggungan: data.UangPertanggungan,
		Deskripsi: data.Deskripsi,
		Syarat: data.Syarat,
		Foto: data.Foto,
	}
	
	res, err := s.repo.CreateAsuransi(Insurance)
	if err != nil {
		return model.Asuransi{}, http.StatusBadRequest, err
	}
	return res, http.StatusOK, nil
}