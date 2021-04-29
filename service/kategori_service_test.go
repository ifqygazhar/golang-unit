package service

import (
	"golang-unit/entity"
	"golang-unit/repo"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var kategoriRepo = &repo.KategoriRepoMock{Mock: mock.Mock{}}
var kategoriService = KategoriService{Repository: kategoriRepo}

func TestKategoriService_GetNotfound(t *testing.T) {

	//program mock
	kategoriRepo.Mock.On("FindId", "1").Return(nil)

	category, err := kategoriService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestKategoriService_GetSucces(t *testing.T) {
	category := entity.Kategori{
		Id:   "2",
		Name: "Laptop",
	}

	kategoriRepo.Mock.On("FindId", "2").Return(category)

	result, err := kategoriService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
