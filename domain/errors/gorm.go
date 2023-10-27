package errors

type GormErr struct {
	Number  int    `json:"number"`
	Message string `json:"message"`
}
