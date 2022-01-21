package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/niroopreddym/caymanislandsquerybuilder/database"
)

//DatabaseService is the class implementation for ProductServicesIface interface
type DatabaseService struct {
	DatabaseService database.DbIface
}

//NewDatabaseServicesInstance instantiates the struct
func NewDatabaseServicesInstance() *DatabaseService {
	return &DatabaseService{
		DatabaseService: database.DBNewHandler(),
	}
}

//Read reads the data from DB
func (service *DatabaseService) Read(query string) ([]map[int]string, error) {
	defer service.DatabaseService.DbClose()

	db, err := service.DatabaseService.InitDbReader()
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	lstResults := []map[int]string{}
	for rows.Next() {
		cols1, err := rows.Values()
		fmt.Println(err)
		resultMap := map[int]string{}
		for index, val := range cols1 {
			val1 := interface{}(val)
			switch t := val1.(type) {
			case int, int32, int64:
				fmt.Println(t)
				resultMap[index] = strconv.Itoa(int(t.(int32)))
			case string:
				fmt.Println(t)
				resultMap[index] = t
			default:
				fmt.Println(t)
			}
		}

		lstResults = append(lstResults, resultMap)
	}

	fmt.Println(lstResults)

	return lstResults, nil
}
