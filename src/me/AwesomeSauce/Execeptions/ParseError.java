package me.AwesomeSauce.Execeptions;

public class ParseError extends RuntimeException {

    public ParseError() {
        super("Exception While Lexical Tokenizing");
    }

    public ParseError(String message) {
        super(message);
    }

    public ParseError(String message, Throwable cause) {
        super(message, cause);
    }
}
