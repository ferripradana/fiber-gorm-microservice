package medicine

import (
	domainMedicine "fiber-gorm-microservice/domain/medicine"
)

func (n *NewMedicine) toDomainMapper() *domainMedicine.Medicine {
	return &domainMedicine.Medicine{
		Name:        n.Name,
		Description: n.Description,
		EANCode:     n.EANCode,
		Laboratory:  n.Laboratory,
	}
}
