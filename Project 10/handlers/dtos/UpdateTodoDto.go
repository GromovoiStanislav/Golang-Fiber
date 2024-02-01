package dtos

type UpdateTodoDto struct {
	Title       string `json:"title" validate:""`
	Description string `json:"description" validate:""`
	Progress    int    `json:"progress" validate:"min=0,max=100"`
}