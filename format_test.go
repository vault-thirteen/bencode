package bencode

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_format(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(FileSectionAnnounce, "announce")
	aTest.MustBeEqual(FileSectionAnnounceList, "announce-list")
	aTest.MustBeEqual(FileSectionCreationDate, "creation date")
	aTest.MustBeEqual(FileSectionComment, "comment")
	aTest.MustBeEqual(FileSectionCreatedBy, "created by")
	aTest.MustBeEqual(FileSectionEncoding, "encoding")
	aTest.MustBeEqual(FileSectionInfo, "info")
}
