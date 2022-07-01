package me.AwesomeSauce.Lexer;

public enum Type {
    // This Scheme-like language has three token types:
    // open parens, close parens, and an "atom" type
    LPAREN, RPAREN, ATOM, EQUAL, ADDITION, SUBTRACTION, DIVISION, MULTIPLICATION, MODULO, NUMBER, POWER, STRING, KEYWORD, END, VARIABLE;
}
