package bencode

//	Special symbols of 'bencode' encoding.

// 1. Headers.
const (
	HeaderDictionary               byte = 'd'
	HeaderInteger                  byte = 'i'
	HeaderList                     byte = 'l'
	HeaderStringSizeValueDelimiter byte = ':'
)

// 2. Footers.
const (
	FooterCommon byte = 'e'
)
