package rescode

// RC is a struct that contains the code, message, http status, and translateable.
// Code is the code of the error.
type RC struct {

	// Code is the code of the error.
	Code uint64

	// Message is the message of the error.
	Message string

	// HttpStatus is the http status of the error.
	HttpStatus int

	// Translateable is the flag to determine whether the message is translatable.
	Translateable bool
}

// Extra is a struct that contains the http status and translateable.
type Extra struct {
	// HttpStatus is the http status of the error.
	HttpStatus int

	// Translateable is the flag to determine whether the message is translatable.
	Translateable bool
}

// New is a function to create a new RC.
func New(code uint64, message string, extra ...Extra) *RC {
	e := Extra{
		HttpStatus:    400,
		Translateable: true,
	}
	if len(extra) > 0 {
		e = extra[0]
	}
	return &RC{
		Code:          code,
		Message:       message,
		HttpStatus:    e.HttpStatus,
		Translateable: e.Translateable,
	}
}

// JSON is a function to return the RC as a JSON.
func (r *RC) JSON(msgs ...string) map[string]interface{} {
	msg := r.Message
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return map[string]interface{}{
		"code":    r.Code,
		"message": msg,
	}
}

// Error is a function to return the message of the RC.
func (r *RC) Error() string {
	return r.Message
}
