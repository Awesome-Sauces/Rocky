package me.AwesomeSauce.Lexer;

import java.util.List;

public class LineProfiler {

    public List<Token> tokens;
    public int lineNumber;

    // Contains a list of tokens that the specified line contains will read them and further
    // expand upon said tokens while also building a profile and code to execute for that line
    // which will be executed by line number
    public LineProfiler(List<Token> tokens, int lineNumber){

        this.tokens = tokens;
        this.lineNumber = lineNumber;



    }

}
