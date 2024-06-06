package models

type Fault struct {
	Turbine     string `json:"turbine"`
	Date        string `json:"date"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
