package errno

import (
	"errors"
	"fmt"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/base"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
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
	Success                = NewErrNo(int64(base.ErrCode_SuccessCode), "Success")
	ServiceErr             = NewErrNo(int64(base.ErrCode_ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int64(base.ErrCode_ParamErrCode), "Wrong Parameter has been given")
	AuthorizationFailedErr = NewErrNo(int64(base.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
	UserNotFoundErr        = NewErrNo(int64(base.ErrCode_UserNotFoundErrCode), "User not found")
	UserCannotLoginErr     = NewErrNo(int64(base.ErrCode_UserCannotLoginErrCode), "User cannot login")
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
