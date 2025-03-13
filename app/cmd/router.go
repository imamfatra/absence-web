package cmd

import (
	"absensi-web/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(absenController controller.AbsensiController) *httprouter.Router {
    router := httprouter.New()

    router.POST("/", absenController.Save)
    router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        http.ServeFile(w, r, "./templates/index.html")
    })
    return router
}
