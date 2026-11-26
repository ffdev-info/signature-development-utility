// Structures needed for creating standard (non-container) signature
// files compatible with DROID.

package sigdev

import (
	"encoding/xml"
	"fmt"
)

// Relativity constants.
const (
	BOF string = "BOFoffset" // BOF represents a beginning of file offset in DROID.
	EOF string = "EOFoffset" // EOF represents a beginning of file offset in DROID.
	// Variable offsets are not written verbatim in DROID.
	//
	// NB. we have more control over these, but testing showed a few
	// divergences. Pick one, and remain consistent throughout the
	// code. Ensure tests exist that verify this. DROID only needs
	// a blank field.
	VAR      string = "Variable"
	VAROFF   string = "VARoffset"
	VARBASIC string = "VAR"
)

// FFSignatureFile encapsulates an entire DROID signature file and all
// roads lead to Rome from here...
type FFSignatureFile struct {
	InternalSignatureCollection intSig `xml:"InternalSignatureCollection"`
	FileFormatCollection        ffColl `xml:"FileFormatCollection"`
	Xmlns                       string `xml:"xmlns,attr"`
	Version                     string `xml:"Version,attr"`
	DateCreated                 string `xml:"DateCreated,attr"`
}

// ffColl ...
type ffColl struct {
	FileFormat []ffformat `xml:"FileFormat"`
}

// ffformat ...
type ffformat struct {
	ID       string `xml:"ID,attr"`
	Name     string `xml:"Name,attr"`
	PUID     string `xml:"PUID,attr"`
	Version  string `xml:"Version,attr"`
	MIMEType string `xml:"MIMEType,attr"`

	InternalSignatureID         string `xml:"InternalSignatureID,omitempty"`
	Extension                   string `xml:"Extension,omitempty"`
	HasPriorityOverFileFormatID string `xml:"HasPriorityOverFileFormatID,omitempty"`
}

// intSig ...
type intSig struct {
	InternalSignature []byteSequences `xml:"InternalSignature"`
}

// byteSequences...
type byteSequences struct {
	ByteSequence []byteSequence `xml:"ByteSequence"`
	ID           string         `xml:"ID,attr"`
	Specificity  string         `xml:"Specificity,attr,omitempty"`
}

// byteSequence...
type byteSequence struct {
	Reference string `xml:"Reference,attr,omitempty"`
	// DROID simple signature.
	Sequence  string `xml:"Sequence,attr,omitempty"`
	MinOffset string `xml:"MinOffset,attr,omitempty"`
	MaxOffset string `xml:"MaxOffset,attr,omitempty"`
}

// String returns the signature file as a string.
func (signatureFile FFSignatureFile) String() string {
	out, _ := xml.MarshalIndent(signatureFile, "", " ")
	return fmt.Sprintf("%s%s", xml.Header, out)
}

// String returns the XML representation of the struct and includes the
// XML's declaration.
func (intSigColl intSig) String() string {
	out, _ := xml.MarshalIndent(intSigColl, "", " ")
	return fmt.Sprintf("%s%s", xml.Header, out)
}

// String returns the XML representation of the struct and includes the
// XML's declaration.
func (formatColl ffColl) String() string {
	out, _ := xml.MarshalIndent(formatColl, "", " ")
	return fmt.Sprintf("%s%s", xml.Header, out)
}
