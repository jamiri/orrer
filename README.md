# Orrer

To prefer being terse over scattered error handling in Go. Used preferably in more than three or four consequent error checking.

## Api

### Types

##### Fn    func() (interface{}, error)

The building block for using the whole library. This is a wrapper function to generalize handling result or error.

##### FnArg func(interface{}) (interface{}, error)

Same as *Fn* but accepts an input as argument.

### Functions

##### GetAny(errors ...error) error

Returns an error if either of errors passed to it is not nil. Otherwise the return value would be nil.

###### Use case:

There are function calls that either return error or update a pointer passed to them the usual Go way to deal with them
is:

```go
if err := fooFunc(&fooStruct); err != nil {
    return err
}

if err := barFunc(&barStruct); err != nil {
    return err
}

if err := bazFun(&someOtherStructAgain); err != nil {
    return err
}

return nil
```

We can summarize them as:

```go
return GetAny(fooFunc(&foorStruct), barFunc(&barStruct), bazFun(&someOtherStructAgain))
```

##### <a name="GetValsOrError"></a>GetValsOrError(fns ...Fn) ([]interface{}, error)

Runs a series of functions; in case any of functions return an error, it will be returned, otherwise results is returned
in an array.

###### Use case:

There are function calls that either return error and nil for a value or return a value with a nil error:

```go
if v1, err := fooFunc(fooVal); err != nil {
    return err
}

if v2, err := barFunc(barVal); err != nil {
    return err
}

if v3, err := bazFun(bazVal); err != nil {
    return err
}

return nil
```

We can summarize (really?) them as:

```go
vals, err: = GetValsOrError(
    func () (interface{}, error) {
        return sayHi("John")
    },
    func () (interface{}, error) {
        return sayBy("John")
    },
    func () (interface{}, error) {
        return sayError("John")
    }
)
println(err) // The error returned by sayError is printed

vals, err: = GetValsOrError(
    func () (interface{}, error) {
        return sayHi("John")
    },
    func () (interface{}, error) {
        return sayBy("John")
    }
)
println(err) // nil
println(vals[0].(string)) // Hi John
```

##### GoGetValsOrError(fns ...Fn) ([]interface{}, error)

Same as [GetValsOrError](#GetValsOrError), just runs the functions in go routines and waits till either all results are ready in
the results array, or an error is returned from any functions.

##### GetSeriesOrError(kickOff interface{}, fns ...FnArg) (interface{}, error)

Runs a series of functions passing the returning value of first one to the second and so on

###### Use case:

When you want to run a series of functions passing the first output as second input and so on. If any errors occurs then
that error is returned instead.

```go
// We have these functions for demonstration------------
func AddOne(inp int) (int, error) {
	return inp + 1, nil
}

func AddTwo(inp int) (int, error) {
	return inp + 2, nil
}

func MulBy3(inp int) (int, error) {
	return inp * 3, nil
}
//------------------------------------------------------

v1, err := AddOne(1)
if err != nil {
	return nil
}
v2, err := MulBy3(v1)
if err != nil {
    return nil
}
v3, err := AddTwo(v2)
if err != nil {
    return nil
}

println(v3) // 8
```

We can summarize them as:

```go
if v, err := GetSeriesOrError(
    1,
    func(i interface{}) (interface{}, error) {
        return AddOne(i.(int))
    },
    func(i interface{}) (interface{}, error) {
        return MulBy3(i.(int))
    },
    func(i interface{}) (interface{}, error) {
        return AddTwo(i.(int))
    },
); err != nil {
	println(v) // 8
}
```

The usefulness of package is more apparent when dealing with lots of errors to check. 

For getting more familiar with the usage, take a look at the tests.