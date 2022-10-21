package infoPromo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bagasalim/simas/auth"
	"github.com/bagasalim/simas/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type responseLoginData struct {
	Name     string `json:"name"`
	Role     int8   `json:"role"`
	Username string `json:"username"`
}
type resposeLogin struct {
	Data  responseLoginData `json:"data"`
	Token string            `json:"token"`
}
type responseSuccess struct {
	Message string     `json:"message"`
	Data    []model.InfoPromo `json:"data"`
}

func initialRepoAuth(t *testing.T) *auth.Handler{
	db := newTestDB(t)
	repoUser := auth.NewRepository(db)

	repoService := auth.NewService(repoUser)
	repoHandler := auth.NewHandler(repoService)
	return repoHandler
}

func initialRepo(t *testing.T) *Handler{
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return handler
}

type ResponseMessage struct {
	Message string
}

func getToken(t *testing.T) string{
	handler := initialRepoAuth(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/login", handler.Login)
	payload := `{"username": "CS01", "password":"123456"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	data := resposeLogin{}
	r.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &data)
	return data.Token
}

func TestGetRecentHandler(t *testing.T){
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getrecentpromos", handler.GetRecentInfos)
	
	//link not found
	req, _ := http.NewRequest("GET", "/getrecentpromos/wa", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resMes:= ResponseMessage{}
	assert.Equal(t, 404, w.Code)
	assert.Error(t, json.Unmarshal(w.Body.Bytes(), &resMes))

	//link found
	req, _ = http.NewRequest("GET", "/getrecentpromos", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
}

func TestGetHandler(t *testing.T){
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getpromos", handler.GetInfos)
	
	//link not found
	req, _ := http.NewRequest("GET", "/getpromosa/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resMes:= ResponseMessage{}
	assert.Equal(t, 404, w.Code)
	assert.Error(t, json.Unmarshal(w.Body.Bytes(), &resMes))

	//link found
	req, _ = http.NewRequest("GET", "/getpromos", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
}

func TestAddInfoHandler(t *testing.T){
	post := "/postinfopromo"
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST(post, handler.AddInfo)

	payload := ``

	req, err := http.NewRequest("POST", post, strings.NewReader(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.NoError(t, err)
	assert.NotNil(t, req)
	resMes:= ResponseMessage{}
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resMes))

	payload = `{"Judul": "test", "Kategori":"test", "Startdate":"2022-10-10", "Enddate":"2022-10-10", "Kodepromo":"201212", "Foto":"test", "Deskripsi":"test", "Syarat":"test"}`

	req, err = http.NewRequest("POST", post, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res:= responseSuccess{}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
}  