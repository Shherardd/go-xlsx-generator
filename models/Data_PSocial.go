package models

type BancosNoIdentificados struct {
	Success   bool       `json:"success"`
	Registros []Registro `json:"registros"`
	Totales   Totales    `json:"totales"`
}

type Registro struct {
	Id             int     `json:"id"`
	CuentaBancaria string  `json:"cuenta_bancaria"`
	Banco          string  `json:"banco"`
	Referencia     string  `json:"referencia"`
	Importe        float64 `json:"importe"`
}

type Totales struct {
	Importe float64 `json:"importe"`
}
