package models

import "time"

type ReportRequest struct {
	Empresa       string        `json:"empresa"`
	Periodo       string        `json:"periodo"`
	Ejericio      string        `json:"ejercicio"`
	IdControl     int           `json:"idControl"`
	NameEmpresa   string        `json:"nameEmpresa"`
	DatosCaratula DatosCaratula `json:"datosCaratula"`
}

type DatosCaratula struct {
	Renglones []Renglon `json:"renglones"`
}

type Renglon struct {
	Descripcion string `json:"descripcion"`
}

type MetadataReport struct {
	Company string
	Month   string
	Year    string
	Time    time.Time
}
