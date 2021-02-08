package utilities

// CheckError panic error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}