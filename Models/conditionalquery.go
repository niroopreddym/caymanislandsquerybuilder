package models

//TableQuery ...
type TableQuery struct {
	TableName string
	Columns   []ConditionalQueryData
}

//ConditionalQueryData inner struct for building the query operations
type ConditionalQueryData struct {
	Operator   string `json:"operator"`
	FieldName  string `json:"fieldName"`
	FieldValue string `json:"fieldValue"`
}
