package enum

import (
	"strconv"
)

type ErrorCode int32

const (
	// ErrorCodeInvalidArgument パラメータが不正
	ErrorCodeInvalidArgument ErrorCode = iota + 1
	// ErrorCodePermissionDenied アクセス拒否
	ErrorCodePermissionDenied
	// ErrorCodeUnauthenticated 認証が必要
	ErrorCodeUnauthenticated
	// ErrorCodeNotFound 見つからない
	ErrorCodeNotFound
	// ErrorCodeInternal サーバー内部エラー
	ErrorCodeInternal
)

func (e ErrorCode) String() string {
	switch e {
	case ErrorCodeInvalidArgument:
		return "InvalidArgument"
	case ErrorCodePermissionDenied:
		return "PermissionDenied"
	case ErrorCodeUnauthenticated:
		return "Unauthenticated"
	case ErrorCodeNotFound:
		return "CodeNotFound"
	case ErrorCodeInternal:
		return "Internal"
	default:
		return strconv.FormatInt(int64(e), 10)
	}
}
