package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	app.InitDB(db)
	recordRepository := repository.NewRecordRepository()
	recordService := service.NewRecordService(recordRepository, db, validate)
	recordController := controller.NewRecordController(recordService)

	router := app.NewRouter(recordController)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicifError(err)

}
