package models

//Input ...
type Input struct {
	Tables []string               `json:"tables"`
	JoinOn []JoinOn               `json:"joinOn"`
	Where  []ConditionalQueryData `json:"where"`
}

//JoinOn ..
type JoinOn struct {
	JoinType string `json:"joinType"`
	Key      Key    `json:"key"`
	Value    Key    `json:"value"`
}

//Key joinKey
type Key struct {
	TableName  string `json:"table"`
	ColumnName string `json:"column"`
}
