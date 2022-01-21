package handlers

import (
	"fmt"
	"io/ioutil"
	"log"

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

	handler.queryTheSQLDB(queryRequest)
	return queryRequest
}

//Response ...

func (handler *AssignemntHandler) queryTheSQLDB(queryRequest string) {
	data, err := handler.DatabaseService.Read(queryRequest)
	fmt.Println(err)
	fmt.Println(data)
}
