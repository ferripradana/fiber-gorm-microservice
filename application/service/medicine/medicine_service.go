package medicine

import domainMedicine "fiber-gorm-microservice/domain/medicine"

type MedicineService interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
	Create(medicine *NewMedicine) (*domainMedicine.Medicine, error)
	GetById(id int) (*domainMedicine.Medicine, error)
	Delete(id int) error
	Update(id int, medicineMap map[string]interface{}) (*domainMedicine.Medicine, error)
}
