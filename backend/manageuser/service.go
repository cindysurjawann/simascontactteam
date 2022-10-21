package manageuser

import (
	"errors"
	"log"
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	GetUser() ([]model.User, int, error)
	GetUserReq(data GetUserRequest) (model.User, int, error)
	UpdateUser(data UpdateUserRequest, username string) (model.User, int, error)
	DeleteUser(id string) (model.User, int, error)
}


type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *service {
	return &service{repo}
}

func (s *service) GetUser() ([]model.User, int, error) {
	user, err := s.repo.GetUser()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func (s *service) GetUserReq(data GetUserRequest) (model.User, int, error) {

	user, err := s.repo.GetUserReq(data.Username)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func (s *service) UpdateUser(data UpdateUserRequest, username string) (model.User, int, error) {

	found, err := s.repo.GetUserReq(username)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	if err == nil && found.Username == "" {
		return model.User{}, http.StatusBadRequest, errors.New("wrong username")
	}

	user := model.User{
		Email: data.Email,
		Role: data.Role,
		Name: data.Name,
	}

	res, err := s.repo.UpdateUser(user, username)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}

func (s *service) DeleteUser(id string) (model.User, int, error) {
	user, err := s.repo.DeleteUser(id)

	if err != nil {
		log.Println("Internal server error : ", err)
		return model.User{}, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}


