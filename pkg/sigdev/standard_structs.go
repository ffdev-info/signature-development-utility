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
	EOF        = "EOFoffset" // EOF represents a beginning of file offset in DROID.
	// Variable offsets are not written verbatim in DROID.
	VAR    = "Variable"
	VAROFF = "VARoffset"
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

/*
```xml
<?xml version="1.0" encoding="UTF-8"?>
<!-- Signature file namespace and metadata is important but none so more as
     the version here which helps DROID to list signature files in its
     graphical user interface.
-->
<FFSignatureFile xmlns="http://www.nationalarchives.gov.uk/pronom/SignatureFile" Version="1788527087" DateCreated="2026-09-03T12:22:45">
  <!-- Section 1: InternalSignatureCollection contains 1..* internal signatures
       associated with a file format.
  -->
  <InternalSignatureCollection>
    <InternalSignature ID="1">
      <ByteSequence Reference="BOFoffset" Sequence="" MinOffset="" MaxOffset=""/>
      <ByteSequence Reference="BOFoffset" Sequence="" MinOffset="" MaxOffset=""/>
    </InternalSignature>
    <InternalSignature ID="2">
      <ByteSequence Reference="BOFoffset" Sequence="" MinOffset="" MaxOffset=""/>
      <ByteSequence Reference="BOFoffset" Sequence="" MinOffset="" MaxOffset=""/>
    </InternalSignature>
    <InternalSignature ID="3">
      <ByteSequence Reference="BOFoffset" Sequence="" MinOffset="" MaxOffset=""/>
    </InternalSignature>
  </InternalSignatureCollection>
  <!-- Section 2: FileFormatCollection contains 1..* file format metadata
       sections and mappings to internal signatures with priorities, and
       file format extensions.
  -->
  <FileFormatCollection>
    <FileFormat ID="1" Name="FILE FORMAT NAME" PUID="CUSTOM/PUID" Version="V0" MIMEType="application/octet-stream">
      <InternalSignatureID>1</InternalSignatureID>
      <InternalSignatureID>2</InternalSignatureID>
      <Extension>ext</Extension>
    </FileFormat>
    <FileFormat ID="3" Name="FILE FORMAT NAME" PUID="CUSTOM/PUID" Version="V1" MIMEType="application/octet-stream">
      <InternalSignatureID>3</InternalSignatureID>
      <Extension>ext</Extension>
      <!-- A priority means that this record is selected by DROID if a file
           is also moatched to one it has a priority over.
      -->
      <HasPriorityOverFileFormatID>1</HasPriorityOverFileFormatID>
    </FileFormat>
  </FileFormatCollection>
</FFSignatureFile>
```
*/

// byteSequences...
type byteSequences struct {
	ByteSequence []byteSequence `xml:"ByteSequence"`
	ID           string         `xml:"ID,attr"`
	Specificity  string         `xml:"Specificity,attr,omitempty"`
}

// byteSequence...
type byteSequence struct {
	//SubSeq    sequence `xml:"SubSequence"`
	Reference string `xml:"Reference,attr,omitempty"`
	// Simple sig...
	Sequence  string `xml:"Sequence,attr,omitempty"`
	MinOffset string `xml:"MinOffset,attr,omitempty"`
	MaxOffset string `xml:"MaxOffset,attr,omitempty"`
}

// Compatibility.
type subSeq = byteSequence
type byteSeq = byteSequences

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
