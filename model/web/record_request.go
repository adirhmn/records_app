package web

type RecordRequest struct {
	StartDate string `validate:"required,date" json:"startDate"`
	EndDate   string `validate:"required,date" json:"endDate"`
	MinCount  int    `validate:"required,numeric" json:"minCount"`
	MaxCount  int    `validate:"required,numeric" json:"maxCount"`
}
