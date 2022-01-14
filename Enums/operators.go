package enums

// Operator limits the operators used in the query builder
type Operator int8

// State transition operations
const (
	IN Operator = iota
	LIKE
	EQUAL
	LESSTHAN
	GREATERTHAN
	LESSTHANEQUAL
	GREATERTHANEQUAL
	NOTEQUAL
	BETWEEN
)

var operationToString = map[Operator]string{
	IN:               "IN",
	LIKE:             "LIKE",
	EQUAL:            "=",
	LESSTHAN:         "<",
	GREATERTHAN:      ">",
	LESSTHANEQUAL:    "<=",
	GREATERTHANEQUAL: ">=",
	NOTEQUAL:         "<>",
	BETWEEN:          "BETWEEN",
}

//StringToOperation returns the map of available operations
var StringToOperation = map[string]Operator{
	"IN":               IN,
	"LIKE":             LIKE,
	"EQUAL":            EQUAL,
	"LESSTHAN":         LESSTHAN,
	"GREATERTHAN":      GREATERTHAN,
	"LESSTHANEQUAL":    LESSTHANEQUAL,
	"GREATERTHANEQUAL": GREATERTHANEQUAL,
	"NOTEQUAL":         NOTEQUAL,
	"BETWEEN":          BETWEEN,
}

func (o Operator) String() string {
	return operationToString[o]
}

//ToOperation converts teh data operation to str
func ToOperation(str string) string {
	operator := StringToOperation[str]
	return operator.String()
}
