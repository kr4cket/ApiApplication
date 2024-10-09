package models

type Project struct {
	Id          int    `json:"-"`
	Name        string `json:"name" binding: "required"`
	Description string `json:"description" binding: "required"`
	Status      string `json:"status" binding: "required"`
	CompletedAt string `json:"completed_at" binding: "required"`
}
