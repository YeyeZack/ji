package modules

// Check Error Handling
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
