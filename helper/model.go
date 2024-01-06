package helper

import (
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
)


func ToRecordResponse(record domain.Record) web.RecordResponse {
	return web.RecordResponse{
		ID:         record.ID,
		CreatedAt:  record.CreatedAt,
		TotalMarks: record.TotalMarks,
	}
}

func ToRecordResponses(records []domain.Record) []web.RecordResponse {
	var recordResponses []web.RecordResponse
	for _, record := range records {
		recordResponses = append(recordResponses, ToRecordResponse(record))
	}
	return recordResponses
}
