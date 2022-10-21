package model

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	LinkValue string `json:"linkvalue" gorm:"type:varchar(256);"`
	LinkType  string `json:"linktype" gorm:"type:varchar(100); not null;unique"`
	UpdatedBy string `json:"updatedby" gorm:"type:varchar(30);"`
}
