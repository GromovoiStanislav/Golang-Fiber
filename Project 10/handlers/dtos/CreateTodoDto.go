package dtos

type CreateTodoDto struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Progress    int    `json:"progress" validate:"min=0,max=100"`
}