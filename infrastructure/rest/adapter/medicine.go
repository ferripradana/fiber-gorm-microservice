package adapter

import (
	service "fiber-gorm-microservice/application/service/medicine"
	repository "fiber-gorm-microservice/infrastructure/repository/medicine"
	"fiber-gorm-microservice/infrastructure/rest/controllers/medicine"
	"gorm.io/gorm"
)

func MedicineAdapter(db *gorm.DB) medicine.MedicineController {
	medicineRepository := repository.NewMedicineRepositoryImpl(db)
	medicineService := service.NewMedicineServiceImpl(medicineRepository)
	return medicine.NewMedicineControllerImpl(medicineService)
}
