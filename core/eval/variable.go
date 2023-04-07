package eval

// Variable Registrar
var variables = make(map[string]*Variable)

// Function to register new variable
func RegisterVariable(variable *Variable) {
	variables[variable.Name] = variable
}

// Get a Variable
func GetVariable(name string) *Variable {
	return variables[name]
}

type Variable struct {
	Name string
	Type VTYPE
	Data string
}

func NewVariable(Name string, Type VTYPE, Data string) *Variable {
	return &Variable{Name: Name, Type: Type, Data: Data}
}

func NewVType(Type string) VTYPE {
	switch Type {
	case "int":
		return INT
	case "String":
		return STRING
	case "double":
		return DOUBLE
	case "class":
		return CLASS
	default:
		return NONE
	}
}

// TOKEN - Custom type to hold value for token type
type VTYPE int

// Declare related constants for each Token type starting with index 1
const (
	STRING VTYPE = iota + 1 // EnumIndex = 1
	INT                     // EnumIndex = 2
	CLASS                   // EnumIndex = 3
	DOUBLE                  // EnumIndex = 4
	NONE                    // EnumIndex = 3
)

func VTypeFromString(tk string) VTYPE {
	if len(tk) == 0 || tk == " " {
		return NONE
	}

	switch tk {
	case "STRING":
		return STRING
	case "INT":
		return INT
	case "CLASS":
		return CLASS
	case "DOUBLE":
		return DOUBLE
	default:
		return NONE
	}
}

// Converting Enum to string
func (vtype VTYPE) ToString() string {
	switch vtype {
	case STRING:
		return "STRING"
	case INT:
		return "INT"
	case DOUBLE:
		return "DOUBLE"
	case CLASS:
		return "CLASS"
	case NONE:
		return "NONE"
	default:
		return "ERROR"
	}
}
