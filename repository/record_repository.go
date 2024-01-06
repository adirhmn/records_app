package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	"context"
	"database/sql"
)

type RecordRepository interface {
	FindRecords(ctx context.Context, tx *sql.Tx, recordRequest web.RecordRequest) ([]domain.Record, error) 
}

type RecordRepositoryImpl struct {
}

func NewRecordRepository() RecordRepository {
	return &RecordRepositoryImpl{}
}

func (r *RecordRepositoryImpl) FindRecords(ctx context.Context, tx *sql.Tx, recordRequest web.RecordRequest) ([]domain.Record, error) {
	SQL := `
	SELECT 
		id, 
		createdAt, 
		(
			SELECT SUM(value) 
			FROM UNNEST(records.marks) AS t(value)
		) AS total_marks 
	FROM records 
	WHERE 
		(
			SELECT SUM(value) 
			FROM UNNEST(records.marks) AS t(value)
		) BETWEEN $1 AND $2 
		AND records.createdAt BETWEEN $3 AND $4;
	`
	minCount := recordRequest.MinCount
	maxCount := recordRequest.MaxCount
	startDate := recordRequest.StartDate
	endDate := recordRequest.EndDate

	rows, err := tx.QueryContext(ctx, SQL, minCount, maxCount, startDate, endDate)
	helper.PanicifError(err)
	defer rows.Close()

	var records []domain.Record
	for rows.Next() {
		record := domain.Record{}
		err := rows.Scan(&record.ID, &record.CreatedAt, &record.TotalMarks)
		helper.PanicifError(err)
		records = append(records, record)
	}
	return records, nil
}
