package bencode

//	1. Special symbols of 'bencode' encoding.

//	1.1. Headers.
const (
	HeaderDictionary               byte = 'd'
	HeaderInteger                  byte = 'i'
	HeaderList                     byte = 'l'
	HeaderStringSizeValueDelimiter byte = ':'
)

//	1.2. Footers.
const (
	FooterCommon byte = 'e'
)

//	1.3. Sections of a BitTorrent file.
const (
	FileSectionAnnounce     = "announce"
	FileSectionAnnounceList = "announce-list"
	FileSectionCreationDate = "creation date"
	FileSectionComment      = "comment"
	FileSectionCreatedBy    = "created by"
	FileSectionEncoding     = "encoding"
	FileSectionInfo         = "info"
)
