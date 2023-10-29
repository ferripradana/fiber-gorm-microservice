package medicine

import domainMedicine "fiber-gorm-microservice/domain/medicine"

type MedicineRepository interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
	Create(newMedicine *domainMedicine.Medicine) (*domainMedicine.Medicine, error)
	GetById(id int) (*domainMedicine.Medicine, error)
	Delete(id int) error
	Update(id int, medicineMap map[string]interface{}) (*domainMedicine.Medicine, error)
}
