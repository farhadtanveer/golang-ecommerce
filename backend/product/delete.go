package product

func (svc *service) Delete(id int) error {
	return svc.prdctRepo.Delete(id)
}