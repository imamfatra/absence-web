package main

import (
	"absensi-web/cmd"
	"absensi-web/controller"
	"absensi-web/db"
	"absensi-web/middleware"
	"absensi-web/repository"
	"absensi-web/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main()  {
    db := db.NewDB()
    validator := validator.New()

    absensiRepository := repository.NewAbsensiRepository()
    absensiService := service.NewAbsensiService(absensiRepository, db, validator)
    absensiController := controller.NewAbsensiController(absensiService)
    router := cmd.NewRouter(*absensiController)

    directory := http.Dir("./static")
    fileServer := http.FileServer(directory)
    router.Handler("GET", "/static/*filepath", http.StripPrefix("/static/", fileServer))    

    server := http.Server {
        Addr: ":3000",
        Handler: middleware.EnableCors(router),
    }
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
