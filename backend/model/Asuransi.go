package model

import (
	"gorm.io/gorm"
)

type Asuransi struct {
	gorm.Model
	Judul string `json:"judul" gorm:"type:varchar(256); not null; unique"`
	Premi int32 `json:"premi" gorm:"type:integer; not null"`
	UangPertanggungan int64 `json:"uangpertanggungan" gorm:"type:integer; not null"`
	Deskripsi string `json:"deskripsi" gorm:"type:text; not null"`
	Syarat string `json:"syarat" gorm:"type:text; not null"`
	Foto string `json:"foto" gorm:"type:varchar(256); not null"`
}