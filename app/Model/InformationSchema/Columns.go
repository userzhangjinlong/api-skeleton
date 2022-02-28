package InformationSchema

type Columns struct {
	TableCatalog           string `gorm:"column:TABLE_CATALOG" json:"tableCatalog"`
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`
	TableNames             string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"columnName"`
	OrdinalPosition        string `gorm:"column:ORDINAL_POSITION" json:"ordinalPosition"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"isNullable"`
	DataType               string `gorm:"column:DATA_TYPE" json:"dataType"`
	CharacterMaximumLength int    `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"`
	CharacterOctetLength   int    `gorm:"column:CHARACTER_OCTET_LENGTH" json:"characterOctetLength"`
	NumericPrecision       int    `gorm:"column:NUMERIC_PRECISION" json:"numericPrecision"`
	NumericScale           int    `gorm:"column:NUMERIC_SCALE" json:"numericScale"`
	DatetimePrecision      int    `gorm:"column:DATETIME_PRECISION" json:"datetimePrecision"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"characterSetName"`
	CollationName          string `gorm:"column:COLLATION_NAME" json:"collationName"`
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"columnType"`
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"columnKey"`
	Extra                  string `gorm:"column:EXTRA" json:"extra"`
	Privileges             string `gorm:"column:PRIVILEGES" json:"privileges"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"columnComment"`
	GenerationExpression   string `gorm:"column:GENERATION_EXPRESSION" json:"generationExpression"`
}

func (model *Columns) TableName() string {
	return "columns"
}
