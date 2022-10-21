package infoPromo

import (
	"errors"
	"fmt"
	"time"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type PromoRepository interface {
	GetInfos() ([]model.InfoPromo, error)
	GetRecentInfos() ([]model.InfoPromo, error)
	AddInfo(Info model.InfoPromo) (model.InfoPromo, error)
}

type repository	struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetInfos() ([]model.InfoPromo, error){
	var infos []model.InfoPromo
	if err := r.db.Find(&infos).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.InfoPromo{}, errors.New("information not found")
		}
		return []model.InfoPromo{}, err
	}
	return infos, nil
}

func (r *repository) GetRecentInfos() ([]model.InfoPromo, error){
	var infos []model.InfoPromo
	now := time.Now().UTC().Format("2006-01-02")
	fmt.Println(now)
	if err := r.db.Where("enddate >= ?", string(now)).Find(&infos).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.InfoPromo{}, errors.New("information not found")
		}
		return []model.InfoPromo{}, err
	}
	return infos, nil
}

func (r *repository) AddInfo(Info model.InfoPromo) (model.InfoPromo, error){
	res := r.db.Create(&Info)
	if res.Error != nil {
		return model.InfoPromo{}, res.Error
	}
	return Info, nil
}