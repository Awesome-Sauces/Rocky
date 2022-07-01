package me.AwesomeSauce.Lexer;

import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class LineProfiler {

    public List<Token> tokens;
    public int lineNumber;

    // Contains a list of tokens that the specified line contains will read them and further
    // expand upon said tokens while also building a profile and code to execute for that line
    // which will be executed by line number
    public LineProfiler(List<Token> tokens, int lineNumber){

        List<Token> aTokens;

        for(Token token : tokens){
            if(token.c.contains("\"")){
                Pattern p = Pattern.compile("\"([^\"]*)\"");
                Matcher m = p.matcher(token.c);
                while (m.find()) {
                    tokens.set(tokens.indexOf(token), new Token(Type.STRING, m.group(1)));
                }

            }
        }

        this.tokens = tokens;
        this.lineNumber = lineNumber;



    }

}
