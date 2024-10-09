package models

type Task struct {
	Id          int    `json:"-"`
	Name        string `json:"name" binding: "required"`
	Description string `json:"description" binding: "required"`
	Priority    string `json:"priority" binding: "required"`
	Status      string `json:"status" binding: "required"`
	CreatedAt   string `json:"created_at" binding: "required"`
	CompletedAt string `json:"completed_at" binding: "required"`
	ProjectId   int    `json:"project_id" binding: "not required"`
}
