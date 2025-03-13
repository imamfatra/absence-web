package model

type AbsensiRequest struct {
    Name string `validate:"required,min=3,max=255" json:"name"`
    Nim int `validate:"min=1000" json:"nim"`
    MataKuliah string `validate:"required,min=3,max=255" json:"mata_kuliah"`
    Jurusan string `validate:"required,min=3,max=255" json:"jurusan"`
}
