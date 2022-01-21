package models

//QueryRequest input to perform the DB operation
type QueryRequest struct {
	QueryID         string
	OutColsDataType []string
	Query           string
}
