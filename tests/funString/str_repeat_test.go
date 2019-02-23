package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func StrRepeat_test(t *testing.T) {

	Convey("重复一个字符串", t, func() {

		So(funString.StrRepeat("ce", 2), ShouldEqual, "cece")

	})
}