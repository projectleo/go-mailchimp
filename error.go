package mailchimp

import (
	"fmt"
)

type SubError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Type   string     `json:"type"`
	Title  string     `json:"title"`
	Status int        `json:"status"`
	Detail string     `json:"detail"`
	Errors []SubError `json:"errors"`
}

// Error ...
func (e ErrorResponse) Error() string {
	err := fmt.Sprintf("Error %d %s (%s)", e.Status, e.Title, e.Detail)
	for _, subError := range e.Errors {
		err += fmt.Sprintf("\nSubError: %s (%s)", subError.Message, subError.Field)
	}
	return err
}
