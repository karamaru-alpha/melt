package merrors

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"

	"github.com/karamaru-alpha/melt/pkg/domain/enum"
)

type MeltError struct {
	// エラーパターン
	ErrorPattern ErrorPattern
	// エラーメッセージ
	errorMessage string
	// xerrors参照用
	err error
	// xerrorsでのcallStack記録
	frame xerrors.Frame
}

// ErrorPattern エラーパターン管理
type ErrorPattern struct {
	// エラーコード
	ErrorCode enum.ErrorCode
	// HTTPステータスコード
	HTTPStatusCode int
	// gRPCステータスコード
	GRPCStatusCode codes.Code
}

var (
	// Unknown 予期せぬエラー
	Unknown = ErrorPattern{
		HTTPStatusCode: http.StatusInternalServerError,
		GRPCStatusCode: codes.Unknown,
	}
	// InvalidArgument パラメータが不正
	InvalidArgument = ErrorPattern{
		ErrorCode:      enum.ErrorCodeInvalidArgument,
		HTTPStatusCode: http.StatusBadRequest,
		GRPCStatusCode: codes.InvalidArgument,
	}
	// Internal サーバー内部エラー
	Internal = ErrorPattern{
		ErrorCode:      enum.ErrorCodeInternal,
		HTTPStatusCode: http.StatusInternalServerError,
		GRPCStatusCode: codes.Internal,
	}
	// Unauthenticated 認証必要
	Unauthenticated = ErrorPattern{
		ErrorCode:      enum.ErrorCodeUnauthenticated,
		HTTPStatusCode: http.StatusUnauthorized,
		GRPCStatusCode: codes.Unauthenticated,
	}
	// PermissionDenied アクセス拒否
	//PermissionDenied = ErrorPattern{
	//	ErrorCode:      enum.ErrorCodePermissionDenied,
	//	HTTPStatusCode: http.StatusForbidden,
	//	GRPCStatusCode: codes.PermissionDenied,
	//}
	// NotFound 見つからない
	//NotFound = ErrorPattern{
	//	ErrorCode:      enum.ErrorCodeNotFound,
	//	HTTPStatusCode: http.StatusNotFound,
	//	GRPCStatusCode: codes.NotFound,
	//}
)

// New Meltエラー生成する
//func New(errorPattern ErrorPattern) error {
//	return newError(nil, errorPattern, "")
//}

// Newf Meltエラーを生成する
func Newf(errorPattern ErrorPattern, format string, args ...interface{}) error {
	return newError(nil, errorPattern, fmt.Sprintf(format, args...))
}

// Wrap 外部のエラーをMeltエラーでラップする
func Wrap(cause error, errorPattern ErrorPattern) error {
	var message string
	var meltError *MeltError
	if errors.As(cause, &meltError) {
		message = meltError.errorMessage
	} else {
		message = cause.Error()
	}
	return newError(cause, errorPattern, message)
}

// Wrapf 外部のエラーをMeltエラーでラップする
func Wrapf(cause error, errorPattern ErrorPattern, format string, a ...interface{}) error {
	return newError(cause, errorPattern, fmt.Sprintf(format, a...))
}

// stackError stacktrace用エラー
type stackError struct {
	*MeltError
}

// Stack エラーをスタックしてフレームを明示的に積む
func Stack(err error) error {
	pattern := Unknown
	message := ""
	var meltError *MeltError
	if errors.As(err, &meltError) {
		pattern = meltError.ErrorPattern
		message = meltError.errorMessage
	}
	return &stackError{
		MeltError: &MeltError{
			ErrorPattern: pattern,
			errorMessage: message,
			err:          err,
			frame:        xerrors.Caller(1),
		},
	}
}

func newError(cause error, errorPattern ErrorPattern, errorMessage string) error {
	return &MeltError{
		ErrorPattern: errorPattern,
		errorMessage: errorMessage,
		err:          cause,
		frame:        xerrors.Caller(2),
	}
}

func (e *MeltError) Error() string {
	return fmt.Sprintf("error: code = %v, message = %s", e.ErrorPattern.ErrorCode, e.errorMessage)
}

func (e *MeltError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *MeltError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *MeltError) FormatError(p xerrors.Printer) error {
	p.Print(e.Message())
	e.frame.Format(p)
	return e.Unwrap()
}

func (e *MeltError) Message() string {
	if e == nil {
		return ""
	}
	return e.errorMessage
}

func (e *stackError) FormatError(p xerrors.Printer) error {
	e.frame.Format(p)
	return e.Unwrap()
}
