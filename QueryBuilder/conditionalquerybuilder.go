package querybuilder

import (
	"encoding/json"
	"log"
	"strings"

	enums "github.com/niroopreddym/caymanislandsquerybuilder/Enums"
	models "github.com/niroopreddym/caymanislandsquerybuilder/Models"
)

//ConditionalQueryBuilder builds the condtional queries on the tablename passed
type ConditionalQueryBuilder struct {
	QueryString string
}

//NewConditionalQueryBuilder is ctor
func NewConditionalQueryBuilder(previousQuery string) *ConditionalQueryBuilder {
	return &ConditionalQueryBuilder{
		QueryString: previousQuery,
	}
}

//GetQueryPattern returns the query that needs to be searched
func (pattern *ConditionalQueryBuilder) GetQueryPattern(queryData []byte) string {
	tableQuery := models.TableQuery{}
	err := json.Unmarshal(queryData, &tableQuery)
	if err != nil {
		log.Panic(err)
	}

	queryString := pattern.buildConditionalQuery(tableQuery)
	return queryString
}

func (pattern *ConditionalQueryBuilder) buildConditionalQuery(queryData models.TableQuery) string {
	if len(queryData.Columns) > 0 {
		pattern.QueryString = pattern.QueryString + " where "
	}

	for index, value := range queryData.Columns {
		pattern.buildSubQuery(value)
		if index != len(queryData.Columns)-1 {
			pattern.QueryString += "and "
		}
	}

	return pattern.QueryString
}

func (pattern *ConditionalQueryBuilder) buildSubQuery(column models.ConditionalQueryData) {
	pattern.withFieldName(column)
	pattern.withOperator(column)
	pattern.withFieldValue(column)
}

func (pattern *ConditionalQueryBuilder) withFieldName(column models.ConditionalQueryData) {
	pattern.QueryString += column.FieldName + " "
}

func (pattern *ConditionalQueryBuilder) withOperator(column models.ConditionalQueryData) {
	operator := strings.ToUpper(column.Operator)
	operatorASCIISymbol := enums.ToOperation(operator)
	pattern.QueryString += operatorASCIISymbol + " "
}

func (pattern *ConditionalQueryBuilder) withFieldValue(column models.ConditionalQueryData) {
	if enums.StringToOperation[strings.ToUpper(column.Operator)] == enums.EQUAL {
		pattern.QueryString += "'" + column.FieldValue + "'"
		return
	}
	pattern.QueryString += column.FieldValue + " "
}
