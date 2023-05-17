package helpers

func IsFieldNotNull(field interface{}) bool {
	switch t := field.(type) {
	case int:
		// if type is an integer
		return t != 0
	case string:
		// if type is a string
		return t != ""
	default:
		// if type is other than above
		return false
	}
}
