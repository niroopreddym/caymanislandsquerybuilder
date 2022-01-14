package querybuilder

//IQueryBuilder exposes the methods that are needed to be implemented for query building
type IQueryBuilder interface {
	GetQueryPattern(queryData []byte) string
}
