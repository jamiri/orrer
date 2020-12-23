package orrer

import (
	"fmt"
	"testing"
	"time"
)

func TestGetAny(t *testing.T) {
	msg1 := &Result{}
	msg2 := &Result{}
	msg3 := &Result{}
	err1 := hello("You", msg1)
	err2 := bye("You", msg2)

	if GetAny(err1, err2) != nil {
		t.Fail()
	}

	err3 := errrr("You", msg3)

	if GetAny(err1, err2, err3) == nil {
		t.Fail()
	}
}

type Result struct {
	Message string
}

func hello(name string, r *Result) error {
	r.Message = fmt.Sprintf("Hello %s", name)
	return nil
}

func bye(name string, r *Result) error {
	r.Message = fmt.Sprintf("bye %s", name)
	return nil
}

func errrr(name string, r *Result) error {
	return fmt.Errorf("Et errorum, ergo sum")
}

// ------------------------------------------------------------------------------------

func TestGetValsOrError(t *testing.T) {
	v, err := GetValsOrError(
		func() (interface{}, error) {
			return sayHi("John")
		},
		func() (interface{}, error) {
			return sayBye("Jona")
		},
		func() (interface{}, error) {
			return sayBye("David")
		},
		func() (interface{}, error) {
			return sayHi("Jake")
		},
		func() (interface{}, error) {
			return sayErr("Javad")
		},
	)

	if err == nil || v != nil {
		t.Fail()
	}

	v, err = GetValsOrError(
		func() (interface{}, error) {
			return sayHi("John")
		},
		func() (interface{}, error) {
			return sayBye("Jona")
		},
		func() (interface{}, error) {
			return sayBye("David")
		},
		func() (interface{}, error) {
			return sayHi("Jake")
		},
	)

	if err != nil {
		t.Fail()
	}

	if v[0].(string) != "Hello John" {
		t.Fail()
	}
}

func sayHi(name string) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}

func sayBye(name string) (string, error) {
	return fmt.Sprintf("Bye %s", name), nil
}

func sayErr(name string) (string, error) {
	return "", fmt.Errorf("Something went wrong")
}

// ------------------------------------------------------------------------------------

func TestGoGetValsOrError(t *testing.T) {

	v, err := GoGetValsOrError(
		func() (interface{}, error) {
			time.Sleep(5 * time.Millisecond)
			return sayHi("John")
		},
		func() (interface{}, error) {
			time.Sleep(4 * time.Millisecond)
			return sayBye("Jona")
		},
		func() (interface{}, error) {
			time.Sleep(3 * time.Millisecond)
			return sayBye("David")
		},
		func() (interface{}, error) {
			time.Sleep(1 * time.Millisecond)
			return sayHi("Jake")
		},
	)

	if err != nil {
		t.Fail()
	}

	if v[0].(string) != "Hello John" {
		t.Fail()
	}
}

// ------------------------------------------------------------------------------------

func TestGetSeriesOrError(t *testing.T) {
	v, err := GetSeriesOrError(
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
	)

	if err != nil {
		t.Fail()
	}

	if v != 8 {
		t.Fail()
	}
}

func AddOne(inp int) (int, error) {
	return inp + 1, nil
}

func AddTwo(inp int) (int, error) {
	return inp + 2, nil
}

func MulBy3(inp int) (int, error) {
	return inp * 3, nil
}