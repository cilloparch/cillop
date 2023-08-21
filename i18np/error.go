package i18np

import "fmt"

// I18nError is a struct that contains the error key and the error params
// to be used in the translation
// It implements the error interface
// It can be used to return errors from functions
type Error struct {
	// Key is the error key to be used in the translation
	// It is required
	// If it is empty, the error is not an error
	Key string

	// Params is the error params to be used in the translation
	// It is optional
	// If it is nil, the error params will be empty
	Params *P

	// Details is the error details to be used in the translation
	// It is optional
	// If it is nil, the error details will be empty
	Details interface{}
}

// IsDetails returns true if the error has details
func (e *Error) IsDetails() bool {
	return e.Details != nil
}

// P is a map of params to be used in the translation
// It is used to pass the params to the I18nError
// It is used to pass the params to the I18n.TranslateWithParams function
type P map[string]interface{}

// Error returns the error key and the error params
// It implements the error interface
func (e *Error) Error() string {
	msg := e.Key
	if e.Params != nil {
		for k, v := range *e.Params {
			msg += fmt.Sprintf(" %s: %v", k, v)
		}
	}
	return msg
}

// IsErr returns true if the error is an error
func (e *Error) IsErr() bool {
	return e.Key != ""
}

// NewError returns a new I18nError
// It is used to return errors from functions
// example:
// i18n.NewError("error.key", i18np.P{"param1": "value1"})
func NewError(key string, params ...P) *Error {
	p := &P{}
	if len(params) > 0 {
		p = &params[0]
	}
	return &Error{Key: key, Params: p}
}

// NewErrorDetails returns a new I18nError with details
//
//	It is used to return errors from functions
//
// example:
// i18n.NewErrorDetails("error.key", "details", i18np.P{"param1": "value1"})
func NewErrorDetails(key string, details interface{}, params ...P) *Error {
	p := &P{}
	if len(params) > 0 {
		p = &params[0]
	}
	return &Error{Key: key, Params: p, Details: details}
}
