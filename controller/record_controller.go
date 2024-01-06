package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RecordController interface {
	FindRecords(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type RecordControllerImpl struct {
	RecordService service.RecordService
}

func NewRecordController(RecordService service.RecordService) RecordController {
	return &RecordControllerImpl{
		RecordService: RecordService,
	}
}


func (controller RecordControllerImpl) FindRecords(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	recordRequest := web.RecordRequest{}
	helper.ReadFromRequestBody(request, &recordRequest)


	recordResponse := controller.RecordService.FindRecords(request.Context(), recordRequest)
	webResponse := web.WebResponse{
		Code:  0,
		Status: "Success",
		Records:   recordResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
