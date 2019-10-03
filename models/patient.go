package models

import (

)

type Patient struct {
	ID			string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName	string `json:"lastName"`
	Diagnosis	string `json:"diagnosis"`
	Physician	string `json:"physician"`
	DOV			string `json:"dov"`
}