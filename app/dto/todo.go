package dto

//todo DTO is a model that used by client when POST from /todo url
type TodoDTO struct {
	ID          string `json:"id" form:"id" binding:"omitempty,uuid"`
	Title       string `json:"title" form:"title"  binding:"required"`
	Description string `json:"description" form:"description" `
}
