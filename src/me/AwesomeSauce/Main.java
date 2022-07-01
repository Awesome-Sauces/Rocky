package me.AwesomeSauce;

import me.AwesomeSauce.Lexer.LineProfiler;
import me.AwesomeSauce.Lexer.Token;
import me.AwesomeSauce.Lexer.Type;
import me.AwesomeSauce.utils.FileDigest;

import java.util.ArrayList;
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
        lTokens.add(";");
        //lTokens.add("\"");


	// write your code here

        /*
        if(args.length < 1) {
            System.out.println("Usage: java Lexer \"((some Scheme) (code to) lex)\".");
            return;
        }

         */

        ArrayList<String> list = null;

        try{
            list = new FileDigest("main.rocky").file;
        }catch (Exception ignored){}


        assert list != null;
        for(String string: list){


            List<Token> tokens = lex(string);

            tokens = new LineProfiler(tokens, 0).tokens;

            for(Token t : tokens){
                System.out.println(t.t + ":" + t.c);
            }

            if(!tokens.isEmpty() && tokens.get(tokens.size()-1).t.equals(Type.END)){
                System.out.println("LINE:END:" + Math.max((list.indexOf(string)-1), 0));
            }
        }



    }
}
