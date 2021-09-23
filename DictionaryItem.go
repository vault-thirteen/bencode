package bencode

// DictionaryItem is a 'bencode' dictionary item.
type DictionaryItem struct {
	// System Fields.
	Key   []byte
	Value interface{}

	// Additional Fields for special purposes.
	KeyStr   string
	ValueStr string
}
