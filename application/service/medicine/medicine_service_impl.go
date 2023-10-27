package medicine

import (
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

func (m MedicineServiceImpl) GetAll(page int64, limit int64) (*PaginationResultMedicine, error) {
	all, err := m.MedicineRepository.GetAll(page, limit)
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
