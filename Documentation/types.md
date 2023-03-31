# Rocky Lang Data Types
Here I will describe all the Data Types that are/will be in Rocky. I will show how to make a variable of them, all the functions related to that type and many other things.

bool: represents a boolean value and is either true or false
int: represents integer types, floating point values, and complex types
String: represents a string value
list: represents a list in Rocky, similar to python lists.
class: represents a class value built from whatever.

## Bool use cases
Here is how to use a bool in Rocky:
```python
import stdlib as utils
import concurrency as con
import iterate as iter

func loop() bool {
    iter.loop(iter.seconds(1)){
        @Override
        func run(){
            # This will print the current milisecond the loop is on
            this.out().status();
        }
    }

    return true;
}

bool case = false;

con.run(loop());

if (case){
    utils.print("Hello World");
}
```
This will simply make a second thread that will run for 1 second and print out the status for each time it is ran.

Here I'll show you how to pass a variable through multiple functions and classes
```python
import stdlib as utils
import memory as mem

class Dog {
    # Will generate automatic getter and setters
    # Private keyword in rocky will automaticly generate getters and setters
    # for that variable. To avoid that add the closed keyword instead of
    # private keyword
    # Example: closed bool alive;
    private bool alive;

    func Dog(bool alive){
        this.alive = alive;
    }
}

# Instantize new Dog Class
Dog dog = new Dog(true);

# dog.getAlive() will return a mirror of the original bool
# If we wish to get the pointer we would have to do this:
# mem.pointer(dog.getAlive()); 
bool alive = dog.getAlive();

```
So this is how to use a bool or really any vanilla data type provided by Rocky's standard lib.
Now we will show you how to create a custom Data Type that is constructed based off of itself.

# Custom Data Type Creation
```python
import stdlib as stdlib

class ArrayList extends DataType {
    closed DataType type;
    
    # We must work with bytes to create this ArrayList
    # We must avoid using the list type to create this ArrayList type
    # That way we don't have coliding types.
    closed bytes storage;
    
    func ArrayList(DataType type){
        this.type = type;
    }
    
    func AddValue(var data){
        if (stdlib.type(data) != this.type){
            return new stdlib.Error("Cannot use %s DataType", stdlib.toString(stdlib.type(data)));
    
    # Write a simple parser to parse the byte variable.
    func parseBytes(){
        
    }
}    
