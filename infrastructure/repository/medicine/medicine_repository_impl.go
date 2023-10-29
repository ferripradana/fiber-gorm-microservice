package medicine

import (
	"encoding/json"
	"fiber-gorm-microservice/domain/errors"
	domainMedicine "fiber-gorm-microservice/domain/medicine"
	"github.com/gofiber/fiber/v2"
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

func (m *MedicineRepositoryImpl) GetAll(page int64, limit int64) (*PaginationResultMedicine, error) {
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

func (m *MedicineRepositoryImpl) Create(newMedicine *domainMedicine.Medicine) (createdMedicine *domainMedicine.Medicine, err error) {
	medicine := fromDomainMapper(newMedicine)
	tx := m.DB.Create(medicine)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return createdMedicine, err
		}
		switch newError.Number {
		case 1062:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return createdMedicine, err
	}

	createdMedicine = medicine.toDomainMapper()
	return createdMedicine, nil
}

func (m *MedicineRepositoryImpl) GetById(id int) (*domainMedicine.Medicine, error) {
	var medicine Medicine
	err := m.DB.Where("id = ?", id).First(&medicine).Error
	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = errors.NewAppErrorWithType(errors.NotFound)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return &domainMedicine.Medicine{}, err
	}

	return medicine.toDomainMapper(), nil
}

func (m *MedicineRepositoryImpl) Delete(id int) (err error) {
	tx := m.DB.Delete(&domainMedicine.Medicine{}, id)
	if tx.Error != nil {
		err = errors.NewAppErrorImpl(tx.Error, errors.UnknownError, fiber.StatusInternalServerError)
		return err
	}

	if tx.RowsAffected == 0 {
		err = errors.NewAppErrorWithType(errors.NotFound)
	}

	return
}

func (m *MedicineRepositoryImpl) Update(id int, medicineMap map[string]interface{}) (*domainMedicine.Medicine, error) {
	var medicine Medicine
	medicine.ID = id

	err := m.DB.Where("id = ?", id).First(&medicine).Error
	if err != nil {
		err = errors.NewAppErrorWithType(errors.NotFound)
		return &domainMedicine.Medicine{}, err
	}

	err = m.DB.Model(&medicine).
		Select("name", "description", "ean_code", "laboratory").
		Updates(medicineMap).Error

	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError errors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			err = errors.NewAppErrorImpl(err, errors.UnknownError, fiber.StatusInternalServerError)
			return &domainMedicine.Medicine{}, err
		}

		switch newError.Number {
		case 1062:
			err = errors.NewAppErrorWithType(errors.ResourceAlreadyExists)
		default:
			err = errors.NewAppErrorWithType(errors.UnknownError)
		}
		return &domainMedicine.Medicine{}, err
	}

	return medicine.toDomainMapper(), nil
}
