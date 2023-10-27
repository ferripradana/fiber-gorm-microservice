package adapter

import (
	service "fiber-gorm-microservice/application/service/medicine"
	repository "fiber-gorm-microservice/infrastructure/repository/medicine"
	"fiber-gorm-microservice/infrastructure/rest/controllers"
	"gorm.io/gorm"
)

func MedicineAdapter(db *gorm.DB) controllers.MedicineController {
	medicineRepository := repository.NewMedicineRepositoryImpl(db)
	medicineService := service.NewMedicineServiceImpl(medicineRepository)
	return controllers.NewMedicineControllerImpl(medicineService)
}
