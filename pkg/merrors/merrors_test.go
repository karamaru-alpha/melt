package merrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	err := errors.New("new error")
	err = Wrapf(err, InvalidArgument, "wrapf the error")
	err = Stack(err)
	err = Wrap(err, Internal)
	fmt.Printf("%+v\n%v\n", err, err.Error())
	//	wrapf the error:
	//		github.com/karamaru-alpha/melt/pkg/merrors.TestWrap
	//			/melt/pkg/merrors/merrors_test.go:13
	//		- github.com/karamaru-alpha/melt/pkg/merrors.TestWrap
	//			/melt/pkg/merrors/merrors_test.go:12
	//		- wrapf the error:
	//			github.com/karamaru-alpha/melt/pkg/merrors.TestWrap
	//			/melt/pkg/merrors/merrors_test.go:11
	//		- new error
	//  error: code = Internal, message = wrapf the error
	t.Skip("出力の見た目確認用")
}
