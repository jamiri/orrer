package orrer

import "sync"

type (
	Fn    func() (interface{}, error)
	FnArg func(interface{}) (interface{}, error)
)

// GetAny returns an error if either of errors passed to it is not nil. Otherwise the return value would be nil.
func GetAny(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetValsOrError runs a series of functions; in case any of functions return an error, it will be returned, otherwise
// results is returned in an array
func GetValsOrError(fns ...Fn) ([]interface{}, error) {
	res := make([]interface{}, len(fns))
	for idx, fn := range fns {
		val, err := fn()
		if err != nil {
			return nil, err
		}
		res[idx] = val
	}
	return res, nil
}

// GoGetValsOrError runs a series of functions concurrently; in case any of the passed functions return an error,
// it will be returned, otherwise the results is returned in an array
func GoGetValsOrError(fns ...Fn) ([]interface{}, error) {
	res := make([]interface{}, len(fns))
	errCh := make(chan error)
	doneCh := make(chan struct{})
	wg := sync.WaitGroup{}

	for idx, fn := range fns {
		wg.Add(1)
		go func(i int, toRun Fn) {
			val, err := toRun()
			if err != nil {
				errCh <- err
				return
			}
			res[i] = val
			wg.Done()
		}(idx, fn)
	}

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			return nil, err
		case <-doneCh:
			return res, nil
		}
	}
}

// GetSeriesOrError runs a series of functions passing the returning value of first one to the second and so on.
func GetSeriesOrError(kickOff interface{}, fns ...FnArg) (interface{}, error) {
	var val interface{}
	var err error
	for idx, fn := range fns {
		if idx == 0 {
			val, err = fn(kickOff)
			continue
		}
		val, err = fn(val)
		if err != nil {
			return nil, err
		}
	}
	return val, nil
}
