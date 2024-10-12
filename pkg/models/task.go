package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          uuid.UUID `json:"id" binding: "not required"`
	Name        string    `json:"name" binding: "required"`
	Description string    `json:"description" binding: "required"`
	Priority    int       `json:"priority" binding: "required"`
	Status      int       `json:"status" binding: "required"`
	CreatedAt   time.Time `json:"created_at" binding: "required"`
	CompletedAt time.Time `json:"completed_at" binding: "required"`
	ProjectId   uuid.UUID `json:"project_id" binding: "not required"`
}
