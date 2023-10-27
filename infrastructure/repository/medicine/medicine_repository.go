package medicine

type MedicineRepository interface {
	GetAll(page int64, limit int64) (*PaginationResultMedicine, error)
}
