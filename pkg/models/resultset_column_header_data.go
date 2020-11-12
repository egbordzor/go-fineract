package models

// ResultsetColumnHeaderData struct for ResultsetColumnHeaderData
type ResultsetColumnHeaderData struct {
	ColumnName            string `json:"columnName,omitempty"`
	ColumnType            string `json:"columnType,omitempty"`
	ColumnLength          int64  `json:"columnLength,omitempty"`
	ColumnDisplayType     string `json:"columnDisplayType,omitempty"`
	ColumnCode            string `json:"columnCode,omitempty"`
	Optional              bool   `json:"optional,omitempty"`
	String                bool   `json:"string,omitempty"`
	DateDisplayType       bool   `json:"dateDisplayType,omitempty"`
	DateTimeDisplayType   bool   `json:"dateTimeDisplayType,omitempty"`
	IntegerDisplayType    bool   `json:"integerDisplayType,omitempty"`
	DecimalDisplayType    bool   `json:"decimalDisplayType,omitempty"`
	BooleanDisplayType    bool   `json:"booleanDisplayType,omitempty"`
	CodeValueDisplayType  bool   `json:"codeValueDisplayType,omitempty"`
	CodeLookupDisplayType bool   `json:"codeLookupDisplayType,omitempty"`
	Mandatory             bool   `json:"mandatory,omitempty"`
}
