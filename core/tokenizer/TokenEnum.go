package tokenizer

// TOKEN - Custom type to hold value for token type
type TOKEN int

// Declare related constants for each Token type starting with index 1
const (
	ADD            TOKEN = iota + 1 // EnumIndex = 1
	SUBTRACT                        // EnumIndex = 2
	DIVIDE                          // EnumIndex = 3
	MULTIPLY                        // EnumIndex = 4
	LPAREN                          // EnumIndex = 5
	RPAREN                          // EnumIndex = 6
	RBRACKET                        // EnumIndex = 7
	LBRACKET                        // EnumIndex = 8
	NUMBER                          // EnumIndex = 9
	IDENTIFIER                      // EnumIndex = 10
	TYPE                            // EnumIndex = 11
	COMMENT                         // EnumIndex = 12
	RSQUAREBRACKET                  // EnumIndex = 13
	LSQUAREBRACKET                  // EnumIndex = 14
)

// Converting Enum to string
func (tk TOKEN) ToString() string {
	switch tk {
	case ADD:
		return "ADD"
	case SUBTRACT:
		return "SUBTRACT"
	case DIVIDE:
		return "DIVIDE"
	case MULTIPLY:
		return "MULTIPLY"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case RBRACKET:
		return "RBRACKET"
	case LBRACKET:
		return "LBRACKET"
	case NUMBER:
		return "NUMBER"
	case IDENTIFIER:
		return "IDENTIFIER"
	case TYPE:
		return "TYPE"
	case COMMENT:
		return "COMMENT"
	case RSQUAREBRACKET:
		return "RSQUAREBRACKET"
	case LSQUAREBRACKET:
		return "LSQUAREBRACKET"
	default:
		return "ERROR"
	}
}
