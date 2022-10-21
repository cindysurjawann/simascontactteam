package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/bagasalim/simas/custom"
	"github.com/bagasalim/simas/model"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(data LoginRequest) (model.User, int, error)
	CreateAccount(data RegisterRequest) (model.User, int, error)
	SetOtp(Username string) (string,string, int, error)
	UpdateLastLogin(data LastLoginRequest) (model.User, int, error)
}

type service struct {
	repo AuthRepository
}

func NewService(repo AuthRepository) *service {
	return &service{repo}
}
func (s *service) Login(data LoginRequest) (model.User, int, error) {
	username := data.Username
	User, err := s.repo.FindUser(username)

	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "not found" {
			status = http.StatusUnauthorized
			err =  errors.New("username or password is wrong")
		}
		return model.User{}, status, err
	}
	UserOtp, err := s.repo.FindOTP(User.ID)
	loc, _ := time.LoadLocation("Asia/Jakarta")
		
	if err != nil || UserOtp.Code == "" || UserOtp.Code != data.Code{
		return model.User{}, http.StatusUnauthorized, errors.New("OTP is wrong")
	}
	if UserOtp.Expire.In(loc).Before(time.Now().In(loc)) || UserOtp.Used{
		return model.User{}, http.StatusUnauthorized, errors.New("OTP is expire")
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(data.Password))
	if err != nil {
		return model.User{}, http.StatusUnauthorized, errors.New("Password is wrong")
	}
	err = s.repo.UpdateOTPExpire(UserOtp.ID)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	User.Password = ""
	return User, http.StatusOK, nil
}
func (s *service) CreateAccount(data RegisterRequest) (model.User, int, error) {
	found, err := s.repo.FindUser(data.Username)
	if err == nil && found.Name != "" {
		return model.User{}, http.StatusBadRequest, errors.New("duplicate data")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	User := model.User{
		Username: data.Username,
		Password: string(passwordHash),
		Name:     data.Name,
		Email: data.Email,
		Role:     2,
	}
	res, err := s.repo.AddUser(User)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
func(s *service) SetOtp(Username string) (string, string, int, error){
	User, err := s.repo.FindUser(Username)

	if err != nil {
		if err.Error() == "not found" {
			return  "","",http.StatusInternalServerError, errors.New("Username not found")
		}
		return  "","", http.StatusInternalServerError, err
	}
	data, err := s.repo.FindOTP(User.ID)
	if err != nil {
		return  "","", http.StatusInternalServerError, err
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	if data.Code != "" && data.Expire.Add(-2 * time.Minute).In(loc).After(time.Now().In(loc)) && data.Used == false{
		return data.Code, User.Email, http.StatusOK, nil
	}
	loc, _ = time.LoadLocation("UTC")
	var code string 
	if os.Getenv("testing") != "y"{
		code = custom.RandStringBytes(6)
	}else{
		code = "123456"
	}
	userLog := model.UserOTP{
		UserID: User.ID,
		Code: code,
		Expire: time.Now().Add(5 * time.Minute).In(loc),
	}
	err = s.repo.AddOTP(userLog)
	if err != nil {
		return  "","", http.StatusInternalServerError, err
	}
	return userLog.Code, User.Email, http.StatusOK, nil
}

func (s *service) UpdateLastLogin(data LastLoginRequest) (model.User, int, error) {
	res, err := s.repo.AddLastLogin(data.Username, data.LastLogin)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
