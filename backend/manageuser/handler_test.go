package manageuser

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func initialRepo(t *testing.T) *Handler {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return handler
}

func initialRepoFailed(t *testing.T) *Handler {
	db := newTestDB(t)
	var User model.User
	db.Migrator().DropTable(&User)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return handler
}

type responseSuccess struct{
	Message string        `json:"message"`
	Data    []model.User `json:"data"`
}

type responseSingle struct{
	Message string        `json:"message"`
	Data    model.User `json:"data"`
}

type responseMessage struct{
	Message string        `json:"message"`
}

type responseError struct{
	Message []string        `json:"message"`
}

const getlink = "/getUser"
const updatelink = "/updateuser"
const deletelink = "/deleteuser"

func TestGetUserHandler(t *testing.T){
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET(getlink, handler.GetUser)

	req, _ := http.NewRequest("GET", getlink, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSuccess{}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

	handler = initialRepoFailed(t)
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()
	r.GET(getlink, handler.GetUser)

	req, _ = http.NewRequest("GET", getlink, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resMes := responseMessage{}
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resMes))
}

func TestGetUserReqHandler(t *testing.T){
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET(getlink, handler.GetUserReq)

	req, _ := http.NewRequest("GET", getlink+"?username=cindu", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSingle{}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

	req, _ = http.NewRequest("GET", getlink+"?username=", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resMes := responseMessage{}
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resMes))

	req, _ = http.NewRequest("GET", getlink+"?username=test", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resMes = responseMessage{}
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resMes))
}

func TestUpdateUserHandler(t *testing.T){
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.PUT(updatelink, handler.UpdateUser)

	req, _ := http.NewRequest("PUT", updatelink+"?username=", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseError{}
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

	payload := `{"nama": "cayo", "email":"calvin@gmail.com", "kategori":"Kredit","keterangan":"gatau"}`
	req, _ = http.NewRequest("PUT", updatelink+"?username=cindu", strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res = responseError{}
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))	

	payload = ``
	req, _ = http.NewRequest("PUT", updatelink+"?username=cindu", strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res = responseError{}
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))	

	payload = `{"Email": "cayo", "Role":2, "Name":"1a"}`
	req, _ = http.NewRequest("PUT", updatelink+"?username=cin", strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resMes := responseSingle{}
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resMes))	

	req, _ = http.NewRequest("PUT", updatelink+"?username=cindu", strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resSucc := responseSingle{}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resSucc))	
}

func TestDeleteHandler(t *testing.T){
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.DELETE(deletelink+"/:id", handler.DeleteUser)

	req, _ := http.NewRequest("DELETE", deletelink+"/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resSucc := responseSingle{}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resSucc))	

	handler = initialRepoFailed(t)
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()
	r.DELETE(deletelink+"/:id", handler.DeleteUser)
	req, _ = http.NewRequest("DELETE", deletelink+"/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resSucc = responseSingle{}
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resSucc))	
}
