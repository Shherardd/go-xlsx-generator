package main

import (
	"excel-gen/models"
	"excel-gen/reports"
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

func main() {
	f := excelize.NewFile()

	//FAKE DATA
	registros := []models.Registro{
		{
			Id:             255651,
			CuentaBancaria: "234990038",
			Banco:          "Citi-Mexico (Banamex)",
			Referencia:     "000000000001",
			Importe:        400304.5,
		},
		{
			Id:             255647,
			CuentaBancaria: "234786038",
			Banco:          "Santander",
			Referencia:     "000000000002",
			Importe:        173593.87,
		},
		{
			Id:             255645,
			CuentaBancaria: "234945038",
			Banco:          "AMEX",
			Referencia:     "000000000003",
			Importe:        164504.33,
		},
	}

	totales := models.Totales{
		Importe: 307555061.5500001,
	}

	bancosNoIdentificados := models.BancosNoIdentificados{
		Success:   true,
		Registros: registros,
		Totales:   totales,
	}

	metadata := &models.MetadataReport{
		Company: "Nombre de la empresa",
		Month:   "Enero",
		Year:    "2021",
		Time:    time.Now(),
	}

	reports.CrearExcelBancosNoIdentificadas(f, &bancosNoIdentificados, "BancosNoIdentificados", metadata)
	reports.CrearExcelBancosNoIdentificadas(f, &bancosNoIdentificados, "BancosIdentificados", metadata)
	if err := f.SaveAs("BancosNoIdentificados.xlsx"); err != nil {
		fmt.Println(err)
	}
}
