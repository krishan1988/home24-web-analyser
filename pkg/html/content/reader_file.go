package content

import (
	"io/ioutil"
	"strings"

	"golang.org/x/net/html"
)

// readerWithFile represents a reader that reads HTML content from a file.
type readerWithFile struct {
	filename string // Filename is the name of the HTML file to read.
}

// NewReaderWithFile creates a new instance of ReaderWithFile.
func NewReaderWithFile(filename string) Reader {
	return &readerWithFile{
		filename: filename,
	}
}

// Read reads the HTML content from the file and returns the root node of the parsed HTML tree.
func (r *readerWithFile) Read() (*html.Node, error) {
	// Read the file contents.
	fileContents, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}

	// Parse the HTML content.
	rootNode, err := html.Parse(strings.NewReader(string(fileContents)))
	if err != nil {
		return nil, err
	}

	return rootNode, nil
}