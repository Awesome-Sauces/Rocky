package tokenizer

type Token struct {
	Type  TOKEN
	Value string
	Order int
}

// Create Token object and return by pointer
func CreateToken(TYPE TOKEN, VALUE string, ORDER int) *Token {
	return &Token{Type: TYPE, Value: VALUE, Order: ORDER}
}

// Token Type Getter
func (t *Token) GetType() TOKEN {
	return t.Type
}

// Value Getter
func (t *Token) GetValue() string {
	return t.Value
}

// Order Getter
func (t *Token) GetOrder() int {
	return t.Order
}

// Set Order
func (t *Token) SetOrder(ORDER int) {
	t.Order = ORDER
}
