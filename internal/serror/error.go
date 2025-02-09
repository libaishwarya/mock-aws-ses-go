package serror

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// getJSONFieldName extracts the JSON field name from a struct field
func getJSONFieldName(obj interface{}, fieldName string) string {
	objType := reflect.TypeOf(obj)
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if field.Name == fieldName {
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				return strings.Split(jsonTag, ",")[0] // Remove `omitempty` or other tags
			}
			return fieldName // Fallback to struct field name if JSON tag is missing
		}
	}
	return fieldName
}

// getErrorMessage maps validation errors to user-friendly messages
func getErrorMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "numeric":
		return "This field must be a number"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "len":
		return "Invalid length"
	default:
		return "Invalid value"
	}
}

// HandleBindError formats binding errors in one line
func HandleBindError(c *gin.Context, err error, obj interface{}) {
	var validationErrors validator.ValidationErrors

	if errors, ok := err.(validator.ValidationErrors); ok {
		validationErrors = errors
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fields provided"})
		return
	}

	var errorMessages []string
	for _, e := range validationErrors {
		field := getJSONFieldName(obj, e.Field()) // Get JSON field name
		message := getErrorMessage(e.Tag())       // Custom error message
		errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", field, message))
	}

	// Convert errors to a single-line string
	errorMessageString := strings.Join(errorMessages, ", ")

	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("validation failed: %s", errorMessageString),
	})
}
