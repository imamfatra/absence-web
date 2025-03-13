package repository

import (
	"absensi-web/model"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type AbsensiRepository interface {
    Save(ctx context.Context, tx *sql.Tx, absensi model.Absensi) (model.Absensi, error)
}

type AbsensiRepositoryImpl struct {}

func NewAbsensiRepository() AbsensiRepository {
    return &AbsensiRepositoryImpl{}
}

func (r *AbsensiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, absensi model.Absensi) (model.Absensi, error) {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    var id int
    query := "INSERT INTO students (name, nim, mata_kuliah, jurusan) VALUES ($1, $2, $3, $4) RETURNING id"
    err := tx.QueryRowContext(ctx, query, absensi.Name, absensi.Nim, absensi.MataKuliah, absensi.Jurusan).Scan(&id)
    if err != nil {
        return model.Absensi{}, fmt.Errorf("failed save data: %v", err)
    }

    absensi.Id = id
    return absensi, nil
}
