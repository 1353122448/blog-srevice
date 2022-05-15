package model

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"createdOn"`
	ModifiedOn uint32 `json:"ModifiedOn"`
	DeleteOn   uint32 `json:"deleteOn"`
	IsDel      uint8  `json:"is_del"`
}
