package config

import (
	"api-skeleton/app/Model"
	"api-skeleton/app/Model/InformationSchema"
)

var (
	Default            []interface{}
	InformationSchema2 []interface{}
)

func init() {
	InformationSchema2 = append(
		InformationSchema2,
		&InformationSchema.Columns{},
		"COLUMNS",
	)
	Default = append(Default, &Model.User{})
}
