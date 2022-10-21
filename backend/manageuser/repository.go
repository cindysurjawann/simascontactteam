package manageuser

import (
	"errors"
	"log"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() ([]model.User, error)
	GetUserReq(username string) (model.User, error)
	UpdateUser(user model.User, username string) (model.User, error)
	DeleteUser(id string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]model.User, error) {
	var User []model.User
	res := r.db.Where("role=2").Find(&User)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return User, nil
}

func (r *repository) GetUserReq(username string) (model.User, error) {
	var User model.User
	if err := r.db.Where("username = ?", username).First(&User).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("username not found")
		}
		return model.User{}, err
	}
	return User, nil
}

func (r *repository) UpdateUser(user model.User, username string) (model.User, error) {
	_, err := r.GetUserReq(username)
	if err != nil {
		return model.User{}, errors.New("wrong username")
	}

	res := r.db.Where("username=?", username).Updates(model.User{
		Email: user.Email,
		Role: user.Role,
		Name: user.Name,
	})
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}

func (r *repository) DeleteUser(id string) (model.User, error) {
	user := model.User{}

	err := r.db.Model(&user).Where("id", id).Delete(&user).Error
	if err != nil {
		log.Println("Delete error : ", err)
		return model.User{}, err
	}
	return model.User{}, nil
}


