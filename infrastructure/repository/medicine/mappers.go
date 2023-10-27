package medicine

import (
	domainMedicine "fiber-gorm-microservice/domain/medicine"
)

func (m *Medicine) toDomainMapper() *domainMedicine.Medicine {
	return &domainMedicine.Medicine{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		EANCode:     m.EANCode,
		Laboratory:  m.Laboratory,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func fromDomainMapper(m *domainMedicine.Medicine) *Medicine {
	return &Medicine{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		EANCode:     m.EANCode,
		Laboratory:  m.Laboratory,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func sliceToDomainMapper(medicines *[]Medicine) *[]domainMedicine.Medicine {
	medicinesDomain := make([]domainMedicine.Medicine, len(*medicines))
	for i, medicine := range *medicines {
		medicinesDomain[i] = *medicine.toDomainMapper()
	}
	return &medicinesDomain
}
