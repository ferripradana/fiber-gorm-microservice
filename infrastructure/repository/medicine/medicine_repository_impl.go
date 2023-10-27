package medicine

import (
	"fiber-gorm-microservice/domain/errors"
	"gorm.io/gorm"
)

type MedicineRepositoryImpl struct {
	DB *gorm.DB
}

func NewMedicineRepositoryImpl(db *gorm.DB) MedicineRepository {
	return &MedicineRepositoryImpl{
		DB: db,
	}
}

func (m MedicineRepositoryImpl) GetAll(page int64, limit int64) (*PaginationResultMedicine, error) {
	var medicines []Medicine
	var total int64

	err := m.DB.Model(&Medicine{}).Count(&total).Error
	if err != nil {
		return &PaginationResultMedicine{}, err
	}

	offset := (page - 1) * limit
	err = m.DB.Limit(int(limit)).Offset(int(offset)).Find(&medicines).Error
	if err != nil {
		return &PaginationResultMedicine{}, err
	}

	if limit < 1 {
		return &PaginationResultMedicine{}, errors.NewAppErrorWithType(errors.ValidationError)
	}

	numPages := (total + limit - 1) / limit
	var nextCursor, prevCursor uint
	if page < numPages {
		nextCursor = uint(page + 1)
	}
	if page > 1 {
		prevCursor = uint(page - 1)
	}

	return &PaginationResultMedicine{
		Data:       sliceToDomainMapper(&medicines),
		Total:      total,
		Limit:      limit,
		Current:    page,
		NextCursor: nextCursor,
		PrevCursor: prevCursor,
		NumPages:   numPages,
	}, nil

}
