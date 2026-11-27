// Create standard signature patterns needed to trigger container
// matching.

package sigdev

var zipSequences = []sequences{
	{
		Sequence:   "504B0304",
		Offset:     0,
		MaxOffset:  4,
		Relativity: "BOF",
	},
	{
		Sequence:   "504B01",
		Offset:     61,
		MaxOffset:  65565,
		Relativity: "EOF",
	},
	{
		Sequence:   "504B0506",
		Offset:     0,
		MaxOffset:  65535, // 65565?
		Relativity: "EOF",
	},
}

// Create the entries needed for the ZIP trigger in the standard
// signature file.
func getZIPTrigger() (byteSequences, ffformat) {
	// Create ZIP file signature sequence.
	/*
		X-FMT/263
		Position type	Absolute from BOF
		Offset	0
		Maximum Offset	4
		Byte order
		Value	504B0304
		Position type	Absolute from EOF
		Offset	0
		Byte order	Little-endian
		Value	504B01{43-65531}504B0506{18-65531}
	*/
	var zipSig SignatureInterface
	zipSig.PUID = "x-fmt/263"
	zipSig.FormatName = "ZIP Format"
	zipSig.MimeType = "application/zip"
	zipSig.Extension = "zip"
	zipSig.Sequences = zipSequences
	zip := zipSig.ToDROID(false)
	zip.InternalSignatureCollection.InternalSignature[0].ID = "2"
	zip.FileFormatCollection.FileFormat[0].InternalSignatureID = "2"
	zip.FileFormatCollection.FileFormat[0].ID = "2"
	return zip.InternalSignatureCollection.InternalSignature[0], zip.FileFormatCollection.FileFormat[0]
}

var ooxmlSequences = []sequences{
	{
		Sequence:   "504B0304",
		Offset:     0,
		MaxOffset:  0,
		Relativity: "BOF",
	},
	{
		Sequence:   "5B436F6E74656E745F54797065735D2E786D6C20A2",
		Offset:     4,
		MaxOffset:  30,
		Relativity: "BOF",
	},
	{
		Sequence:   "504B0102",
		Offset:     0,
		MaxOffset:  65535,
		Relativity: "EOF",
	},
	{
		Sequence:   "504B0506",
		Offset:     0,
		MaxOffset:  65535,
		Relativity: "EOF",
	},
}

// Create the entries needed for the OOXML trigger in the standard
// signature file.
func getOOXMLTrigger() (byteSequences, ffformat) {
	// Create OOXML sequence.
	/*
		FMT/189
		Position type	Absolute from BOF
		Offset	0
		Maximum Offset	0
		Byte order	Little-endian
		Value	504B0304{26}5B436F6E74656E745F54797065735D2E786D6C20A2*504B0102*504B0506
	*/
	var ooxmlSig SignatureInterface
	ooxmlSig.PUID = "fmt/189"
	ooxmlSig.FormatName = "Microsoft Office Open XML"
	ooxmlSig.MimeType = "application/octet-stream"
	ooxmlSig.Extension = ""
	ooxmlSig.Sequences = ooxmlSequences
	ooxml := ooxmlSig.ToDROID(false)
	ooxml.InternalSignatureCollection.InternalSignature[0].ID = "3"
	ooxml.FileFormatCollection.FileFormat[0].InternalSignatureID = "3"
	ooxml.FileFormatCollection.FileFormat[0].ID = "3"
	return ooxml.InternalSignatureCollection.InternalSignature[0], ooxml.FileFormatCollection.FileFormat[0]
}

var ole2Sequences = []sequences{
	{
		Sequence:   "D0CF11E0A1B11AE1",
		Offset:     0,
		MaxOffset:  0,
		Relativity: "BOF",
	},
	{
		Sequence:   "FEFF",
		Offset:     0,
		MaxOffset:  28,
		Relativity: "BOF",
	},
}

// Create the entries needed for the OLE2 trigger in the standard
// signature file.
func getOLE2Trigger() (byteSequences, ffformat) {
	// Create OLE2 signature sequence.
	/*
		FMT/111
		Position type	Absolute from BOF
		Offset	0
		Byte order
		Value	D0CF11E0A1B11AE1{20}FEFF
	*/
	var ole2Sig SignatureInterface
	ole2Sig.PUID = "fmt/111"
	ole2Sig.FormatName = "OLE2 Compound Document Format"
	ole2Sig.MimeType = "application/octet-stream"
	ole2Sig.Extension = ""
	ole2Sig.Sequences = ole2Sequences
	ole2 := ole2Sig.ToDROID(false)
	ole2.InternalSignatureCollection.InternalSignature[0].ID = "4"
	ole2.FileFormatCollection.FileFormat[0].InternalSignatureID = "4"
	ole2.FileFormatCollection.FileFormat[0].ID = "4"
	return ole2.InternalSignatureCollection.InternalSignature[0], ole2.FileFormatCollection.FileFormat[0]
}
