package asuransi

import (
	"errors"
	"log"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type AsuransiRepository interface {
	GetAsuransi() ([]model.Asuransi, error)
	CreateAsuransi(asuransi model.Asuransi) (model.Asuransi, error)
	UpdateAsuransi(asuransi model.Asuransi) (model.Asuransi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAsuransi(Asuransi model.Asuransi) (model.Asuransi, error){
	res := r.db.Create(&Asuransi)
	if res.Error != nil {
		return model.Asuransi{}, res.Error
	}
	return Asuransi, nil
}

func (r *repository) GetAsuransi() ([]model.Asuransi, error) {
	var asuransi []model.Asuransi
	res := r.db.Find(&asuransi)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return asuransi, nil
}

func (r *repository) UpdateAsuransi(asuransi model.Asuransi) (model.Asuransi, error) {
	_, err := r.GetAsuransi()
	if err != nil {
		return model.Asuransi{}, errors.New("wrong data insurance")
	}

	res := r.db.Where("id=?", asuransi.ID).Updates(model.Asuransi{
		Judul: asuransi.Judul,
	})
	if res.Error != nil {
		return model.Asuransi{}, res.Error
	}
	return asuransi, nil
}