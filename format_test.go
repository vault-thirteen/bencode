package bencode

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_format(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(HeaderDictionary, byte('d'))
	aTest.MustBeEqual(HeaderInteger, byte('i'))
	aTest.MustBeEqual(HeaderList, byte('l'))
	aTest.MustBeEqual(HeaderStringSizeValueDelimiter, byte(':'))

	aTest.MustBeEqual(FooterCommon, byte('e'))
}
