package main

// pass interface into the function to return a string representation of type
func typeof(v interface{}) string {
	switch v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	default:
		return "unknown"
	}
}
