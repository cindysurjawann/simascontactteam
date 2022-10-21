package zoomhistory

import (
	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type ZoomRepository interface {
	AddUser(Riwayat model.Riwayat) (model.Riwayat, error)
	GetRiwayat() ([]model.Riwayat, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) AddUser(Riwayat model.Riwayat) (model.Riwayat, error) {
	res := r.db.Create(&Riwayat)
	if res.Error != nil {
		return model.Riwayat{}, res.Error
	}

	return Riwayat, nil
}

func (r *repository) GetRiwayat() ([]model.Riwayat, error) {
	var Riwayat []model.Riwayat
	res := r.db.Limit(50).Order("created_at desc").Find(&Riwayat)
	if res.Error != nil {
		return nil, res.Error
	}

	return Riwayat, nil
}
