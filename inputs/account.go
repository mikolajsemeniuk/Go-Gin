package inputs

type Account struct {
	Code  string `json:"code" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}
