package bencode

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

// DecodedObject is a decoded object with some meta-data.
type DecodedObject struct {

	// Primary Information.
	FilePath        string
	SourceData      []byte
	DecodedObject   any
	DecodeTimestamp int64

	// Secondary Information.
	IsSelfChecked bool
	BTIH          BtihData
}

// CalculateBtih calculates the BitTorrent Info Hash (BTIH) check sum.
func (do *DecodedObject) CalculateBtih() (err error) {

	// Get the 'info' section from the decoded object.
	var infoSection any
	infoSection, err = do.GetInfoSection()
	if err != nil {
		return err
	}

	// Encode the 'info' section.
	var infoSectionBA []byte
	infoSectionBA, err = NewEncoder().EncodeAnInterface(infoSection)
	if err != nil {
		return err
	}

	// Calculate the BTIH check sum.
	do.BTIH.Bytes, do.BTIH.Text = CalculateSha1(infoSectionBA)

	return nil
}

// GetInfoSection gets an 'info' section from the object.
func (do *DecodedObject) GetInfoSection() (result any, err error) {

	// Get the dictionary.
	var dictionary []DictionaryItem
	var ok bool
	dictionary, ok = do.DecodedObject.([]DictionaryItem)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	// Get the 'info' section from the decoded object.
	var dictItem DictionaryItem
	for _, dictItem = range dictionary {
		if string(dictItem.Key) == FileSectionInfo {
			return dictItem.Value, nil
		}
	}

	return nil, errors.New(ErrSectionDoesNotExist)
}

// MakeSelfCheck performs a simple self-check. It encodes the decoded data and
// compares it with the source.
func (do *DecodedObject) MakeSelfCheck() (success bool) {

	// Encode the decoded data.
	var err error
	var baEncoded []byte
	baEncoded, err = NewEncoder().EncodeAnInterface(do.DecodedObject)
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

// CalculateSha1 calculates the SHA-1 check sum and returns it as a hexadecimal
// text and byte array.
func CalculateSha1(data []byte) (resultAsBytes Sha1Sum, resultAsText string) {
	resultAsBytes = sha1.Sum(data)
	resultAsText = hex.EncodeToString(resultAsBytes[:])

	return resultAsBytes, resultAsText
}
