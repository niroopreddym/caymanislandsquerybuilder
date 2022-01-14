package handlers

import (
	"io/ioutil"
	"log"

	querybuilder "github.com/niroopreddym/caymanislandsquerybuilder/QueryBuilder"
)

//AssignemntHandler handles the assignment biz layer
type AssignemntHandler struct {
	JoinBuilder             querybuilder.IQueryBuilder
	ConditionalQueryBuilder querybuilder.IQueryBuilder
}

//NewHandler ctor
func NewHandler() *AssignemntHandler {
	return &AssignemntHandler{
		JoinBuilder:             querybuilder.NewJoinBuilder(""),
		ConditionalQueryBuilder: querybuilder.NewConditionalQueryBuilder(""),
	}
}

//Assignment1 ...
func (handler *AssignemntHandler) Assignment1() string {
	queryData, err := ioutil.ReadFile("query.json")
	if err != nil {
		log.Panic(err)
	}

	str := handler.ConditionalQueryBuilder.GetQueryPattern(queryData)
	return str
}

//Assignment2 ...
func (handler *AssignemntHandler) Assignment2() string {
	queryData, err := ioutil.ReadFile("query2.json")
	if err != nil {
		log.Panic(err)
	}

	str := handler.JoinBuilder.GetQueryPattern(queryData)
	return str
}
