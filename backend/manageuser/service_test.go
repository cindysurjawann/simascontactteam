package manageuser

import (
	"net/http"
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/stretchr/testify/assert"
)

func TestGetUserService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	res, status, err := service.GetUser()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)

	var User model.User
	db.Migrator().DropTable(&User)
	repo = NewRepository(db)
	service = NewService(repo)
	res, status, err = service.GetUser()
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Nil(t, res)
}

func TestGetUserReqService(t *testing.T)  {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := GetUserRequest{
		Username: "cindu",
	}
	res, status, err := service.GetUserReq(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)

	req = GetUserRequest{
		Username: "cindo",
	}
	
	var User model.User
	res, status, err = service.GetUserReq(req)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, res, User)
}

func TestUpdateUserService(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := UpdateUserRequest{
		Email: "tes",
		Role: 2,
		Name: "tes",
	}

	res, status, err := service.UpdateUser(req, "cindu")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)

	res, status, err = service.UpdateUser(req, "cind")
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.NotNil(t, res)
	
}

func TestDeleteService(t *testing.T){
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	res, status, err := service.DeleteUser("1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)

	var User model.User
	db.Migrator().DropTable(&User)
	repo = NewRepository(db)
	service = NewService(repo)

	res, status, err = service.DeleteUser("1")
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.NotNil(t, res)
}