package bencode

import (
	"bufio"
	"bytes"
	"fmt"
)

//	1.	Parser's settings.

//	1.1.	Integer size (number of ASCII letters allowed).
// N.B.: Maximum value of UInt64 is '18446744073709551615'.
const IntegerMaxLength = 20

//	1.2.	Byte string size header length (number of ASCII letters allowed).
const ByteStringSizeHeaderMaxLength = IntegerMaxLength

// Decoder is a 'bencode' decoder.
type Decoder struct {
	reader *bufio.Reader
}

// NewDecoder is the decoder's constructor.
func NewDecoder(
	reader *bufio.Reader,
) (d *Decoder) {
	d = &Decoder{
		reader: reader,
	}

	return d
}

// Decode decodes a 'bencoded' byte stream into an interface.
func (d Decoder) Decode() (result interface{}, err error) {
	return d.readBencodedValue()
}

// readBencodedValue reads a raw "bencoded" value, including its sub-values.
func (d Decoder) readBencodedValue() (result interface{}, err error) {

	// Get the first byte from stream to know its type.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	// Analyze the type.
	if b == HeaderDictionary {
		return d.readDictionary()

	} else if b == HeaderList {
		return d.readList()

	} else if b == HeaderInteger {
		return d.readInteger()

	} else if isByteNonNegativeAsciiNumeric(b) {
		// It must be an ASCII number indicating a byte string.
		// => Byte string.

		// Rewind the Cursor back, as it does not have a type-Prefix !
		// The 'bencode' encoding is ugly ...
		err = d.reader.UnreadByte()
		if err != nil {
			return nil, err
		}

		// Read the byte string.
		return d.readByteString()
	}

	// Otherwise, it is a syntax error.
	var errorArea = []byte{b}

	return nil, fmt.Errorf(ErrFSyntaxErrorAt, errorArea)
}

// readByteString reads a byte string from the stream (reader).
func (d Decoder) readByteString() (ba []byte, err error) {

	// Read the size header and verify it.
	var byteStringLen uint
	byteStringLen, err = d.readByteStringSizeHeader()
	if err != nil {
		return nil, err
	}

	// Now we should read the byte string.
	var b byte
	var i uint = 0
	var bytesAccumulator *bytes.Buffer
	bytesAccumulator = bytes.NewBuffer([]byte{})
	for i < byteStringLen {

		// Read the next symbol.
		b, err = d.reader.ReadByte()
		if err != nil {
			return nil, err
		}

		// Save the byte to the accumulator.
		err = bytesAccumulator.WriteByte(b)
		if err != nil {
			return nil, err
		}

		i++
	}

	ba = bytesAccumulator.Bytes()

	return ba, nil
}

// readByteStringSizeHeader reads the size header of a byte string from the
// stream (reader) and converts its value into an integer.
func (d Decoder) readByteStringSizeHeader() (byteStringLen uint, err error) {

	// Read the first byte.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return 0, err
	}

	var sizeHeader []byte
	for b != HeaderStringSizeValueDelimiter {

		// Syntax check.
		if !isByteNonNegativeAsciiNumeric(b) {
			var errorArea = append(sizeHeader, []byte{b}...)

			return 0, fmt.Errorf(ErrFSyntaxErrorAt, errorArea)
		}

		// Save the byte into the size header.
		if len(sizeHeader) < ByteStringSizeHeaderMaxLength {
			sizeHeader = append(sizeHeader, b)
		} else {
			// The length header is too big !
			var errorArea = append(sizeHeader, []byte{b}...)

			return 0, fmt.Errorf(
				ErrHeaderLengthError,
				errorArea,
			)
		}

		// Read the next byte.
		b, err = d.reader.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// Check the header's length.
	if len(sizeHeader) == 0 {
		var errorArea = append(sizeHeader, []byte{b}...)

		return 0, fmt.Errorf(ErrFSyntaxErrorAt, errorArea)
	}

	// Convert the size header into a normal integer size value.
	var byteStringLenUint64 uint64
	byteStringLenUint64, err = convertByteStringToNonNegativeInteger(sizeHeader)
	if err != nil {
		return 0, err
	}

	return uint(byteStringLenUint64), nil
}

// readDictionary reads a dictionary. We suppose that the header of the
// dictionary ('d') has already been read from the stream.
func (d Decoder) readDictionary() (result interface{}, err error) {

	// Prepare the data.
	var dictionary = make([]DictionaryItem, 0)

	// Probe the next byte to check the end of the dictionary.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	for b != FooterCommon {

		// That single byte (we probed) was not an End !
		// We must get back, rewind that byte.
		err = d.reader.UnreadByte()
		if err != nil {
			return nil, err
		}

		// Get the key.
		var dictKey []byte
		dictKey, err = d.readDictionaryKey()
		if err != nil {
			return nil, err
		}

		// Get the value.
		var dictValue interface{}
		dictValue, err = d.readDictionaryValue()
		if err != nil {
			return nil, err
		}

		// Save the item into the dictionary.
		dictionary = append(
			dictionary,
			DictionaryItem{
				// System Fields.
				Key:   dictKey,
				Value: dictValue,

				// Additional Fields for special purposes.
				KeyStr:   string(dictKey),
				ValueStr: convertInterfaceToString(dictValue),
			},
		)

		// Probe the next byte to check the end of the dictionary.
		b, err = d.reader.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	return dictionary, nil
}

// readDictionaryKey reads a dictionary's key.
func (d Decoder) readDictionaryKey() ([]byte, error) {
	return d.readByteString()
}

// readDictionaryValue reads a dictionary's value.
func (d Decoder) readDictionaryValue() (interface{}, error) {
	return d.readBencodedValue()
}

// readInteger reads an integer from the stream (reader). We suppose that the
// header of the integer ('i') has already been read from the stream.
func (d Decoder) readInteger() (value int64, err error) {

	// Prepare the data.
	var valueBA []byte

	// Read the first byte.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return 0, err
	}

	for b != FooterCommon {

		// Syntax check.
		if !isByteAsciiNumeric(b) {
			var errorArea = append(valueBA, []byte{b}...)

			return 0, fmt.Errorf(ErrFSyntaxErrorAt, errorArea)
		}

		// Save the byte into the value byte array.
		if len(valueBA) < IntegerMaxLength {
			valueBA = append(valueBA, b)
		} else {
			// The integer is too big !
			var errorArea = append(valueBA, []byte{b}...)

			return 0, fmt.Errorf(ErrFIntegerLengthError, errorArea)
		}

		// Read the next byte.
		b, err = d.reader.ReadByte()
		if err != nil {
			return 0, err
		}
	}

	// We have read the value.
	// Check that it is not empty.
	if len(valueBA) == 0 {
		var errorArea = append(valueBA, []byte{b}...)

		return 0, fmt.Errorf(ErrFSyntaxErrorAt, errorArea)
	}

	// Convert the value into a normal integer value.
	return convertByteStringToInteger(valueBA)
}

// readList reads a list from the stream (reader). We suppose that the header
// of the list ('l') has already been read from the stream.
func (d Decoder) readList() (list []interface{}, err error) {

	// Prepare the data.
	list = make([]interface{}, 0)

	// Probe the next byte to check the end of the list.
	var b byte
	b, err = d.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	for b != FooterCommon {

		// That single byte (we probed) was not an End !
		// We must get back, rewind that byte.
		err = d.reader.UnreadByte()
		if err != nil {
			return nil, err
		}

		// Get the item.
		var listItem interface{}
		listItem, err = d.readBencodedValue()
		if err != nil {
			return nil, err
		}

		// Save the item into the dictionary.
		list = append(list, listItem)

		// Probe the next byte to check the end of the list.
		b, err = d.reader.ReadByte()
		if err != nil {
			return nil, err
		}
	}

	return list, nil
}
