package product

import "ecommerce/domain"

func (svc *service) List() ([]*domain.Product, error) {
	return svc.prdctRepo.List()
}