// Code generated by hertz generator.

package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/internal/middleware"
)

func rootMw() []app.HandlerFunc {
	return nil
}

func _apiMw() []app.HandlerFunc {
	return nil
}

func _v1Mw() []app.HandlerFunc {
	return nil
}

func _listfriendlinkMw() []app.HandlerFunc {
	return nil
}

func _createfriendlinkMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _loginMw() []app.HandlerFunc {
	return nil
}

func _createmeetplanMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _deletemeetplansMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _createmeetplanandorderMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _createorderMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _updatetermdaterangeMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _gettermdaterangeMw() []app.HandlerFunc {
	return nil
}

func _createupdaterecordMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _listupdaterecordMw() []app.HandlerFunc {
	return nil
}

func _listuserMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _meetplanMw() []app.HandlerFunc {
	return nil
}

func _getmeetplanMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _meetplan0Mw() []app.HandlerFunc {
	return nil
}

func _listmeetplanMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _updatemeetplanMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _deletemeetplanMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _orderMw() []app.HandlerFunc {
	return nil
}

func _listorderMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _updateorderMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _order0Mw() []app.HandlerFunc {
	return nil
}

func _getorderMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _userMw() []app.HandlerFunc {
	return nil
}

func _getuserMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _getselfMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _createuserMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.Jwt()}
}

func _updateuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _user0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateoptionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Jwt()}
}

func _getoptionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Jwt()}
}
