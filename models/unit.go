package models


type Unit struct {
	ID            string  `json:"id"`
	Unit          *string `json:"unit"`
	IndicatesTime *bool   `json:"indicatesTime"`
}
