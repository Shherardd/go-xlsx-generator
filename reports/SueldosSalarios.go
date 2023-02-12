package reports

import (
	"excel-gen/models"
	"excel-gen/reports/styles"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func CrearExcelBancosNoIdentificadas(spreadsheet *excelize.File, datos *models.BancosNoIdentificados, title string,
	metadata *models.MetadataReport) {
	//														//Initial Values
	sheetTitle := "Partidas Bancarias No Identificadas"
	headRowsCount := 3
	headersRow := 5
	initialDataRow := 6
	lastSheetColumn := "D"
	//														//Style declaration
	moneyFormatStyle, _ := spreadsheet.NewStyle(styles.DataMoneyFormat)
	headStyle, _ := spreadsheet.NewStyle(styles.HeadStyles)
	HeadersStyle, _ := spreadsheet.NewStyle(styles.HeaderStyles)
	totalStyle, _ := spreadsheet.NewStyle(styles.TotalesStyles)

	//														//Head gen
	spreadsheet.NewSheet(title)
	spreadsheet.SetCellValue(title, "A1", metadata.Company)
	spreadsheet.SetCellValue(title, "A2", sheetTitle)
	spreadsheet.SetCellValue(title, "A3", metadata.Month+" de "+metadata.Year+" "+
		metadata.Time.Format("15:04:05"))

	// 														//Headers
	spreadsheet.SetCellValue(title, "A5", "Cuenta Bancaria")
	spreadsheet.SetCellValue(title, "B5", "Banco")
	spreadsheet.SetCellValue(title, "C5", "Referencia")
	spreadsheet.SetCellValue(title, "D5", "Importe")

	//														//Data gen
	cell := initialDataRow
	for _, value := range datos.Registros {
		row := strconv.Itoa(cell)
		spreadsheet.SetCellValue(title, "A"+row, value.CuentaBancaria)
		spreadsheet.SetCellValue(title, "B"+row, value.Banco)
		spreadsheet.SetCellValue(title, "C"+row, value.Referencia)
		spreadsheet.SetCellValue(title, "D"+row, value.Importe)
		cell++
	}
	lastDataRow := cell - 1

	cell++
	i := strconv.Itoa(cell)
	spreadsheet.SetCellValue(title, "C"+i, "Total:")
	spreadsheet.SetCellValue(title, "D"+i, datos.Totales.Importe)

	//Merge Cells
	spreadsheet.MergeCell(title, "A1", lastSheetColumn+"1")
	spreadsheet.MergeCell(title, "A2", lastSheetColumn+"2")
	spreadsheet.MergeCell(title, "A3", lastSheetColumn+"3")

	//Set Totales
	spreadsheet.SetCellStyle(title, "C"+i, "D"+i, totalStyle)

	//Set Money Format columns
	j := strconv.Itoa(lastDataRow)
	spreadsheet.SetCellStyle(title, "D6", "D"+j, moneyFormatStyle)

	//Set Column Width
	spreadsheet.SetColWidth(title, "A", lastSheetColumn, 20)

	//Apply style to cells
	vCellHead := lastSheetColumn + strconv.Itoa(headRowsCount)
	vCellHeader := lastSheetColumn + strconv.Itoa(headersRow)
	strHeadersRow := strconv.Itoa(headersRow)
	spreadsheet.SetCellStyle(title, "A1", vCellHead, headStyle)
	spreadsheet.SetCellStyle(title, "A"+strHeadersRow, vCellHeader, HeadersStyle)

}
