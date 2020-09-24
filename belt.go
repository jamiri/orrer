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


