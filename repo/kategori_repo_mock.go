package repo

import (
	"golang-unit/entity"

	"github.com/stretchr/testify/mock"
)

type KategoriRepoMock struct {
	Mock mock.Mock
}

func (repo *KategoriRepoMock) FindId(id string) *entity.Kategori {
	argumen := repo.Mock.Called(id)
	if argumen.Get(0) == nil {
		return nil
	} else {
		kategori := argumen.Get(0).(entity.Kategori)
		return &kategori
	}
}
