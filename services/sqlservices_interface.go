package services

//ISQLService provides the data interface
type ISQLService interface {
	Read(query string) ([]map[int]string, error)
}
