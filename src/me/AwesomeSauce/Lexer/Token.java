package me.AwesomeSauce.Lexer;

public class Token {
    // could have column and line number fields too, for reporting errors later
    public Type t;
    public String c;

    public Token(Type t, String c){
        this.t = t;
        this.c = c;
    }

    public String toString() {
        if (t == Type.ATOM) {
            return "ATOM<" + c + ">";
        }else if (t == Type.NUMBER) {
            return "NUMBER<" + c + ">";
        }else if (t == Type.STRING) {
            return "STRING<" + c + ">";
        }
        return t.toString();
    }
}
