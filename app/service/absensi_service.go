package service

import (
	"absensi-web/model"
	"absensi-web/repository"
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AbsensiService interface {
    Save(ctx context.Context, request model.AbsensiRequest) (model.AbsensiResponse, error)
}

type AbsensiServiceImpl struct {
    AbsensiRepository repository.AbsensiRepository
    DB *sql.DB
    Validate *validator.Validate
}

func NewAbsensiService(absensiRepository repository.AbsensiRepository, db *sql.DB, validate *validator.Validate) AbsensiService {
    return &AbsensiServiceImpl{
        AbsensiRepository: absensiRepository,
        DB: db,
        Validate: validate,
    }
}

func (s *AbsensiServiceImpl) Save(ctx context.Context, request model.AbsensiRequest) (model.AbsensiResponse, error) {
    err := s.Validate.Struct(request)
    if err != nil {
        return model.AbsensiResponse{}, fmt.Errorf("validation failed, err: %v", err)
    }

    tx, err := s.DB.Begin()
    if err != nil {
        return model.AbsensiResponse{}, fmt.Errorf("Failed start transaktion, err: %v", err)
    }
    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

    absensi := model.Absensi {
        Name: request.Name,
        Nim: request.Nim,
        MataKuliah: request.MataKuliah,
        Jurusan: request.Jurusan,
    }
    absensi, err = s.AbsensiRepository.Save(ctx, tx, absensi)
    if err != nil {
        return model.AbsensiResponse{}, fmt.Errorf("Failed save data, err: %v", err)
    }

    if tx.Commit(); err != nil {
        return model.AbsensiResponse{}, fmt.Errorf("Failed commit data, err: %v", err)
    }

    return model.AbsensiResponse{Id: absensi.Id}, nil
}
