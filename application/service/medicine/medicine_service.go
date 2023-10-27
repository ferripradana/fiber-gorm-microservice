package medicine

type MedicineService interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
}
