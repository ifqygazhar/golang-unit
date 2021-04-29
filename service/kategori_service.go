package service

import (
	"errors"
	"golang-unit/entity"
	"golang-unit/repo"
)

type KategoriService struct {
	Repository repo.KategoriRepo
}

func (service KategoriService) Get(id string) (*entity.Kategori, error) {
	category := service.Repository.FindId(id)
	if category == nil {
		return nil, errors.New("Kategori not found")
	} else {
		return category, nil
	}
}
