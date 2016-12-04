package lexer

import (
	"io/ioutil"

	. "github.com/smartystreets/goconvey/convey"
)

func newFileLex(s string) *FileLex {
	tmpFile, err := ioutil.TempFile("", "")
	So(err, ShouldBeNil)

	_, err = tmpFile.Write([]byte(s))
	So(err, ShouldBeNil)

	f, err := NewFileLex(tmpFile.Name())
	So(err, ShouldBeNil)

	return f
}
