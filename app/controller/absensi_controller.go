package controller

import (
	"absensi-web/model"
	"absensi-web/service"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AbsensiController struct {
    AbsensiService service.AbsensiService
}

func NewAbsensiController(absensiService service.AbsensiService) *AbsensiController  {
    return &AbsensiController{absensiService}
}

func writeToResponseBody(w http.ResponseWriter, response interface{}) {
    w.Header().Add("Content-Type", "application/json")
    encoder := json.NewEncoder(w)
    err := encoder.Encode(response)
    if err != nil {
        panic(err)
    }
}

func (c *AbsensiController) Save(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    var absensiRequest model.AbsensiRequest

    defer r.Body.Close()

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&absensiRequest)
    if err != nil {
        webResponse := model.WebResponse {
            Code: http.StatusBadRequest,
            Status: "error",
            Data: "Invalid request body",
            // Data: err.Error(),
        }
        writeToResponseBody(w, webResponse)
        return
    }

    absensiResponse, err := c.AbsensiService.Save(r.Context(), absensiRequest)
    if err != nil {
         webResponse := model.WebResponse {
            Code: http.StatusInternalServerError,
            Status: "error",
            Data: err.Error(),
        }
        writeToResponseBody(w, webResponse)
        return

    }

    webResponse := model.WebResponse{
        Code: 200,
        Status: "success",
        Data: absensiResponse,
    }
    writeToResponseBody(w, webResponse)
}
