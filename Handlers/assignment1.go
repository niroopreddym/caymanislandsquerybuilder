package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	models "github.com/niroopreddym/caymanislandsquerybuilder/Models"
	querybuilder "github.com/niroopreddym/caymanislandsquerybuilder/QueryBuilder"
	"github.com/niroopreddym/caymanislandsquerybuilder/services"
)

//AssignemntHandler handles the assignment biz layer
type AssignemntHandler struct {
	JoinBuilder             querybuilder.IQueryBuilder
	ConditionalQueryBuilder querybuilder.IQueryBuilder
	DatabaseService         services.ISQLService
	FileName                string
}

//NewHandler ctor
func NewHandler(inputFileName string) *AssignemntHandler {
	return &AssignemntHandler{
		JoinBuilder:             querybuilder.NewJoinBuilder(""),
		ConditionalQueryBuilder: querybuilder.NewConditionalQueryBuilder(""),
		DatabaseService:         services.NewDatabaseServicesInstance(),
		FileName:                inputFileName,
	}
}

//Assignment1 ...
func (handler *AssignemntHandler) Assignment1() string {
	queryData, err := ioutil.ReadFile(handler.FileName)
	if err != nil {
		log.Panic(err)
	}

	str := handler.ConditionalQueryBuilder.GetQueryPattern(queryData)
	return str
}

//Assignment2 ...
func (handler *AssignemntHandler) Assignment2() string {
	queryData, err := ioutil.ReadFile(handler.FileName)
	if err != nil {
		log.Panic(err)
	}

	queryRequest := handler.JoinBuilder.GetQueryPattern(queryData)

	queryResponse := handler.queryTheSQLDB(queryRequest, queryData)
	response, _ := json.Marshal(queryResponse)
	return string(response)
}

func (handler *AssignemntHandler) queryTheSQLDB(queryRequest string, queryData []byte) []map[string]string {
	data, err := handler.DatabaseService.Read(queryRequest)
	fmt.Println(err)

	//map the data to outAttributes
	response := mapOutVariablesToResponse(data, queryData)
	return response
}

func mapOutVariablesToResponse(data []map[int]string, queryData []byte) []map[string]string {
	input := models.Input{}
	err := json.Unmarshal(queryData, &input)
	if err != nil {
		fmt.Println(err)
	}

	requiredOutAttributes := input.Queries[len(input.Queries)-1].OutAttributes

	response := []map[string]string{}
	for _, dataMap := range data {
		individualResponse := map[string]string{}
		for index, value := range dataMap {
			individualResponse[requiredOutAttributes[index].Alias] = value
		}
		response = append(response, individualResponse)
	}

	return response
}
