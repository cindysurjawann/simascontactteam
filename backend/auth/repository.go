package auth

import (
	"errors"
	"time"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindUser(username string) (model.User, error)
	AddUser(user model.User) (model.User, error)
	AddOTP(data model.UserOTP) (error)
	FindOTP(id uint) (model.UserOTP, error)
	UpdateOTPExpire(id uint) error
	AddLastLogin(username string, lastlogin time.Time) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindUser(username string) (model.User, error) {
	var User model.User
	loc, _ := time.LoadLocation("GMT")
	if err := r.db.Where("username = ?", username).First(&User).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("not found")
		}
		return model.User{}, err
	}
	User.LastLogin = User.LastLogin.In(loc)
	return User, nil
}
func (r *repository) AddUser(user model.User) (model.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}
func (r *repository) AddOTP(data model.UserOTP) error{
	res := r.db.Create(&data)
	if res.Error != nil {
		return res.Error
	}
	return nil
	// res := r.db.Create()
}
func (r *repository) FindOTP(id uint) (model.UserOTP, error){
	var dataOtp model.UserOTP
	if err := r.db.Where("user_id = ?", id).Last(&dataOtp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserOTP{}, nil
		}
		return model.UserOTP{}, err
	}
	return dataOtp, nil
}
func (r *repository) UpdateOTPExpire(id uint)( error){
	model := model.UserOTP{}
	if tx:= r.db.Model(&model).Where("id", id).Update("used", "1"); tx.Error != nil{
		return tx.Error
	}
	return nil
}
func (r *repository) AddLastLogin(username string, lastlogin time.Time) (model.User, error) {
	loc, _ := time.LoadLocation("UTC")
	var User model.User
	res := r.db.Model(&User).Where("username=?", username).Updates(model.User{
		LastLogin: lastlogin.In(loc),
	})

	if res.Error != nil {
		return model.User{}, res.Error
	}

	return User, nil
}
