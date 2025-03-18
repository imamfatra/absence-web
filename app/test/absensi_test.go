package test

import (
	"absensi-web/cmd"
	"absensi-web/controller"
	"absensi-web/repository"
	"absensi-web/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	_ "github.com/lib/pq"
)

func loadEnv() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("../.env")
    config.SetConfigType("env")
	config.AddConfigPath("../")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error load .env file: %v", err)
	}

    return config
}

func NewDB() *sql.DB {
    config := loadEnv()

	host := config.GetString("DB_HOST")
	port := config.GetInt("DB_PORT")
	username := config.GetString("DB_USER")
	password := config.GetString("DB_PASSWORD")
	dbName := config.GetString("DB_NAME")

	psqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbName)
	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		log.Fatalf("Connected database error, %v", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
    validate := validator.New()

    absensiRepository := repository.NewAbsensiRepository()
    absenseService := service.NewAbsensiService(absensiRepository, db, validate)
    absensiController := controller.NewAbsensiController(absenseService)
    router := cmd.NewRouter(*absensiController)

    return router
}

func TestSaveSuccess(t *testing.T)  {
    db := NewDB()
    router := setupRouter(db)

    err := db.Ping()
    assert.NoError(t, err)

    requestBody := strings.NewReader((`{"name": "riski", "nim": 489354, "mata_kuliah": "agama", "jurusan": "pendidikan pkn"}`))
    request := httptest.NewRequest(http.MethodPost, "/", requestBody)
    request.Header.Add("Content-Type", "application/json")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    response := recorder.Result()
    defer response.Body.Close()
    assert.Equal(t, http.StatusOK, response.StatusCode)

    body, err := io.ReadAll(response.Body)
    assert.NoError(t, err)

    var responseBody map[string]interface{}
    err = json.Unmarshal(body, &responseBody)
    assert.NoError(t, err)

    assert.Equal(t, 200, int(responseBody["code"].(float64)))
    assert.Equal(t, "success", responseBody["status"])
}
