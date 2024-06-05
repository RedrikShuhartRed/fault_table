package models

type Fault struct {
	id          int    `json:"id"`
	Turbine     int    `json:"turbine"`
	Date        string `json:"date"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
