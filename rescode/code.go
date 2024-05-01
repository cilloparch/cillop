package rescode

type RC struct {
	Code          uint64
	Message       string
	HttpStatus    int
	Translateable bool
}

type Extra struct {
	HttpStatus    int
	Translateable bool
}

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

func (r *RC) Error() string {
	return r.Message
}
