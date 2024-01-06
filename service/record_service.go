package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type RecordService interface {
	FindRecords(ctx context.Context, recordRequest web.RecordRequest) []web.RecordResponse
}

type RecordServiceImpl struct {
	RecordRepository repository.RecordRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewRecordService(recordRepository repository.RecordRepository, DB *sql.DB, validate *validator.Validate) RecordService {
	return &RecordServiceImpl{
		RecordRepository: recordRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service RecordServiceImpl) FindRecords(ctx context.Context, recordRequest web.RecordRequest) []web.RecordResponse{
	tx, err := service.DB.Begin()
	helper.PanicifError(err)
	defer helper.CommitOrRollback(tx)

	records, err := service.RecordRepository.FindRecords(ctx, tx, recordRequest)
	helper.PanicifError(err)

	return helper.ToRecordResponses(records)
}
