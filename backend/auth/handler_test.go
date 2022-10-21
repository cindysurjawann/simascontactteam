package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bagasalim/simas/custom"
	"github.com/bagasalim/simas/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

const (
	login           = "/login"
	createuser      = "/create-user"
	updateLastLogin = "/updatelastlogin"
	SendOTP			= "/send-otp"
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
type responseErrorValidation struct {
	Error []string `json:"error"`
}
type responseError struct {
	Message string `json:"message"`
}
func initialRepoAuth(t *testing.T) *Handler {
	db := newTestDB(t)
	repoUser := NewRepository(db)

	repoService := NewService(repoUser)
	repoHandler := NewHandler(repoService)
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

	r.POST(SendOTP, handler.SendOTP)
	payload := `{"username": "cindu"}`
	req, _ := http.NewRequest("POST", SendOTP, strings.NewReader(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	r.POST("/login", handler.Login)
	payload = `{"username": "cindu", "password":"123456","code":"123456"}`
	req, _ = http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w = httptest.NewRecorder()
	
	data := resposeLogin{}
	r.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &data)
	return data.Token

}

func TestLogin(t *testing.T) {
	type responseSuccess struct {
		Data  map[string]any `json:"data"`
		Token string         `json:"token"`
	}
	os.Setenv("testing","y")
	db := newTestDB(t)
	repo := NewRepository(db)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	User := model.User{
		Username: "remasertu",
		Password: string(passwordHash),
		Name:     "rema",
		Role:     2,
	}
	// task := "task 1"
	repo.AddUser(User)
	service := NewService(repo)
	handler := NewHandler(service)
	service.SetOtp("remasertu")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST(login, handler.Login)

	payload := `{"username": "remasertu", "password":"123456","code":"123456"}`
	req, err := http.NewRequest("POST", login, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var success responseSuccess
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))

	//expire token
	var errorMessage1 responseError = responseError{}
	payload = `{"username": "remasertu", "password":"123456","code":"123456"}`
	req, err = http.NewRequest("POST", login, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errorMessage1))
	

	//validation
	payload = ``
	req, err = http.NewRequest("POST", login, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var errValid responseErrorValidation
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errValid))

	
	//validation otp
	payload = `{"username": "remasertu", "password":"123456","code":"12345"}`
	req, _ = http.NewRequest("POST", login, strings.NewReader(payload))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var errorMessage responseError = responseError{}
	fmt.Println("test login", string(w.Body.Bytes()))
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errorMessage))

	//validation password
	service.SetOtp("remasertu")
	payload = `{"username": "remasertu", "password":"12345","code":"123456"}`
	req, _ = http.NewRequest("POST", login, strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	errorMessage = responseError{}
	// fmt.Println("test login", string(w.Body.Bytes()))
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errorMessage))

	//notfound
	payload = `{"username": "remasertu1", "password":"123456","code":"123456"}`
	req, err = http.NewRequest("POST", login, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	errorMessage = responseError{}
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errorMessage))

}
func TestCreateUser(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST(createuser, handler.CreateUser)

	payload := `{"username": "remasertu", "password":"123456", "name":"rema","email":"ali@gmail.com"}`
	req, err := http.NewRequest("POST", createuser, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	type responseMess struct {
		Message string `json:"message"`
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var success responseMess
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))

	// error validation
	payload = ``
	req, err = http.NewRequest("POST", createuser, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	
	var errValid responseErrorValidation
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errValid))

	//error duplicate
	payload = `{"username": "remasertu", "password":"123456", "name":"rema","email":"ali@gmail.com"}`
	req, err = http.NewRequest("POST", createuser, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))
}

func TestUpdateLastLoginHandler(t *testing.T) {
	token := getToken(t)
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	manageLinkRoute := r.Group("")
	middleware := custom.MiddleWare{}
	manageLinkRoute.Use(middleware.Auth)
	manageLinkRoute.POST(updateLastLogin, handler.UpdateLastLogin)
	type responseErrorValidation struct {
		Error []string `json:"error"`
	}

	//success
	payload := `{"username": "cindu", "lastlogin":"2022-10-18T10:12:07.000Z"}`
	req, _ := http.NewRequest("POST", updateLastLogin, strings.NewReader(payload))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

	//fail wrong input
	payload = `{"username": "cindu", "lastlogin":"2022-10-1810:12:07.000Z"}`
	req, _ = http.NewRequest("POST", updateLastLogin, strings.NewReader(payload))
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseMessage := responseErrorValidation{}
	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))
}
func TestSendOTP(t *testing.T){
	os.Setenv("testing","y")
	handler := initialRepo(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST(SendOTP, handler.SendOTP)
	//sukses
	payload := `{"username": "cindu"}`
	req, _ := http.NewRequest("POST", SendOTP, strings.NewReader(payload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	//force failed validation
	payload = ``
	req, _ = http.NewRequest("POST", SendOTP, strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	var errValid responseErrorValidation
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errValid))

	//force failed validation
	payload = `{"username": "c"}`
	req, _ = http.NewRequest("POST", SendOTP, strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("test send otp mes", string(w.Body.Bytes()))
	assert.Equal(t, 500, w.Code)
	var errMes ResponseMessage
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errMes))

	//fail send message
	os.Setenv("testing","n")
	payload = `{"username": "cindu"}`
	req, _ = http.NewRequest("POST", SendOTP, strings.NewReader(payload))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
	var errMes1 ResponseMessage
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errMes1))
}