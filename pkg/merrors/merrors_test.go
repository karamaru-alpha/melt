package merrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrap(t *testing.T) {
	err := errors.New("new error")
	err = Wrapf(err, InvalidArgument, "wrapf")
	err = Stack(err)
	err = Wrap(err, Internal)
	fmt.Printf("%+v\n%v\n", err, err.Error())
	t.Skip("出力の見た目確認用")
}
