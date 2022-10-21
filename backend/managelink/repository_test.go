package managelink

import (
	"errors"
	"testing"

	_ "errors"

	"github.com/bagasalim/simas/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
	err = db.AutoMigrate(&model.Link{}, &model.User{}, &model.UserOTP{})
	assert.NoError(t, err)

	link := []model.Link{
		{
			LinkType:  "WA",
			LinkValue: "Ini Link WA",
			UpdatedBy: "System",
		},
		{
			LinkType:  "Zoom",
			LinkValue: "Ini Link Zoom",
			UpdatedBy: "System",
		},
	}
	err = db.Create(&link).Error
	assert.NoError(t, err)

	dataUser := []model.User{
		{
			Username: "CS01",
			Password: "$2a$10$BQHCjmHmEsFGJXCGWm7et.2lvVPecg0ibhFd/tgOCCCncTu5ieiA.",
			Name:     "Customer Service",
			Role:     2,
		},
	}
	db.Create(&dataUser)

	return db
}

func TestGetLink(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	//get Link WA
	res, err := repo.GetLink("WA")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link WA")

	//get Link Zoom
	res, err = repo.GetLink("Zoom")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link Zoom")

	//No Link
	res, err = repo.GetLink("No Link")
	assert.Equal(t, err.Error(), errors.New("link not found").Error())
	assert.Equal(t, res, model.Link{})
}

func TestUpdateLink(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	//Link WA
	link := model.Link{
		LinkType:  "WA",
		LinkValue: "Ini Link WA Update",
		UpdatedBy: "System",
	}
	res, err := repo.UpdateLink(link)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link WA Update")

	//Link Zoom
	link = model.Link{
		LinkType:  "Zoom",
		LinkValue: "Ini Link Zoom Update",
		UpdatedBy: "System",
	}
	res, err = repo.UpdateLink(link)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link Zoom Update")

	//No Link
	link = model.Link{
		LinkType:  "No Link Type",
		LinkValue: "No Link Value",
		UpdatedBy: "System",
	}
	res, err = repo.UpdateLink(link)
	assert.Equal(t, err.Error(), errors.New("wrong link type").Error())
	assert.Equal(t, res, model.Link{})

}
