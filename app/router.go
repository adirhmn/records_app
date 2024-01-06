package app

import (
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(recordController controller.RecordController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/records", recordController.FindRecords)

	router.PanicHandler = exception.ErrorHandler

	return router

}
