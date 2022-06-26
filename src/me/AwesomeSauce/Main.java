package me.AwesomeSauce;

import me.AwesomeSauce.Lexer.Token;

import java.util.List;

import static me.AwesomeSauce.Lexer.Lexer.*;

public class Main {

    public static void main(String[] args) {

        lTokens.add("=");
        lTokens.add("+");
        lTokens.add("-");
        lTokens.add("/");
        lTokens.add("%");
        lTokens.add("(");
        lTokens.add(")");
        lTokens.add(" ");
        lTokens.add("*");
        lTokens.add("^");
        //lTokens.add("\"");


	// write your code here

        /*
        if(args.length < 1) {
            System.out.println("Usage: java Lexer \"((some Scheme) (code to) lex)\".");
            return;
        }

         */

        String input = "\"frfrfr\"";


        List<Token> tokens = lex(input);
        for(Token t : tokens) {
            System.out.println(t);
        }

    }
}
