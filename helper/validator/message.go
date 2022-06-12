package validator

import "fmt"

func GetMessage(field string, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprint("The ", field, " field is required")
	case "email":
		return fmt.Sprint("The ", field, " field must be a valid email address.")
	default:
		return tag
	}
}
