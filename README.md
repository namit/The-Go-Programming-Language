# The Go Programming Language

Notes, examples, and exercises from the awesome "The Go Programming Langauge" book.

## Remember
  * Parameter vs Argument: Parameter is what's specified in function declaration / definition. Argument is the actual value that's passed to that function. Ex, "Command Line Arguments."
  * Unicode is a character set (which, for ex, says code-point 'A' is represented by number 65. In hex, that's `x0041`. That's why the Unicode for 'A' is `U+0041`. Note: Go uses the term "rune" for code-points.
  * UTF-8 is an encoding (which, for ex, saves that 65 number in 8-bit binary). UTF-8 may use 1 to 4 bytes to encode any Unicode code-points.

## Why Go?
Even if a lot of people complain about Golang, this is still the best choice for me. None of the issues they complain about are major issues for me (well, except the poor debug support and no generics). And it is the only language with the combination of things I like:
  * It's static type.
  * C-derived (I find it easier to have genuine understanding of how Go works internally compared to other languages like Ruby, Elixir, etc. And perhaps that's because it's C-derived, and I understand C).
  * Has garbage collection.
  * Generates a single binary.
  * Supports concurrency very well.
  * Is opinionated.
  * Is small/simple. Has a community that prefers to just use standard lib for most work.
  * Is designed specifically for web-services.
  * Has enough momentum that it won't become a niche language.
  * Due to the above points, this is an ideal general purpose one-language (to rule them all) that I want.

## General
  * Semicolons are automatically inserted if certain tokens precede a newline. That's why the opening brace of a function declaration or for loop or if condition needs to be on the same line. That makes it obvious that a semicolon should not be automatically inserted on this line. This is also why newline is permitted after `+` in `x + y` but not before `+`.
  * For the most part, the order of declaration of func, var, const, and type doesn't matter.
  * Go doesn't explicitly differentiate between stacks and heaps. That's an implementation detail that can change.
  * All indexing is half-open. So, for `s = "abcd"`, the slice `s[0:2]` will include the character at index 0 ('a') but not the character at index 2 ('c').
  * A struct is a type which contains named fields that are collected together and treated as a single unit.
  * Methods can be defined for any named type except for pointers and interfaces.
  * `a++` is a statement (an increment statement, which means `a = a + 1`). It is not an expression. So, `b := a++` is invalid. This also makes `++a` redundant and, hence, invalid.
  * Vim snippets I use:
    * anon (anonymous function)
    * cons (multiple constants)
    * iota (multiple iota constants)
    * def/defr (defer/defer+recover function)
    * interface
    * if
    * else
    * errn (error return)
    * errh (error handle and return)
    * json (marshaling tag)
    * for (condition)
    * fori (for i loop)
    * forr (for range)
    * func
    * ff (fmt.Printf)
    * fn (fmt.Println)
    * lf (log.Printf)
    * ln (log.Println)
    * main
    * meth (method)
    * select
    * st (struct)
    * switch
    * vars (multiple variables)

## Names
  * Names begin with a Unicode letter and can have any number of letters, digits or underscores.
  * Keywords: break, case, chan, const, continue, for, import, interface, map, package, range, return, select, struct, switch, type, var, default, defer, else, fallthrough, if, func, go, goto.
  * Built-in constants, types, and functions (These names are not reserved and can be used in your own declarations):
    * Constants: true, false, iota, nil.
    * Types: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex128, complex64, bool, byte, rune, string, error.
    * Functions: make, len, cap, new, append, copy, close, delete, complex, real, imag, panic, recover.
  * `true := "Hi"` is valid because true is not reserved. It creates a string variable `true` and sets its value as "Hi". There should be very few reasons to replace built-ins, though.

## Packages
  * Each source file begins with a package declaration indicating which package this file belongs to. Package names are always lowercase.
  * All .go files of a package reside in a single directory. Moreover, one directory can have only one package.
  * Package main is special as it defines an executable instead of a library. Every executable must contain a main package and a main function, which becomes the entry point.
  * Imported packages are added when code is linked (unlike Java, where packages are used at run-time).
  * Missing or unused imported packages are considered errors. Use goimports to automatically manage addition/removal of import statements and also run gofmt.
  * If there are more than one imported packages, then use the list form of package import rather than specifying each import separately. This can also be done automatically by using goimports. The order of imports doesn't matter (gofmt sorts them alphabetically).
  * Imported packages can be renamed (ex, `import format "fmt"` will rename the "fmt" package into "format"). The only scenario where this can be useful is if there are two similarly named packages. Otherwise, don't change package names.
  * The os package provides features to interact with the underlying operating system in a platform-independent fashion. Ex, `os.Args` provides a slice containing command line arguments for the program with the first element being the name of the program itself.
  * For top-level variables (ie, declared in package block), fields of public struct, or type methods, you can have their names start with uppercase letter (ie, Unicode class Lu) to make them public (ie, visible outside the package). In a public struct, only the fields explicitly made public will be considered public.
  * Unlike in Java, private is accessible everywhere in a package, and protected doesn't exist as there's no inheritance.
  * Every statement at package level needs to start with a keyword like var, func, etc (to make it easy for parser).

## Variables & Constants
  * Basic types: bool, string, int, int8/16/32/64, uint, uint8/16/32/64, byte (alias for uint8), rune (alias for int32), float32/64, complex64/128.
  * At the time of declaration, if a variable is not initialized with a specific value, it's initialized to its zero value.
  * Declarations:
```go
var i int // Normal declaration
var p1, p2 *complex128 // Multiple variables of same type
var name, age = "Neo", 30 // Implicit declaration and initialization of different types
p3, x := &i, 12 // Short variable declaration of different types
chan := make(chan interface{}) // Notes about new() and make() below
var ( // Just like import statements, var declarations can be put in blocks
  x = isItVisible() // Initialization through function call
  y int // Zero value 0
  z = y + 1 // Initialization through expression
)
var v1 = Vertex{1,2} // Struct literal that declares and initializes a struct type defined as "type Vertex struct { x,y int }"
var v2 = Vertex{x: 1} // y gets initializes with its zero value, 0
```
  * Short variable declaration and initialization can only be used inside functions. It's not allowed at package level because it can be tricky for parser to handle.
  * Short declaration only declares variables that aren't already declared in the same lexical block. For others, it just acts as assignment. A short declaration should have least one new variable, though.
  * If you care about the size of a number (for example if it's used in data model), use specifically sized integer or floating-point types. The natural machine size int & uint should be used when the size doesn't matter. Unlike C, you can not implicitly cast between, say, an int and int64 because that code may behave differently in different platforms.
  * Command line arguments and environment variables are available globally. Ex,
```go
os.Setenv("FOO", "1")
fmt.Println("FOO has value ", os.Getenv("FOO"))
fmt.Println(os.Args) // [./app-name a b c]
```
  * `new(Type)` function allocates memory. It takes the type as the argument, and returns a pointer to a newly allocated zero value of that type. Ex,
```go
x := new(StructA) // x has type *StructA
```
  * `make(Type, size IntegerType)` function allocates and initializes the types slice, map, or chan only. Like new, the first argument is a type. Unlike new, make’s return type is the same as the type of its argument, not a pointer to it. And the allocated value is initialized (not set to zero value like in new). The reason is that slice, map and chan are structs, and aren't usable without initialization. Them being structs is also why it doesn't make sense for make function to return pointer (unlike new). Ex,
```go
v []int = make([]int, 100) // creates slice v, a struct that has non-zero fields: pointer to an array, length, and capacity. So, v is immediately usable
p *[]int = new([]int) // *p has zero value nil, which would make p useless in its current state
```
  * Constants declarations are similar to variable declarations, but, they cannot be declared using the := (short declaration) syntax.
  * Constants can only be numbers, characters (runes), strings, or booleans. They must be defined by constant expressions, evaluate-able by the compiler. For instance, `1<<3` is a constant expression, while `math.Sin(math.Pi/4)` is not because this needs to be executed at runtime.
  * Within a constant declaration, iota represents successive untyped integer constants. It is reset to 0 whenever the keyword const appears next.
```go
const ( // iota reset to 0
  c0 = iota // c0 == 0
  c1 = iota // c1 == 1
)
const ( // iota reset to 0
  f0 = 1 << iota // f0 == 1. Useful for flag constants
  f1 = 1 << iota // f1 == 2, f2 == 4, and so on
)
const x = iota // x == 0 as iota reset to 0
```
  * Declaring a local variable and then taking its address vs declaring a pointer and using new() to create the actual variable may internally be the same implementation by the compiler. Your way of declaring a variable is a hint for the compiler, not an instruction.
  * You can assign an int variable to your named type "integer" without cast, but you can't assign a variable of type integer to int without cast.
  * There's no implicit type conversion (keep in mind that using a numeric value is not implicit conversion. Ex, `var x float64 = 2` is still valid).
  * Type conversion of an int to a string creates a single character string that has the Unicode decimal code of that int. `fmt.Println(string(65))` will print the letter "A." On the other hand, type-conversion of a string to an int is not allowed.
  * Tuple assignment: `x, y = y, x // x and y swap their values`
  * Printf formatting verbs:
    * General - %v default format, %+v default format and add field names for structs, %#v Go syntax, %T type, %% literal percent sign, %t (default for bool) boolean true / false, %p (default for chan or pointer) pointer in base 16 notation. %v uses the built-in formatter for built-in types, or calls the `String()` method for user-defined types.
    * Integer - %b base 2, %c Unicode rune value of the number (so, %c of 30340 would show 的. You can guess what %c of 65 will show), %d (default for int) base 10, %o base 8, %q single-quoted character safely escaped in Go syntax (%c of '\n' will print newline but %q will print the escaped character '\n'), %x base 16, %X same as %x but with upper-case A-F letters, %U Unicode format U+1234.
    * Floating point - %b exponent power of 2, %e scientific power of e, %E same as %e but with upper-case E, %f decimal, %F same as %f, %g (default for float) Use %f for small numbers and %e for bigger ones, %G Same as %g but with upper-case %E. %9.2f means width 9 and precision 2.
    * String & Byte - %s (default for string) string or slice, %q double-quoted string safely escaped in Go syntax.
    * Compound objects default formats are:
      * Custom struct: `{field0 field1}`
      * Array/slice: `[elem0 elem1]`
      * Map: `map[key1:val1 key2:val2]`
  * The printf functions take empty interfaces as arguments and inspect their types to decide formatting. For ex, %b means base 2 if argument is an integer or exponent power of 2 if float.

## Functions
  * A function declaration contains keyword func, name of function, parameter list, optional results list, and the body. Ex: `func fnName(a string, b int) (x int, y int) { /* body */ }`
  * If the parameter types are the same (as for x & y here), they can be specified together. Ex, `func fnName(a string, b int) (x, y int) { /* body */ }`
  * Return values can be named (in which case, function can have naked return calls). Named return values are treated as variables defined at start of function. Ex,
```go
func add(x, y int) (sum int) {
  sum = x + y
  return // Will return the value of sum variable
}
```
  * Use blank identifier `_` if you want any returned values to be discarded.
  * When specifying variable number of parameters (variadic) in a function, they all need to be of the same type. Ex, `func add(nums ...int) (int) { }`. However, we can specify the variables to be of empty interface type `interface{}` to be able to send different types. In this case, however, we lose the benefits of static typing.
  * Functions are first-class, which means we can pass, return and store them just like variables.
  * A high-order function is a function that inputs or outputs at least one function.
  * A closure is a struct storing a function with its environment (ie values or references to non local variables that the function has access to). Go doesn't support named nested functions, but we can create closures using anonymous nested functions. Ex,
```go
func highOrderFunction(y int) func(int) int { // This high order function takes y int as parameter and returns an anonymous function that takes an int and returns an int. In this case, the returned value is a closure (not just an anonymous function) because it needs to save the value of y in it.
  return func(x int) int {
    return x + y
  }
}
var closure1 = highOrderFunction(1) // closure1(1) = 2 because this closure has saved a value of 1 for y. TODO: Verify that this is true. That the variable closure1 is actually a closure structure. Because printing the type of this variable currently shows "= func(int) int" as if it's just an anonymous function.
var closure2 = highOrderFunction(10) // closure2(1) = 11 because this closure has saved a value of 10 for y
```
  * When defining a method on a type (like a struct), add the type information before the method name. Ex,
```go
func (s *MyStruct) pointerMethod() { } // method on pointer. Useful if struct needs to be modified, or if struct is too big to pass by value. Can't call this method with value because caller won't be expecting value to change
func (s MyStruct) valueMethod() { } // method on value. Can call this with pointer or value, but it will take value in either case. It's better to use value methods for data safety.
// In both cases, the methods are called similarly: s.pointerMethod() and s.valueMethod()
```
  * This way of defining methods on types is just syntactic sugar. In the end, it's just a function that takes the receiver (like the MyStruct s) as another argument. As the receiver is specified in methods (s, in the above example), there's no need for a this / self. Note: Methods can be defined for any user-defined types. Ex, Int defined as `type Int int` can have methods.
  * defer statement is similar to Java's finally. The call in the defer statement is executed when the function exits even if it's exiting due to a runtime panic. Call defer statement as soon as you know you're going to need it (It'll still get executed only on function's exit).
  * Deferred calls are executed in LIFO order. Think unwinding the call stack.
  * Panic `func panic(v interface{})` and recover `func recover() interface{}` are like exceptions and catch.
  * Call `defer recover()` before any code that may start a panic. Recover function outside of defer just returns nil. If panic happens, it will unwind the stack to look for defer code to execute. Any call to recover function in that deferred code will stop the panicking sequence and return the argument sent in panic. If there is no panic, recover will return nil.
```go
defer func() {
  if r := recover(); r != nil {
    fmt.Println("This prints")
  }
  fmt.Println("This prints too")
}()
panic("Don't know what to do")
fmt.Println("Never see this")
```
  * Panics should generally be used where the code doesn't know how to proceed further. Recover only if it is safe to do so. Otherwise exit program or return error to client in case of server app. Don't use panics for basic error handling.

## Loops
  * A for loop can have:
```go
for { /* statements */ } // Zero components (infinite loop)
for i < 5 { /* statements */ } // Just a condition component (while loop)
for i = 0; i < 5; i++ { /* statements */ } // All components (initialization, condition, post)
for i, v := range "abc" { /* statements */ } // Range of values (from data types like string, array, slice, etc) to iterate over. Remember that v is passed by value, not by reference.
```
* Looping through a range is significantly faster as it gives compiler a way to perform bounds check just once for the loop rather than per iteration as in `for i < 10`.
* Initialization can have variable declaration, assignment, or function call. Variables declared in loop initialization are only available in loop scope.
* The break and continue statements behave as expected in a loop. In addition, in case of nested loops, a label can be used to target any specific loop. Ex,
```go
L: // Label useful to break out of multiple loops
  for {
    for {
      break L
    }
  }
```
  * In addition to the for loops, a break can also break out of switch and select statements.

## Conditions
  * Just like the for loop, the if condition can also start with a short initialization statement before the condition:
```go
if v := getValue(); v < 5 {
  w := 2
} else {
  // v will be accessible here too. But not w
}
```
  * In switch, a case body breaks automatically unless it ends with a fallthrough statement. Just like for and if, switch can also have an initialization statement. Ex,
```go
switch x := 2; x {
case 2: // First case that's checked
  // do something
case f(): // function is called and this case body is executed if x matches the value the function returned. Note: Just like f() here, x in the switch statement could also be a function like x(). In that case, the function is executed and the returned value is compared to cases.
  // do something
default: // default can be placed before other cases. It's evaluated ONLY if no case matches (including cases specified after the default)
  // do something else
}
```
* A switch with no condition is like "switch true". It can be used at times when we're not iterating through the values of any specific variable. Think of it as a cleaner way to list long if-else statements. Ex,
```go
x, y := 2, 3
switch {
case x == 2:
  fmt.Printf("x is 2")
case y == 3:
  fmt.Printf("y is 3") // This won't get executed because x == 2 case was matched already
case someFunc()
  fmt.Printf("someFunc() returned true")
default:
  fmt.Printf("Default")
}
```

## Interfaces
  * Interfaces support static duck typing. Duck typing means relying on implicit method implementation rather than explicit declaration of interface implementation. Static means that this check is still done at compile time rather than at runtime (unlike in most scripting languages like Ruby).
  * To summarize, in Go, any type that implements the methods from an interface implicitly implements that interface.
  * For interface inheritance, use interface composition. In the below example, Point inherits from Printer through interface composition. So, to implement the Point interface, the type will need to implement all three methods Print(), X(), and Y():
```go
type Point interface {
  Printer
  X() float64
  Y() float64
}
type Printer interface {
  Print()
}
```
  * An empty interface type `interface{}` means "any type that implements at least no methods." That just means every type (basic types, structs, etc). `interface{}` is similar to Java's Object (although Object doesn't apply to primitive types) and to some extent, C's `void*` (although `void*` can point to even garbage).
  * You can store values to empty interface types, but obviously can't call any methods on them.
  * By convention, a function that creates a new instance of an interface If1 is called NewIf1.
  * The error interface is declared as `type error interface { Error() string }`. The errors package's `errors.New(text string) error` function uses a trivial implementation of error interface called errorString.
  * The most basic and commonly used implementation of the error interface is error package's unexported errorString type which just saves an error string and returns it in the `Error()` method implementation.
  * The errors package's New function returns the errorString type: `func New(text string) error { return &errorString{text} }`.

## Pointers
  * Pointers are completely distinct from integers (unlike in C). No arithmetic allowed on pointers.
  * To convert Go pointers into C-like arbitrary pointers or to perform pointer arithmetic, use the unsafe package. However, other than for interfacing with other languages, there should be no need to use this package.
  * If you send a pointer to function F1's local variable into another function F2, then that local variable is taken from the stack and put to heap so that it can be accessible outside of F1. Keep in mind that stack / heap is implementation detail that end-user shouldn't worry about or rely upon.

## Slices, Arrays, and Lists
  * As there is no pointer arithmetic, pointers and arrays in Go are distinct types (unlike in C)
  * The size of an array is part of its type. The type of a 10 int array is different from a 5 int array. A function with a 5 int array parameter can't accept a 10 int array. Use slices instead of Arrays if you want size flexibility (which is usually the case).
```go
var arr [4]int // Declaration of array
x, y, z := [...]int{1,2}, [...]int{1,2,3}, [...]int{4,5,6} // Implicit declaration and initialization of multiple arrays.
x = y // will throw an error because the array types for x & y aren't same.
z = y // will work because the types are same.
```
* To construct values for structs, arrays, slices (and also maps), you use composite literals. They consist of the type of the literal followed by a brace-bound list of elements. Ex,
```go
s := []int{1,2} // Slice
a := [...]int{1,2} // Array
p := Point{5,5} // Struct Point with x and y fields
```
  * Slice is a view/window of an array. Internally, it's a struct with fields for size, capacity, and pointer to the underlying array. So, they are cheaper to create and take fixed memory and time (unlike arrays). Slices are used in function parameters to allow arrays of any size to be passed.
```go
var aSlice []int{1,2} // Will create a slice and an underlying array
firstHalfSlice, secondHalfSlice, fullSlice := array1[:50], array1[50:], array1[:] // Say, array1 is a [100]int array
fmt.Printf("%T vs %T",fullSlice, array1) // will show "[]int vs [100]int"
```
  * Slices are immutable. Resizing a slice means creating a new slice (or, even a new bigger array if the older array isn't big enough).
  * The append function adds the next element in the underlying array of the slice and returns a new slice that includes the next element
```go
s0 := make([]int, 2, 3) // [0,0] of array [0,0,0]
s1 := append(s0, 2) // [0,0,2] of array [0,0,2]
s2 := append(s0, 3) // [0,0,3] of array [0,0,3] Note: s1 is now [0,0,3] because the underlying array for s0,s1,s2 is the same
s3 := append(s0, 4, 4) // [0,0,4,4] of a new array [0,0,4,4] because old array wasn't big enough. s0,s1,s2 still use the same old array
```
  * Slices of slices:
```go
s := [][]int{[]int{1, 2}, []int{3, 4}} // [ [1, 2], [3, 4] ]
```
  * To create a list:
```go
l := list.New()
l.PushBack(32)
l.PushBack("str")
```
  * Lists are not type-specific. They use empty interface type for stored objects. The reason is that lists are in the standard library, not in the core language. And, you can only pass types as arguments to built-in functions like new and make, not to library functions.
  * As slices are more efficient and type-safer than lists, there's not much need for using lists.

## Strings
  * String literals can have two forms:
    * Interpreted: Runes between double quotes `"`. All characters except newline and unescaped `"` are allowed inside.
    * Raw: Runes between backticks \`. All characters except backtick are allowed inside. Backslashes and newlines have no special meaning.
  * Strings are immutable. However, the comparison operator works on strings' values. Ex,
```go
s1 := "A string"
s2 :=  "A " + "string"
s1 == s2  // will return true although &s1 != &s2.
```
  * Strings can be treated as slices of bytes. However, keep in mind that a Unicode rune can be upto 4 bytes (unlike in C where strings are arrays of characters and one byte stores one character). Iterating a string like "a的c":
```go
for i := 0 ; i<len(str) ; i++ { // Iterating over bytes
  fmt.Printf("%d ", str[i])
} // we get 97 231 154 132 99. Only 97 and 99 are correct Unicode decimal codes. The middle rune took 3 bytes.
for _,v := range str { // Iterate over runes using range
  fmt.Printf("%d ", v)
} // we get 97 30340 99. All three Unicode decimal codes are correct. This is because the range iteration on a string is done on runes, not bytes.
```
  * Regex compilation is expensive. So, they should be compiled once and reused everywhere. Use MustCompile function if you want program to panic on error in compiling regex rather than just returning err as it would in Compile function.
  * Regex compilation (which happens at runtime) should not be confused with compilation of program.
  * Use raw string literals with backtick \` instead of double quotes `"` to specify regex to avoid having to escape the backslashes. Ex,
```go
regexp.MustCompile(`\..*`) // is easier than
regexp.MustCompile("\\..*")
```

## Maps
  * To create a map with int keys and string values:
```go
a := make(map[int] string)
```
  * Any type that can be compared using == can be the key. Note: Equality in structs means "All fields are equal in both structs." Similarly, arrays are equal if all elements are equal.
  * The order of iteration (ex, using range) for a map is not specified and is intentionally randomized to discourage users from expecting a particular order.
  * Using `interface{}` as key type will make maps more flexible, but obviously lose the power of type-checking. So, it's better to not do so.
  * map is a reference to the data structure created by make. So, when a map is passed to a function, the function gets a copy of the reference (pointing to the same structure). So, the changes made to a map in a function will reflect even outside the function.
  * The second argument of the make function can be used to specify the initial capacity of the map. However, there's no need to do so because the overhead of not specifying capacity is small.
  * If we ask for value of a key that doesn't exist in the map, we will get the zero value of the value.
  * To create a Set (unordered unique values), we can simply have a map with keys as the set elements and values as bool true. Keys that don't exist will return the zero value of bool, ie, false.

## Goroutines
  * Goroutine is a function call that completes asynchronously. This function may be run in a new OS thread or multiplexed on the same thread. Goroutines are just green threads (Green threads are usually managed by a VM instead of OS, but they can also be managed by a runtime library, as is the case with Go).
```go
go someFunc() // Call someFunc in a goroutine
```
  * Think about concurrency / goroutines on the order of 100K.
  * Channels are used for communication (passing values of specified types) between goroutines. Don't communicate by sharing; share by communicating.
  * Think of channels as pipes, but handling known types rather than raw stream of data.
  * When one goroutine sends or receives on a channel, it blocks until another goroutine does the corresponding receive or send operation.
  * When two goroutines touch the same memory, there's no guarantee about which one touches it first.
  * Memory access within a goroutine can be reordered as long as it doesn't change the outcome of the routine. Ex,
```go
a = b
use(b)
// can be reordered by complier as:
use(b)
a = b
```
  * Like threads in other languages, functions called through goroutines can access local variables of the function in which it's called. The local variables are passed by value to those functions.
  * However, in case of closures, they are passed by reference. Moreover, because there's no guarantee regarding the order of execution of the goroutines, if a variable is modified after calling a closure through goroutine, there's no guarantee what value (before or after modification) of the variable will be in the closure. In such cases, it's better to pass that variable as an argument to the closure.
  * If a program exits before goroutines complete, then the goroutines are terminated.
  * There is no way to specify priorities of goroutines. However, we can yield in the low priority goroutines frequently by calling `runtime.Gosched()`. goroutines yield on channel send/receive (if they will block), syscalls (including file/network read/writes), memory allocation, `time.Sleep()`, and `runtime.Gosched()`.
  * The sync package provides simple mutex locks that can be held by just one goroutine at a time. Other goroutines wanting to lock them will wait for their turn:
```go
var lock sync.Mutex
lock.Lock()
// Do whatever
lock.Unlock() // Unlocked is the zero value of a mutex
```
  * Go mutexes are really binary semaphores (not traditional mutexes). Meaning, they have no owner and they can't be used recursively. Calling Lock on a mutex in the goroutine that has already locked it will cause deadlock. Using mutexes is not idiomatic Go. Use concurrency design patterns instead.

## HTTP HandleFunc
  * After you add handlers, the HTTP request multiplexer ServeMux checks the URL of each incoming request against a list of registered patterns of these handlers and calls the handler for the pattern that most closely matches the URL.
  * Patterns may optionally begin with a host name, restricting matches to URLs on that host only.
  * To allow the server to serve multiple requests simultaneously, each incoming request is served by a handler run in a separate goroutine.
