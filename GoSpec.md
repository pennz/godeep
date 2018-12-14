# Introduction

- **general-purpose**
- **strongly typed**
- **garbage-collected**
- **concurrent**

# Notation
Extended Backus-Naur Form (EBNF):

```
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operatiors, in increasing precedence:
```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

- Lower-case production names are sued to identify lexical tokens,i.e., ```production_name```.
- Non-terminals are in CamelCase.
- Lexical tokens are enclosed in quotes, "" or ``.
# Source code representation
UTF-8
...

## Characters
```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point classified as "Letter" */ .
unicode_digit  = /* a Unicode code point classified as "Number, decimal digit" */ .
```

## Letters and digits
```
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```

# Lexcial elements
## Comments
- //
- /* xxxxx */

## Tokens
1. identifiers
2. keywords
3. operators and punctuation
4. literals

A newline or end of file may trigger the insertion of a semicolon.

## Semicolons
The formal grammar uses semicolons ";" as terminators in a number of productions. 

Go programs may omit most of these semicolons using the following two rules: 
1. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is
    - an identifier
    - an integer, floating-point, imaginary, rune, or string literal
    - one of the keywords break, continue, fallthrough, or return
    - one of the operators and punctuation ++, --, ), ], or }
2. To allow complex statements to occupy a single line, a semicolon may be omitted before a closing ")" or "}".

## Identifiers
Identifiers name program entities such as variables and **types**.
```
identifier = letter { letter | unicode_digit }
```
Here pay attention to the EBNF: ```Repetition  = "{" Expression "}" .```
```
a
_x9
ThisVariableIsExported
αβ
```

## Keywords
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```
## Operators and puncuation
```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

## Integer literals
```
int_lit     = decimal_lit | octal_lit | hex_lit .
decimal_lit = ( "1" … "9" ) { decimal_digit } .
octal_lit   = "0" { octal_digit } .
hex_lit     = "0" ( "x" | "X" ) hex_digit { hex_digit } .
```
```
42
0600
0xBadFace
170141183460469231731687303715884105727
```

## Floating-point literals
```
float_lit = decimals "." [ decimals ] [ exponent ] |
            decimals exponent |
            "." decimals [ exponent ] .
decimals  = decimal_digit { decimal_digit } .
exponent  = ( "e" | "E" ) [ "+" | "-" ] decimals .
```
```
0.
72.40
072.40  // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
```

## Imaginary literals
```
imaginary_lit = (decimals | float_lit) "i" .
```

## Rune literals
**rune constant**
```
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
```
```
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // rune literal containing single quote character
'aa'         // illegal: too many characters
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
```

## String literals
```
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```
```
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:
```
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

# Constants
- boolean constants
- rune constants
- integer constants
- floating-point constants
- complex constants
- string constants

*numeric constants*
- rune constants
- integer constants
- floating-point constants
- complex constants

A constant value is represented by 
- a rune, integer, floating-point, imaginary, or string literal, 
- an identifier denoting a constant, 
- a constant expression, 
- a conversion with a result that is a constant, or 
- the result value of some built-in functions such as unsafe.Sizeof applied to any value, cap or len applied to some expressions, real and imag applied to a complex constant and complex applied to numeric constants. 
- The boolean truth values are represented by the predeclared constants ```true``` and ```false```

Numeric constants represent exact values of arbitrary precision and do not overflow.

Constants may be typed or untyped. 
untyped:
- literal constants
- true, false, iota
- certain constant expressions containing only untyped constant operands 

given type:
- explicitly
    - a constant declaration or conversion
    - used in a variable declaration or an assignment 
    - as an operand in an expression
- implicitly
    - An untyped constant has a default type which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a short variable declaration such as i := 0

# Variables
A variable is a **storage location** for holding a **value**.
The set of permissible values is determined by the variable's **type**.

- A variable declaration reserves storage for a **named** variable
- the signature of a function declaration or function literal reserves storage for funtion parameters and results.
- Calling the built-in function new or taking the address of a composite literal allocates storage for a variable at run time. Such an anonymous variable is referred to via a (possibly implicit) pointer indirection. 

Structured variables of array, slice, and struct types have elements and fields that may be addressed individually. Each such element acts like a variable. 

The static type (or just type) of a variable is 
- the type given in its declaration, 
- the type provided in the new call or composite literal, or 
- the type of an element of a structured variable

Variables of interface type also have a distinct *dynamic type*, which is the concrete type of the value assigned to the variable at run time (unless the value is the predeclared identifier nil, which has no type). The dynamic type may vary during execution but values stored in interface variables are always assignable to the static type of the variable. ??
```
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
```
A variable's value is **retrieved by referring** to the variable in an expression; it is the most recent value assigned to the variable. If a variable has not yet been assigned a value, its value is the zero value for its type. 

# Types
A type determines **a set of values** together with **operations** and **methods** specific to those values.
```
Type      = TypeName | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
	    SliceType | MapType | ChannelType .
```
The language predeclares certain type names. Others are introduced with type declarations. Composite types—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals. 

Each type T has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself. Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration. 

```
type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)
```
The underlying type of string, A1, A2, B1, and B2 is string. The underlying type of []B1, B3, and B4 is []B1. 

## Method sets
A type may have a *method set* associated with it.
The method set of an **interface type** is its interface. 
The method set of any other type T consists of all methods declared with receiver type T. 
The method set of the corresponding **pointer type \*T** is the set of all methods declared with receiver *T or T.

## Boolean types
```bool```, ```true``` and ```false```

## Numeric types
```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.

There is also a set of predeclared numeric types with implementation-specific sizes:

```
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

Conversions are required when different numeric types are mixed in an expression or assignment. For instance, int32 and int are **not** the same type even though they may have the same size on a particular architecture. 

## String types
A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. Strings are **immutable**.

It is illegal to take the address of such an element; if s[i] is the i'th byte of a string, &s[i] is invalid. 

## Array types
```
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```
The length is part of the array's type; it must evaluate to a **non-negative constant representable** by a value of type **int**. 

```
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

## Slice type
```
SliceType = "[" "]" ElementType.
```
A slice is a descriptor for a contiguous segment of an *underlying array* and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The value of an uninitialized slice is nil.

The array underlying a slice may extend past the end of the slice. The capacity is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by slicing a new one from the original slice. The capacity of a slice a can be discovered using the built-in function cap(a).

A new, initialized slice value for a given element type T is made using the built-in function make, which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with make always allocates a new, hidden array to which the returned slice value refers. That is, executing
```
make([]T, length, capacity)
```
produces the same slice as allocating an array and slicing it, so these two expressions are equivalent:
```
make([]int, 50, 100)
new([100]int)[0:50]
```

## Struct types
```
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName .
Tag           = string_lit .
```
```
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

**promoted**
<br>
...TODO: complete it...

## Pointer types
```
PointerType = "*" BaseType .
BaseType    = Type .
```
```
*Point
*[4]int
```

## Function types
A function type denotes the set of **all functions with the same parameter and result types**. The value of an uninitialized variable of function type is nil. 
```
FunctionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```
```
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```
## Interface types
An interface type specifies a **method set** called its *interface*.

A **variable** of interface type can store a value of **any type** with a method set that is any **superset** of the interface. 
```
InterfaceType      = "interface" "{" { MethodSpec ";" } "}" .
MethodSpec         = MethodName Signature | InterfaceTypeName .
MethodName         = identifier .
InterfaceTypeName  = TypeName .
```
As with all method sets, in an interface type, each method must have a unique non-blank name.
```
// A simple File interface
interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
	Close()
}
```

```
func (p T) Read(b Buffer) bool { return … }
func (p T) Write(b Buffer) bool { return … }
func (p T) Close() { … }
```
More than one type may implement an interface. For instance, if two types S1 and S2 have the method set (where T stands for either S1 or S2) then the File interface is implemented by both S1 and S2, regardless of what other methods S1 and S2 may have or share.

A type implements any interface comprising any subset of its methods and may therefore implement several distinct interfaces. For instance, all types implement the empty interface: 
```
interface{}
```

*embedding* interface E in T, it adds all (exported and non-exported) methods of E to the interface T. 

## Map types
```
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
ElementType = Type .
```
```
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

```
make(map[string]int)
make(map[string]int, 100)
```

## Channel types
```
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```
A channel may be constrained only to send or only to receive by conversion or assignment. 
```
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```
A new, initialized channel value can be made using the built-in function make, which takes the channel type and an optional capacity as arguments:
```
make(chan int, 100)
```

# Properties of types and values

# Blocks
```
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```
implicit blocks:

1. The universe block encompasses all Go source text.
2. Each package has a package block containing all Go source text for that package.
3. Each file has a file block containing all Go source text in that file.
4. Each "if", "for", and "switch" statement is considered to be in its own implicit block.
5. Each clause in a "switch" or "select" statement acts as an implicit block.`

Blocks nest and influence scoping.

# Declarations and scope
...
## Variable declaration
A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value. 
```
VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
```
```go
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
```
If a list of expressions is given, the variables are initialized with the expressions following the rules for assignments. Otherwise, each variable is initialized to its zero value.

If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment. If that value is an untyped constant, it is first converted to its default type; if it is an untyped boolean value, it is first converted to type bool. The predeclared value nil cannot be used to initialize a variable with no explicit type. 
```go
var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal
```

## Short variable declarations
```
ShortVarDecl = IdentifierList ":=" ExpressionList .
```
It is shorthand for a regular variable declaration with initializer expressions but no types:
```
"var" IdentifierList = ExpressionList .
```
Unlike regular variable declarations, a short variable declaration may **redeclare** variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the **same type**, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original.
```go
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
a, a := 1, 2                              // illegal: double declaration of a or no new variable if a was declared elsewhere
```
Short variable declarations may appear **only** inside functions. In some contexts such as the initializers for "if", "for", or "switch" statements, they can be used to declare **local temporary variables**. 

# Expressions
An expression specifies the **computation of a value** by applying operators and functions to operands. 
## Operands
Operands denote the **elementary values** in an expression. An operand may be 
- a literal, 
- a (possibly qualified) non-blank identifier denoting a constant, 
- variable, or 
- function, or 
- a parenthesized expression. 
```
Operand     = Literal | OperandName | "(" Expression ")" .
Literal     = BasicLit | CompositeLit | FunctionLit .
BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
OperandName = identifier | QualifiedIdent.
```

## Qualified identifiers
```
QualifiedIdent = PackageName "." identifier .
```
A qualified identifier accesses an identifier in a different package, which must be imported. The identifier must be exported and declared in the package block of that package.
```
math.Sin	// denotes the Sin function in package math
```
## Composite literals
Composite literals construct values for structs, arrays, slices, and maps and **create a new value each time they are evaluated**. They consist of the type of the literal followed by a brace-bound list of elements. Each element may optionally be preceded by a corresponding key. 
```
CompositeLit  = LiteralType LiteralValue .
LiteralType   = StructType | ArrayType | "[" "..." "]" ElementType |
                SliceType | MapType | TypeName .
LiteralValue  = "{" [ ElementList [ "," ] ] "}" .
ElementList   = KeyedElement { "," KeyedElement } .
KeyedElement  = [ Key ":" ] Element .
Key           = FieldName | Expression | LiteralValue .
FieldName     = identifier .
Element       = Expression | LiteralValue .
```
```go
// list of prime numbers
primes := []int{2, 3, 5, 7, 9, 2147483647}

// vowels[ch] is true if ch is a vowel
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
```
## Function literals
anonymous function
```
FunctionLit = "func" Signature FunctionBody .
```
```go
func(a, b int, z float64) bool { return a*b < int(z) }
```
A function literal can be assigned to a variable or invoked directly.
```go
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

## Primary expressions
Primary expressions are the operands for unary and binary expressions.
```
PrimaryExpr =
	Operand |
	Conversion |
	MethodExpr |
	PrimaryExpr Selector |
	PrimaryExpr Index |
	PrimaryExpr Slice |
	PrimaryExpr TypeAssertion |
	PrimaryExpr Arguments .

Selector       = "." identifier .
Index          = "[" Expression "]" .
Slice          = "[" [ Expression ] ":" [ Expression ] "]" |
                 "[" [ Expression ] ":" Expression ":" Expression "]" .
TypeAssertion  = "." "(" Type ")" .
Arguments      = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
```
```go
x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()
```
## Selectors
`x.f`

A selector f may denote a field or method f of a type T (the type of x), or it may refer to a field or method f of **a nested embedded field** of T. 

**depth**: the number of embedded fields traversed to reach f 

TODO: read the rules to find the selectors

## Method expressions
If M is in the method set of type T, T.M is a function that is callable as a **regular function** with the same arguments as M prefixed by an additional argument that is the receiver of the method.
```
MethodExpr    = ReceiverType "." MethodName .
ReceiverType  = Type .
```
 For a method with a value receiver, one can **derive** a function with an explicit pointer receiver, so
```go
(*T).Mv
```
yields a function value representing Mv with signature
```go
func(tv *T, a int) int
```
Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call. 

The final case, a value-receiver function for a pointer-receiver method, is illegal because **pointer-receiver methods are not in the method set of the value type**. 

> pointer type is more versatile

## Method values
If the expression x has static type T and M is in the method set of type T, x.M is called a _method value_. 
> here is expression `x`, for method expression, it is type `T`

As with selectors, a reference to a **non-interface method with a value receiver** using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`. 

As with method calls, a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`. 
## Index expressions

An index expression on a map a of type `map[K]V` used in an assignment or initialization of the special form
```go
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
```
## Slice expressions
### Simple slice expressions
```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```
### Full slice expressions
```go
a[low : high : max]
```

## Type assertions
For an expression x of **interface** type and a type T, the primary expression
`x.(T)`
asserts that x is not nil and that the value stored in x is of type T. The notation x.(T) is called a type assertion. 

More precisely, if T is not an interface type, x.(T) asserts that the dynamic type of x is identical to the type T. In this case, T must implement the (interface) type of x; otherwise the type assertion is invalid since it is not possible for x to store a value of type T. If T is an interface type, x.(T) asserts that the dynamic type of x implements the interface T.

> ? how is the value saved? default value int? 
> Yes, the value saved is type int. and as you can see. x is interface type. there is the value type underhood.

If the type assertion holds, the value of the expression is the value stored in x and its type is T. If the type assertion is false, a run-time panic occurs. In other words, even though the dynamic type of x is known only at run time, the type of x.(T) is known to be T in a correct program. 

```go
var x interface{} = 7          // x has dynamic type int and value 7
i := x.(int)                   // i has type int and value 7

type I interface { m() }

func f(y I) {
	s := y.(string)        // illegal: string does not implement I (missing method m)
	r := y.(io.Reader)     // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
	…
}
```
## Calls
## Passing arguments to ... parameters
## Operators
## Arithmetic operators
## Comparison operators
## Logical operators
## Address operators
## Receive operator
## Conversions
## Constant expressions
Constant expressions may contain only constant operands and are evaluated at **compile time**. 
**untyped**
```go
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
```

The values of typed constants must always be accurately representable by values of the constant type. The following constant expressions are illegal:
```go
uint(-1)     // -1 cannot be represented as a uint
int(3.14)    // 3.14 cannot be represented as an int
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
Four * 100   // product 400 cannot be represented as an int8 (type of Four)
```
## Order of evaluation

# Statements
## Assignments
```
Assignment = ExpressionList assign_op ExpressionList .

assign_op = [ add_op | mul_op ] "=" .
```

The left-hand side operand must be 
- **addressable**, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. 
- a map index expression
- or (for = assignments only) the blank identifier.

```go
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch
```

The assignment proceeds in two phases. 
1. First, the operands of index expressions and pointer indirections (including implicit pointer indirections in selectors) on the left and the expressions on the right are all evaluated in the usual order. 
2. Second, the assignments are carried out in left-to-right order. 

```go
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x == []int{3, 5, 3}
```
In assignments, each value must be assignable to the type of the operand to which it is assigned, with the following special cases:

1.    Any typed value may be assigned to the blank identifier.
2.    If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first **converted** to its **default type**.
3.    If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first converted to type bool.

# Built-in functions
Built-in functions are predeclared. They are called like any other function but some of them accept a type instead of an expression as the first argument. 

...
## Making slices, maps and channels
The built-in function make takes a type T, which must be a slice, map or channel type, optionally followed by a type-specific list of expressions. It returns a value of type `T` (not `*T`). The memory is initialized as described in the section on initial values. 
```
Call             Type T     Result

make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for approximately n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
```
```go
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for approximately 100 elements
```
