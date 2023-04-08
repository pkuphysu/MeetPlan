package errno

import (
	"errors"
	"fmt"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/base"
)

type ErrNo struct {
	ErrCode base.StatusCode
	ErrMsg  string
}

func (e ErrNo) Is(target error) bool {
	t, ok := target.(ErrNo)
	if !ok {
		return false
	}
	return e.ErrCode == t.ErrCode
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code base.StatusCode, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func (e ErrNo) WithError(err error) ErrNo {
	e.ErrMsg = err.Error()
	return e
}

var (
	Success                = NewErrNo(base.StatusCode_SuccessCode, "Success")
	ServiceErr             = NewErrNo(base.StatusCode_ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(base.StatusCode_ParamErrCode, "Wrong Parameter has been given")
	AuthorizationFailedErr = NewErrNo(base.StatusCode_AuthorizationFailedErrCode, "Authorization failed")
	UserNotFoundErr        = NewErrNo(base.StatusCode_UserNotFoundErrCode, "User not found")
	UserCannotLoginErr     = NewErrNo(base.StatusCode_UserCannotLoginErrCode, "User cannot login")
	OrderNotFoundErr       = NewErrNo(base.StatusCode_OrderNotFoundErrCode, "Order not found")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}

func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return convert2BaseResp(Success)
	}

	e := ErrNo{}
	if errors.As(err, &e) {
		return convert2BaseResp(e)
	}

	s := ServiceErr.WithMessage(err.Error())
	return convert2BaseResp(s)
}

func convert2BaseResp(err ErrNo) *base.BaseResp {
	resp := base.NewBaseResp()
	resp.SetStatusCode(err.ErrCode)
	resp.SetMessage(err.ErrMsg)
	return resp
}
