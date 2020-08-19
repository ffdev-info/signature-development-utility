// Structures needed for creating standard (non-container) signature
// files compatible with DROID.

package sigdevutil

import (
	"encoding/xml"
	"fmt"
)

// Relativity constants.
const (
	BOF string = "BOFoffset" // BOF represents a beginning of file offset in DROID.
	EOF        = "EOFoffset" // EOF represents a beginning of file offset in DROID.
	VAR        = "VARoffset" // VAR is a helper variable that we use to negate a variable offset's output in the code below.
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
	InternalSignature []byteSeq `xml:"InternalSignature"`
}

// byteSeq ...
type byteSeq struct {
	ByteSequence []subSeq `xml:"ByteSequence"`
	ID           string   `xml:"ID,attr"`
	Specificity  string   `xml:"Specificity,attr,omitempty"`
}

// subSeq ...
type subSeq struct {
	SubSeq    sequence `xml:"SubSequence"`
	Reference string   `xml:"Reference,attr,omitempty"`
}

// sequence ...
type sequence struct {
	Sequence   string `xml:"Sequence"`
	Position   string `xml:"Position,attr"`
	MinFragLen string `xml:"MinFragLength,attr,omitempty"`
	SubSeqMin  string `xml:"SubSeqMinOffset,attr"`
	SubSeqMax  string `xml:"SubSeqMaxOffset,attr,omitempty"`
}

// String returns the signature file as a string.
func (signatureFile FFSignatureFile) String() string {
	out, _ := xml.MarshalIndent(signatureFile, "", " ")
	return fmt.Sprintf("%s%s", xml.Header, out)
}

// output returns the XML representation of the struct and includes the
// XML's declaration.
func (intSigColl intSig) String() string {
	out, _ := xml.MarshalIndent(intSigColl, "", " ")
	return fmt.Sprintf("%s%s", xml.Header, out)
}

// output returns the XML representation of the struct and includes the
// XML's declaration.
func (formatColl ffColl) String() string {
	out, _ := xml.MarshalIndent(formatColl, "", " ")
	return fmt.Sprintf("%s%s", xml.Header, out)
}
