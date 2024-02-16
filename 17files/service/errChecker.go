package check

func IsErr(err error) {
	if err != nil {
		panic(err)
	}
}
