package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/neverwinter-nights/bencode"
)

func main() {
	var err = checkFileHashSum()
	checkError(err)
}

func checkFileHashSum() (err error) {
	// Settings.
	const (
		ExampleFolder = "example"
		DataFolder    = "data"
		Btih          = "9ddf6a9b17b624991b39f8afd2edc64f673350e3"
		FileName      = "5942384.torrent"
	)

	// Parse the file.
	var f = bencode.NewFile(
		filepath.Join(ExampleFolder, DataFolder, FileName),
	)

	var decodedObject *bencode.DecodedObject
	decodedObject, err = f.Parse()
	if err != nil {
		return err
	}

	// Check the BTIH.
	var ok = (decodedObject.BTIH.Text == Btih)

	var btihBytesExpected []byte
	btihBytesExpected, err = hex.DecodeString(Btih)
	if err != nil {
		return err
	}

	ok = ok && (bytes.Equal(
		decodedObject.BTIH.Bytes[:],
		btihBytesExpected,
	))

	if !ok {
		err = errors.New(fmt.Sprintf(
			"BTIH Mismatch. Expected:%v. Got:%v.",
			Btih,
			decodedObject.BTIH.Text,
		))

		return err
	}

	return nil
}

// checkError checks an error and prints the result to the std::out.
func checkError(err error) {
	if err != nil {
		fmt.Println("Error.", err)

		return
	}

	fmt.Println("OK")
}
