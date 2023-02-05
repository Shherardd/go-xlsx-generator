package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()

	index := f.NewSheet("Empleados")

	f.SetActiveSheet(index)

	f.SetCellValue("Empleados", "A1", "ID")
	f.SetCellValue("Empleados", "B1", "Nombre")
	f.SetCellValue("Empleados", "C1", "Correo electrónico")

	employees := []struct {
		ID    int
		Name  string
		Email string
	}{
		{1, "John Doe", "johndoe@example.com"},
		{2, "Jane Doe", "janedoe@example.com"},
		{3, "Bob Smith", "bobsmith@example.com"},
	}

	for i, employee := range employees {
		row := i + 2
		f.SetCellValue("Empleados", fmt.Sprintf("A%d", row), employee.ID)
		f.SetCellValue("Empleados", fmt.Sprintf("B%d", row), employee.Name)
		f.SetCellValue("Empleados", fmt.Sprintf("C%d", row), employee.Email)
	}

	if err := f.SaveAs("employees.xlsx"); err != nil {
		fmt.Println(err)
	}
}
