package medicine

import domainMedicine "fiber-gorm-microservice/domain/medicine"

type MedicineRepository interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
	Create(newMedicine *domainMedicine.Medicine) (*domainMedicine.Medicine, error)
}
