package tlog

import (
	"net/http"

	"github.com/getsentry/raven-go"
)

type ErrData struct {
	_error       error
	_err_msg     []string
	_stack_track *raven.Stacktrace
	_request     *http.Request
}

func NewErrData(err error, errmsg string) *ErrData {
	return &ErrData{
		_error:       err,
		_err_msg:     []string{errmsg},
		_stack_track: raven.NewStacktrace(1, 3, []string{}),
	}
}

func (errdata *ErrData) AttachRequest(request *http.Request) *ErrData {
	errdata._request = request

	return errdata
}

func (errdata *ErrData) Error() error {
	return errdata._error
}

func (errdata *ErrData) AttachErrMsg(errmsg string) *ErrData {
	errdata._err_msg = append(errdata._err_msg, errmsg)

	return errdata
}
