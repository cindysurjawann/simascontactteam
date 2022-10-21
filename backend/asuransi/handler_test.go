package asuransi

import (
	"encoding/json"
	"fmt"
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

type ResponseMessage struct {
	Message string
}

type responseSuccess struct {
	Message string         `json:"message"`
	Data    model.Asuransi `json:"data"`
}

func TestGetAsuransiHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/getasuransi", handler.GetAsuransi)
	req, err := http.NewRequest("GET", "/getasuransi", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAddAsuransiHandler(t *testing.T) {
	post := "/postasuransi"
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST(post, handler.CreateAsuransi)

	payload := ``

	req, err := http.NewRequest("POST", post, strings.NewReader(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	resMes := ResponseMessage{}
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resMes))

	payload = `{"Judul": "Asuransi Test", "Premi":20000, "UangPertanggungan":100000, "Deskripsi":"coba deskripsi", "Syarat":"syarat coba", "Foto":"test234"}`

	req, err = http.NewRequest("POST", post, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	fmt.Println(string(w.Body.Bytes()))
	res := responseSuccess{}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
}