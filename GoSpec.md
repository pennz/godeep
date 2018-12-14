
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