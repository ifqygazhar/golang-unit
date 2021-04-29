package repo

import "golang-unit/entity"

type KategoriRepo interface {
	FindId(id string) *entity.Kategori
}
