package errno

import (
	"bytes"
	"errors"
	"net/http"
	"runtime"
	"strconv"
)

type Err struct {
	err        error
	errCode    int
	statusCode int

	pcs        []uintptr
	stackTrace string
}

func (e *Err) Error() string {
	return e.err.Error()
}

func (e *Err) Unwrap() error {
	return e.err
}

func (e *Err) StatusCode() int {
	return e.statusCode
}

func (e *Err) ErrCode() int {
	return e.errCode
}

func (e *Err) SetErrCode(errCode int) *Err {
	e.errCode = errCode
	return e
}

func (e *Err) StackTrace() string {
	if len(e.stackTrace) != 0 {
		return e.stackTrace
	}
	buf := bytes.NewBufferString("Stack Trace:")
	if len(e.pcs) != 0 {
		frame := runtime.CallersFrames(e.pcs)
		for {
			frame, more := frame.Next()
			if !more {
				break
			}
			buf.WriteByte('\n')
			buf.WriteString(frame.Function)
			buf.WriteByte('\n')
			buf.WriteByte('\t')
			buf.WriteString(frame.File)
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(frame.Line))
		}
	}
	if ierr, ok := e.err.(interface{ StackTrace() string }); ok {
		buf.WriteByte('\n')
		buf.WriteString(ierr.StackTrace())
	} else if ierr, ok := e.err.(*Err); ok {
		buf.WriteByte('\n')
		buf.WriteString(ierr.StackTrace())
	}
	e.stackTrace = buf.String()
	return e.stackTrace
}

func ToInternalErr(err error) *Err {
	if err == nil {
		return nil
	}
	return &Err{err: err, statusCode: http.StatusInternalServerError, pcs: callers(), errCode: -1}
}

func NewInternalErr(msg string) *Err {
	return ToInternalErr(errors.New(msg))
}

func ToValidationErr(err error) *Err {
	if err == nil {
		return nil
	}
	return &Err{err: err, statusCode: http.StatusBadRequest, pcs: callers(), errCode: -1}
}

func NewValidationErr(msg string) *Err {
	return ToValidationErr(errors.New(msg))
}

func ToPermissionErr(err error) *Err {
	if err == nil {
		return nil
	}
	return &Err{err: err, statusCode: http.StatusForbidden, pcs: callers(), errCode: -1}
}

func NewPermissionErr(msg string) *Err {
	return ToPermissionErr(errors.New(msg))
}

func ToNotFoundErr(err error) *Err {
	if err == nil {
		return nil
	}
	return &Err{err: err, statusCode: http.StatusNotFound, pcs: callers(), errCode: -1}
}

func NewNotFoundErr(msg string) *Err {
	return ToNotFoundErr(errors.New(msg))
}

func ToAuthenticationErr(err error) *Err {
	if err == nil {
		return nil
	}
	return &Err{err: err, statusCode: http.StatusUnauthorized, pcs: callers(), errCode: -1}
}

func NewAuthenticationErr(msg string) *Err {
	return ToAuthenticationErr(errors.New(msg))
}

func ToDependencyErr(err error) *Err {
	if err == nil {
		return nil
	}
	return &Err{err: err, statusCode: http.StatusFailedDependency, pcs: callers(), errCode: -1}
}

func NewDependencyErr(msg string) *Err {
	return ToDependencyErr(errors.New(msg))
}

func callers() []uintptr {
	var pcs [64]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}
