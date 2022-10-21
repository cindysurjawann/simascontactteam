package model

import "gorm.io/gorm"

type Riwayat struct {
	gorm.Model
	Nama       string `json:"nama" gorm:"type:varchar(30); not null; default:null"`
	Email      string `json:"email" gorm:"type:varchar(100); not null; default:null"`
	Kategori   string `json:"kategori" gorm:"type:varchar(50); not null; default:null"`
	Keterangan string `json:"keterangan" gorm:"type:varchar(255); not null; default:null"`
	Lokasi     string `json:"lokasi" gorm:"type:varchar(255);"`
}
