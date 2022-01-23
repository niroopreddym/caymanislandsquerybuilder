package services

//ISQLService provides the data interface
type ISQLService interface {
	Read(query string) ([]map[int]string, error)
}

//INoSQLService provides the data interface
type INoSQLService interface {
	Read(query string) ([]map[int]string, error)
}

//IDataConnector ...
type IDataConnector interface {
	ISQLService
	INoSQLService
}
