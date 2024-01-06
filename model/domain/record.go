package domain

import "time"

type Record struct {
	ID         int       
	CreatedAt  time.Time 
	TotalMarks int       
}