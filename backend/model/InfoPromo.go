package model

import (
	"gorm.io/gorm"
)

type InfoPromo struct {
	gorm.Model
	Judul string `json:"judul" gorm:"type:varchar(256); not null"`
	Kategori string `json:"kategori" gorm:"type:varchar(30); not null"`
	Startdate string `json:"startdate" gorm:"type:date; not null"`
	Enddate string `json:"enddate" gorm:"type:date; not null"`
	Kodepromo string  `json:"kodepromo" gorm:"type:varchar(10); not null; unique"`
	Foto string `json:"foto" gorm:"type:varchar(200); not null"`
	Deskripsi string `json:"deskripsi" gorm:"type:text; not null"`
	Syarat string `json:"syarat" gorm:"type:text; not null"`
}