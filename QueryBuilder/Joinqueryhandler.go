package querybuilder

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	models "github.com/niroopreddym/caymanislandsquerybuilder/Models"
)

//JoinQueryBuilder builds the join queries on tables list passed
type JoinQueryBuilder struct {
	QueryString string
}

//NewJoinBuilder is ctor
func NewJoinBuilder(previousQuery string) *JoinQueryBuilder {
	return &JoinQueryBuilder{
		QueryString: previousQuery,
	}
}

//GetQueryPattern returns the query that needs to be searched
func (pattern *JoinQueryBuilder) GetQueryPattern(queryData []byte) string {
	input := models.Input{}
	err := json.Unmarshal(queryData, &input)
	if err != nil {
		log.Panic(err)
	}

	queryString := pattern.buildJoinQuery(input)
	return queryString
}

func (pattern *JoinQueryBuilder) buildJoinQuery(queryData models.Input) string {
	namesMap := getTableAliasMap(queryData.Tables)
	pattern.QueryString = "select * from "
	flag := true
	for _, value := range queryData.JoinOn {
		pattern.QueryString = pattern.buildSubQuery(value, namesMap, flag)
		flag = false
	}

	query := pattern.getConditionalQueryStr(queryData, namesMap)
	return query
}

func (pattern *JoinQueryBuilder) getConditionalQueryStr(queryData models.Input, namesMap map[string]string) string {
	conditionalQueryBuilder := NewConditionalQueryBuilder(pattern.QueryString)

	data := alterTheTableNameToAlias(queryData.Where, namesMap)
	tableQuery := models.TableQuery{
		Columns: data,
	}
	byteArr, _ := json.Marshal(tableQuery)

	return conditionalQueryBuilder.GetQueryPattern(byteArr)
}

func alterTheTableNameToAlias(whereClasueData []models.ConditionalQueryData, namesMap map[string]string) []models.ConditionalQueryData {
	newData := []models.ConditionalQueryData{}
	for _, value := range whereClasueData {
		if strings.Contains(value.FieldName, ".") {
			tblColumnArr := strings.Split(value.FieldName, ".")
			tableName := tblColumnArr[0]
			value.FieldName = strings.ReplaceAll(value.FieldName, tableName, namesMap[tableName])
		}

		if strings.Contains(value.FieldValue, ".") {
			tblColumnArr := strings.Split(value.FieldValue, ".")
			tableName := tblColumnArr[0]
			value.FieldValue = strings.ReplaceAll(value.FieldValue, tableName, namesMap[tableName])
		}

		newData = append(newData, value)
	}

	return newData
}

func getTableAliasMap(tableNames []string) map[string]string {
	namesMap := map[string]string{}
	for index, tableName := range tableNames {
		namesMap[tableName] = "t" + strconv.Itoa(index)
	}

	return namesMap
}

func (pattern *JoinQueryBuilder) buildSubQuery(queryData models.JoinOn, namesMap map[string]string, flag bool) string {
	if flag {
		pattern.QueryString += queryData.Key.TableName + " " + namesMap[queryData.Key.TableName] + " "
	}

	pattern.QueryString += queryData.JoinType + " join " + queryData.Value.TableName + " " + namesMap[queryData.Value.TableName] + " on "
	pattern.QueryString += namesMap[queryData.Key.TableName] + "." + queryData.Key.ColumnName + " "
	pattern.QueryString += "= "
	pattern.QueryString += namesMap[queryData.Value.TableName] + "." + queryData.Value.ColumnName + " "
	return pattern.QueryString
}
