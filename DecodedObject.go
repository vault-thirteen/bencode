package bencode

import (
	"bytes"
)

// DecodedObject is a decoded object with some meta-data.
type DecodedObject struct {
	FilePath        string
	SourceData      []byte
	RawObject       any
	DecodeTimestamp int64
	IsSelfChecked   bool
}

// MakeSelfCheck performs a simple self-check. It encodes the decoded data and
// compares it with the source.
func (do *DecodedObject) MakeSelfCheck() (success bool) {

	// Encode the decoded data.
	var err error
	var baEncoded []byte
	baEncoded, err = NewEncoder().EncodeAnInterface(do.RawObject)
	if err != nil {
		return false
	}

	// Compare the encoded decoded data with the original data.
	var checkResult int
	checkResult = bytes.Compare(baEncoded, do.SourceData)
	if checkResult != 0 {
		return false
	}

	do.IsSelfChecked = true

	return true
}
