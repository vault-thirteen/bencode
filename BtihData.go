package bencode

// BtihData is the information about the BitTorrent Info Hash (BTIH) check sum
// stored both as a text and as an array of bytes.
type BtihData struct {
	Bytes Sha1Sum
	Text  string
}
