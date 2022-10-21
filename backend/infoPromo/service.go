package infoPromo

import (
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	GetInfos() ([]model.InfoPromo, int, error)
	GetRecentInfos() ([]model.InfoPromo, int, error)
	AddInfo(data InfoRequest) (model.InfoPromo, int, error)
}

type service struct {
	repo PromoRepository
}

func NewService(repo PromoRepository) *service {
	return &service{repo}
}

func (s *service) GetInfos() ([]model.InfoPromo, int, error){
	infos, err :=  s.repo.GetInfos()
	if err != nil{
		return nil, http.StatusInternalServerError, err
	}
	return infos, http.StatusOK, nil
}

func (s *service) GetRecentInfos() ([]model.InfoPromo, int, error){
	infos, err :=  s.repo.GetRecentInfos()
	if err != nil{
		return nil, http.StatusInternalServerError, err
	}
	return infos, http.StatusOK, nil
}

func (s *service) AddInfo(data InfoRequest) (model.InfoPromo, int, error) {
	Info := model.InfoPromo{
		Judul: data.Judul,
		Kategori: data.Kategori,
		Startdate: data.Startdate,
		Enddate: data.Enddate,
		Kodepromo: data.Kodepromo,
		Foto : data.Foto,
		Deskripsi: data.Deskripsi,
		Syarat: data.Syarat,
	}

	res, err := s.repo.AddInfo(Info)
	if err != nil {
		return model.InfoPromo{}, http.StatusBadRequest, err
	}
	return res, http.StatusOK, nil
}

