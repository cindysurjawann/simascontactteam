package zoomhistory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	post = "/createzoomhistory"
	get  = "/getzoomhistory"
)

func TestCreateZoom(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST(post, handler.CreateZoom)
	payload := `{"nama": "cayo", "email":"calvin@gmail.com", "kategori":"Kredit","keterangan":"gatau"}`
	req, err := http.NewRequest("POST", post, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	type responseMess struct {
		Message string `json:"message"`
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("res", w.Code, string(w.Body.Bytes()[:]))
	var success responseMess
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))
	json.Unmarshal(w.Body.Bytes(), &success)

	// error validation
	payload = ``
	req, err = http.NewRequest("POST", post, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	type responseErrorValidation struct {
		Error []string `json:"error"`
	}
	var errValid responseErrorValidation
	// fmt.Println("res", w.Code, string(w.Body.Bytes()[:]))
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errValid))

}

func TestGetRiwayatHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET(get, handler.GetRiwayat)
	req, err := http.NewRequest("GET", get, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	type responseMess struct {
		Message string `json:"message"`
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("res", w.Code, string(w.Body.Bytes()[:]))
	var success responseMess
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))
	json.Unmarshal(w.Body.Bytes(), &success)

}
