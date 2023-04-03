import json

# define a dictionary to hold variables
variables = {}

# define a dictionary to hold built-in functions
def custom_print(*args):
    print(args)

# define a dictionary to hold built-in functions
functions = {
    "print": custom_print
}

# define a function to evaluate expressions
def evaluate_expression(expression):
    if expression["Type"] == "NUMBER":
        return int(expression["Value"])
    if expression["Type"] == "STRING":
        return str(expression["Value"])
    elif expression["Type"] == "IDENTIFIER":
        if expression["Value"] in variables:
            return variables[expression["Value"]]
        elif expression["Value"] in functions:
            return functions[expression["Value"]]
        else:
            raise ValueError("Unknown variable or function: {}".format(expression["Value"]))
    else:
        raise ValueError("Unknown expression type: {}".format(expression['Type']))

# define a function to evaluate statements
def evaluate_statement(statement):
    if statement["Type"] == "TYPE":
        # create a new variable
        variable_name = statement["Value"]
        if statement["Order"] + 2 >= len(tokens):
            raise ValueError("Expected more tokens after variable declaration")
        variable_value = evaluate_expression(tokens[statement["Order"] + 2])
        variables[variable_name] = variable_value
    elif statement["Type"] == "IDENTIFIER" and tokens[statement["Order"]+1]["Type"] == "LPAREN":
        # call a function
        function_name = statement["Value"]
        args = []
        i = statement["Order"] + 2
        while tokens[i]["Type"] != "RPAREN":
            if tokens[i]["Type"] == "COMMA":
                i += 1
            else:
                args.append(evaluate_expression(tokens[i]))
                i += 1
        if function_name not in functions:
            raise ValueError("Unknown function: {}".format(function_name))
        functions[function_name](*args)
    elif statement["Type"] == "IDENTIFIER":
        # get the value of a variable
        if statement["Value"] not in variables:
            raise ValueError("Unknown variable: {}".format(statement["Value"]))
        print(variables[statement["Value"]])
    elif statement["Type"] == "ORDER":
        # do nothing for order tokens
        pass
    else:
        raise ValueError("Unknown statement type: {}".format(statement["Type"]))

# load tokens from JSON file
with open("rocky.json") as f:
    tokens = json.load(f)

# evaluate each statement
i = 0
while i < len(tokens):
    token = tokens[i]
    if token["Type"] == "FUNCTION":
        # define a new function
        function_name = tokens[i+1]["Value"]
        parameters = []
        j = i + 3
        while tokens[j]["Type"] != "RPAREN":
            if tokens[j]["Type"] == "COMMA":
                j += 1
            else:
                parameter_name = tokens[j]["Value"]
                parameter_type = tokens[j-1]["Value"]
                parameters.append((parameter_name, parameter_type))
                j += 2
        if tokens[j+1]["Type"] != "LBRACKET":
            raise ValueError("Expected LBRACKET after function declaration")
        function_body = []
        k = j + 2
        while tokens[k]["Type"] != "RBRACKET":
            function_body.append(tokens[k])
            k += 1
        functions[function_name] = (parameters, function_body)
        i = k + 1
    else:
        evaluate_statement(token)
        i += 1
