package medicine

import (
	domainMedicine "fiber-gorm-microservice/domain/medicine"
	medicineRepository "fiber-gorm-microservice/infrastructure/repository/medicine"
)

type MedicineServiceImpl struct {
	MedicineRepository medicineRepository.MedicineRepository
}

func NewMedicineServiceImpl(repository medicineRepository.MedicineRepository) MedicineService {
	return &MedicineServiceImpl{
		MedicineRepository: repository,
	}
}

func (service *MedicineServiceImpl) GetAll(page int64, limit int64) (*PaginationResultMedicine, error) {
	all, err := service.MedicineRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}
	return &PaginationResultMedicine{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

func (service *MedicineServiceImpl) Create(medicine *NewMedicine) (*domainMedicine.Medicine, error) {
	medicineDomain := medicine.toDomainMapper()
	return service.MedicineRepository.Create(medicineDomain)
}
