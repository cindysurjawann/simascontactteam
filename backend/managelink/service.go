package managelink

import (
	"errors"
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	GetLink(data GetLinkRequest) (model.Link, int, error)
	UpdateLink(data UpdateLinkRequest) (model.Link, int, error)
}

type service struct {
	repo LinkRepository
}

func NewService(repo LinkRepository) *service {
	return &service{repo}
}

func (s *service) GetLink(data GetLinkRequest) (model.Link, int, error) {

	link, err := s.repo.GetLink(data.LinkType)
	if err != nil {
		return model.Link{}, http.StatusInternalServerError, err
	}

	return link, http.StatusOK, nil
}

func (s *service) UpdateLink(data UpdateLinkRequest) (model.Link, int, error) {
	found, err := s.repo.GetLink(data.LinkType)
	if err != nil {
		return model.Link{}, http.StatusInternalServerError, err
	}
	if err == nil && found.LinkType == "" {
		return model.Link{}, http.StatusBadRequest, errors.New("wrong link type")
	}

	Link := model.Link{
		LinkType:  data.LinkType,
		LinkValue: data.LinkValue,
		UpdatedBy: data.UpdatedBy,
	}

	res, err := s.repo.UpdateLink(Link)
	if err != nil {
		return model.Link{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
