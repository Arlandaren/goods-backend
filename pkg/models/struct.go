package models

import "time"

type Good struct {
	ID          uint   `json:"id" db:"id"`
	ProjectId   uint   `json:"projectId" db:"project_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Priority    int    `json:"priority" db:"priority"`
	Removed     bool   `json:"removed" db:"removed"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}
type Meta struct {
    Limit   int `json:"limit"`
    Offset  int `json:"offset"`
    Removed int `json:"removed"`
    Total   int `json:"total"`
}
