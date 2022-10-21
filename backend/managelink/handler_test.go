package managelink

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bagasalim/simas/auth"
	"github.com/bagasalim/simas/custom"
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
	Data    model.Link `json:"data"`
}

func initialRepoAuth(t *testing.T) *auth.Handler {
	db := newTestDB(t)
	repoUser := auth.NewRepository(db)

	repoService := auth.NewService(repoUser)
	repoHandler := auth.NewHandler(repoService)
	return repoHandler
}
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

func getToken(t *testing.T) string {
	os.Setenv("testing","y")
	// getOtp(t)
	handler := initialRepoAuth(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/send-otp", handler.SendOTP)
	payload := `{"username": "CS01"}`
	req, _ := http.NewRequest("POST", "/send-otp", strings.NewReader(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	

	r.POST("/login", handler.Login)
	payload = `{"username": "CS01", "password":"123456","code":"123456"}`
	req, _ = http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w = httptest.NewRecorder()
	data := resposeLogin{}
	r.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &data)
	return data.Token

}
func TestGetLinkRequest(t *testing.T) {
	// token := getToken(t)
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/get-link/:type", handler.GetLinkRequest)

	//link not found
	req, _ := http.NewRequest("GET", "/get-link/wa", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resMes := ResponseMessage{}
	assert.Equal(t, 500, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resMes))

	//link found
	req, _ = http.NewRequest("GET", "/get-link/WA", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
}
func TestGetLinkHandler(t *testing.T) {

	token := getToken(t)
	handler := initialRepo(t)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	manageLinkRoute := r.Group("")
	middleware := custom.MiddleWare{}
	manageLinkRoute.Use(middleware.Auth)
	manageLinkRoute.Use(middleware.IsCS)
	manageLinkRoute.GET("/getlink", handler.GetLink)

	//not foundLink
	req, _ := http.NewRequest("GET", "/getlink?linktype=wa", nil)
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("res test getlink",string(w.Body.Bytes()))
	responseMessage := ResponseMessage{}
	assert.Equal(t, 500, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	//link == ""
	req, _ = http.NewRequest("GET", "/getlink?linktype=", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseMessage = ResponseMessage{}
	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	//found
	req, _ = http.NewRequest("GET", "/getlink?linktype=WA", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	type DataResponse struct {
		LinkValue string `json:"linkvalue"`
	}

	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

}
func TestHandlerUpdateLink(t *testing.T) {
	token := getToken(t)
	handler := initialRepo(t)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	manageLinkRoute := r.Group("")
	middleware := custom.MiddleWare{}
	manageLinkRoute.Use(middleware.Auth)
	manageLinkRoute.Use(middleware.IsCS)
	manageLinkRoute.POST("/update-link", handler.UpdateLink)

	//test fail update because no linktype
	req, _ := http.NewRequest("POST", "/update-link", nil)
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	type responseErrorValidation struct {
		Error []string `json:"error"`
	}
	responseMessage := responseErrorValidation{}
	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	//test fail update because no data insert
	req, _ = http.NewRequest("POST", "/update-link?linktype=wa", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseMessage = responseErrorValidation{}
	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	// fail data because not found linktype
	payload := `{"linktype": "WA", "linkvalue":"test", "updatedby":"rema"}`
	req, _ = http.NewRequest("POST", "/update-link?linktype=wa", strings.NewReader(payload))
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)

	//success
	payload = `{"linktype": "WA", "linkvalue":"test", "updatedby":"rema"}`
	req, _ = http.NewRequest("POST", "/update-link?linktype=WA", strings.NewReader(payload))
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Code, string(w.Body.Bytes()[:]))
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

}
