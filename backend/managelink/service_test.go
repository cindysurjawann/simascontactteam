package managelink

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/stretchr/testify/assert"
)

func TestGetLinkService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	//Get WA
	req := GetLinkRequest{
		LinkType: "WA",
	}

	link, status, err := service.GetLink(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, link)

	//Get Zoom
	req = GetLinkRequest{
		LinkType: "Zoom",
	}

	link, status, err = service.GetLink(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, link)

}

func TestUpdateLinkService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	//Link WA
	link := UpdateLinkRequest{
		LinkType:  "WA",
		LinkValue: "Ini Link WA Update",
		UpdatedBy: "System",
	}
	res, status, err := service.UpdateLink(link)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link WA Update")

	//Link Zoom
	link = UpdateLinkRequest{
		LinkType:  "Zoom",
		LinkValue: "Ini Link Zoom Update",
		UpdatedBy: "System",
	}
	res, status, err = service.UpdateLink(link)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link Zoom Update")

	//No Link
	link = UpdateLinkRequest{
		LinkType:  "No Link",
		LinkValue: "No Link",
		UpdatedBy: "System",
	}
	res, status, err = service.UpdateLink(link)
	assert.Equal(t, err.Error(), errors.New("link not found").Error())
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, res, model.Link{})

}
