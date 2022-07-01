package me.AwesomeSauce.Lexer;

import java.util.List;
import java.util.ArrayList;

/*
 * Lexical analyzer for Scheme-like minilanguage:
 * (define (foo x) (bar (baz x)))
 */
public class Lexer {

    public static List<String> lTokens = new ArrayList<>();
    private static boolean onString = false;
    private static boolean lastOnString = false;

    /*
     * Given a String, and an index, get the atom starting at that index
     */
    public static String getAtom(String s, int i) {

        int j = i;
        while (j < s.length()) {
            if(!lTokens.contains(String.valueOf(s.charAt(j)))
            ) {
                j++;
            } else {
                return s.substring(i, j);
            }
        }
        return s.substring(i, j);
    }

    public static List<Token> lex(String input) {
        List<Token> result = new ArrayList<Token>();
        for(int i = 0; i < input.length(); ) {
            switch(String.valueOf(input.charAt(i))) {
                case "=":
                    result.add(new Token(Type.EQUAL, "="));
                    i++;
                    break;
                case "^":
                    result.add(new Token(Type.POWER, "^"));
                    i++;
                    break;
                case "*":
                    result.add(new Token(Type.MULTIPLICATION, "*"));
                    i++;
                    break;
                case "+":
                    result.add(new Token(Type.ADDITION, "+"));
                    i++;
                    break;
                case "-":
                    result.add(new Token(Type.SUBTRACTION, "-"));
                    i++;
                    break;
                case "/":
                    result.add(new Token(Type.DIVISION, "/"));
                    i++;
                    break;
                case "%":
                    result.add(new Token(Type.MODULO, "%"));
                    i++;
                    break;
                case "(":
                    result.add(new Token(Type.LPAREN, "("));
                    i++;
                    break;
                case ")":
                    result.add(new Token(Type.RPAREN, ")"));
                    i++;
                    break;
                case ";":
                    result.add(new Token(Type.END, ";"));
                    return result;
                default:

                    if(Character.isWhitespace(input.charAt(i))) {
                        i++;
                    }else{
                        /*
                        if(String.valueOf(input.charAt(i)).equals("\"")) {

                            if(!lastOnString) onString = !onString;
                        }

                        System.out.println(onString + ":" + input.charAt(i) + ":" + lastOnString);

                        if(onString ){
                            String atom = getAtom(input, i);
                            i += atom.length();
                            if(lastOnString) {
                                result.add(new Token(Type.STRING, atom));
                                lastOnString = false;
                                break;
                            }
                            //onString = false;
                            lastOnString = true;
                            break;
                        }else */ if(Character.isDigit(input.charAt(i))){
                            String atom = getAtom(input, i);
                            i += atom.length();
                            result.add(new Token(Type.NUMBER, atom));
                            break;
                        } else {
                            String atom = getAtom(input, i);
                            i += atom.length();
                            result.add(new Token(Type.ATOM, atom));
                            break;
                        }
                    }


                    break;
            }
        }
        return result;
    }

}