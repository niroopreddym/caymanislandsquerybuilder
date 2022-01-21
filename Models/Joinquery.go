package models

//Input ...
type Input struct {
	Queries []Query `json:"queries"`
}

//Query master structure
type Query struct {
	QueryID       string                 `json:"queryId"`
	OutAttributes []OutAttributes        `json:"outAttributes"`
	Tables        []string               `json:"tables"`
	JoinOn        []JoinOn               `json:"joinOn"`
	Where         []ConditionalQueryData `json:"where"`
}

//JoinOn ..
type JoinOn struct {
	JoinType string `json:"joinType"`
	Key      Key    `json:"key"`
	Value    Key    `json:"value"`
}

//OutAttributes ...
type OutAttributes struct {
	TableName  string `json:"table"`
	ColumnName string `json:"column"`
	Alias      string `json:"alias"`
}

//Key joinKey
type Key struct {
	TableName  string `json:"table"`
	ColumnName string `json:"column"`
}
