package querybuilder

import (
	"encoding/json"
	"fmt"
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

	query := ""
	responseMap := map[string]models.QueryRequest{}
	for _, value := range input.Queries {
		_, prevQuery, queryID := containsQuery(value.Tables, responseMap)

		queryString := pattern.buildJoinQuery(value, prevQuery, queryID)
		responseMap[value.QueryID] = models.QueryRequest{
			QueryID:         value.QueryID,
			OutColsDataType: make([]string, len(value.OutAttributes)),
			Query:           queryString,
		}

		query = queryString
	}

	// data := combineSeriesOfQueries(input, responseMap)
	// byteArr, _ := json.Marshal(responseMap)
	return query
}

func containsQuery(tables []string, responseMap map[string]models.QueryRequest) (bool, string, string) {
	for _, val := range tables {
		if value, isExists := responseMap[val]; isExists {
			return true, value.Query, value.QueryID
		}
	}
	return false, "", ""
}

// func combineSeriesOfQueries(input models.Input, responseMap map[int]models.QueryRequest) string {
// 	for _, val := range responseMap {

// 	}
// }

func (pattern *JoinQueryBuilder) buildJoinQuery(queryData models.Query, previousQuery string, previousQueryID string) string {
	//build on prevQuery If it Exists
	pattern.QueryString = embedPreviousQToCTE(previousQuery, previousQueryID)

	namesMap := getTableAliasMap(queryData.Tables)
	pattern.QueryString = pattern.selectOutAttributes(queryData, namesMap)
	flag := true

	for _, value := range queryData.JoinOn {
		pattern.QueryString = pattern.buildSubQuery(value, namesMap, flag)
		flag = false
	}

	query := pattern.getConditionalQueryStr(queryData, namesMap)

	fmt.Println(query)
	return query
}

func embedPreviousQToCTE(previousQuery string, previousQueryID string) string {
	if previousQuery == "" {
		return ""
	}

	return "with " + previousQueryID + " as (" + previousQuery + ")"
}

func (pattern *JoinQueryBuilder) selectOutAttributes(queryData models.Query, namesMap map[string]string) string {
	if len(queryData.OutAttributes) == 0 {
		return pattern.QueryString + "select * from "
	}

	prefixColSelection := "select "
	for index, value := range queryData.OutAttributes {
		prefixColSelection += namesMap[value.TableName] + "." + value.ColumnName
		if index != len(queryData.OutAttributes)-1 {
			prefixColSelection += " as " + value.Alias + ","
		} else {
			prefixColSelection += " from "
		}
	}

	return pattern.QueryString + prefixColSelection
}

func (pattern *JoinQueryBuilder) getConditionalQueryStr(queryData models.Query, namesMap map[string]string) string {
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
