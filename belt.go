package orrer

// GetAny returns an error if either of errors passed to it is not nil. Otherwise the return value would be nil.
func GetAny(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}

type Fn func() (interface{}, error)

// GetValsOrError runs a series of functions; in case any of functions return an error, it will be returned, unless the
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
