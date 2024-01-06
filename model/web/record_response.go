package web

import "time"

type RecordResponse struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks int       `json:"totalMarks"`
}