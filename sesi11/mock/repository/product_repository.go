package repository

import "golang_sesi11/mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
