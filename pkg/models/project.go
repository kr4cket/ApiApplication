package models

import "github.com/google/uuid"

type Project struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Name        string    `json:"name" binding: "required"`
	Description string    `json:"description" binding: "required"`
	Status      string    `json:"status" binding: "required"`
	CompletedAt string    `json:"completed_at" binding: "required"`
}
