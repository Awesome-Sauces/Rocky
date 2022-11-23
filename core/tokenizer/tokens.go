package tokenizer

type Token struct {
	Type  TOKEN
	Value string
}

// Create Token object and return by pointer
func CreateToken(TYPE TOKEN, VALUE string) *Token {
	return &Token{Type: TYPE, Value: VALUE}
}

// Token Type Getter
func (t *Token) GetType() TOKEN {
	return t.Type
}

// Value Getter
func (t *Token) GetValue() string {
	return t.Value
}
