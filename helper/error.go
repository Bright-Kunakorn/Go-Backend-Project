package helper

func ErrorPanic(err error) error {
	if err != nil {
		panic(err)
	}
	return err
}
