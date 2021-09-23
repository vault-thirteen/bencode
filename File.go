package bencode

import (
	"bufio"
	"io/ioutil"
	"os"
	"time"
)

// File is a file.
type File struct {
	path   string
	osFile *os.File
}

// NewFile is a file's constructor.
func NewFile(
	filePath string,
) (f *File) {
	f = &File{
		path: filePath,
	}

	return f
}

// close closes a file.
func (f *File) close() (err error) {
	return f.osFile.Close()
}

// getContents reads the contents of an opened file.
func (f File) getContents() (fileContents []byte, err error) {

	// Fool check.
	if f.osFile == nil {
		return nil, ErrFileNotInitialized
	}

	_, err = f.osFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	fileContents, err = ioutil.ReadAll(f.osFile)
	if err != nil {
		return nil, err
	}

	return fileContents, nil
}

// open opens a file.
func (f *File) open() (err error) {
	f.osFile, err = os.Open(f.path)
	if err != nil {
		return err
	}

	return nil
}

// Parse parses an input file into an interface. It also stores some
// additional data, all packed into an object.
func (f File) Parse() (result *DecodedObject, err error) {

	// Open the file and prepare a stream reader.
	err = f.open()
	if err != nil {
		return nil, err
	}

	defer func() {
		// Close the file.
		var derr error
		derr = f.close()
		if derr != nil {
			err = combineErrors(err, derr)
		}
	}()

	var bufioReader = bufio.NewReader(f.osFile)

	// Parse the file encoded with 'bencode' encoding into an object.
	var decoder = NewDecoder(bufioReader)
	var ifc interface{}
	ifc, err = decoder.readBencodedValue()
	if err != nil {
		return nil, err
	}

	// Get the file contents.
	var fileContents []byte
	fileContents, err = f.getContents()
	if err != nil {
		return nil, err
	}

	// Prepare the result.
	var decodedObject *DecodedObject
	decodedObject = &DecodedObject{
		FilePath:        f.path,
		SourceData:      fileContents,
		DecodedObject:   ifc,
		DecodeTimestamp: time.Now().Unix(),
	}

	// Perform a self-check.
	var ok bool
	ok = decodedObject.MakeSelfCheck()
	if !ok {
		return nil, ErrSelfCheck
	}

	// Calculate the BTIH.
	err = decodedObject.CalculateBtih()
	if err != nil {
		return nil, err
	}

	return decodedObject, nil
}
