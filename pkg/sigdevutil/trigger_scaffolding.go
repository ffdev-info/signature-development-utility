// Create standard signature patterns needed to trigger container
// matching.

package sigdevutil

// Create the entries needed for the ZIP trigger in the standard
// signature file.
func getZIPTrigger() (byteSeq, ffformat) {
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
	seqs := []string{"504B0304", "504B01{43-65531}504B0506{18-65531}"}
	rels := []string{BOF, EOF}
	zipSequences := make([]sequences, len(seqs))
	for idx := 0; idx < len(seqs); idx++ {
		zipSequences[idx].Sequence = seqs[idx]
		zipSequences[idx].Offset = 0
		zipSequences[idx].MaxOffset = 4
		zipSequences[idx].Relativity = rels[idx]
	}
	zipSig.Sequences = zipSequences

	zip := zipSig.ToDROID(false)
	zip.InternalSignatureCollection.InternalSignature[0].ID = "2"
	zip.FileFormatCollection.FileFormat[0].InternalSignatureID = "2"
	zip.FileFormatCollection.FileFormat[0].ID = "2"

	return zip.InternalSignatureCollection.InternalSignature[0], zip.FileFormatCollection.FileFormat[0]
}

// Create the entries needed for the OOXML trigger in the standard
// signature file.
func getOOXMLTrigger() (byteSeq, ffformat) {
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
	seqs := []string{"504B0304{26}5B436F6E74656E745F54797065735D2E786D6C20A2*504B0102*504B0506"}
	rels := []string{BOF}
	ooxmlSequences := make([]sequences, len(seqs))
	for idx := 0; idx < len(seqs); idx++ {
		ooxmlSequences[idx].Sequence = seqs[idx]
		ooxmlSequences[idx].Offset = 0
		ooxmlSequences[idx].MaxOffset = 0
		ooxmlSequences[idx].Relativity = rels[idx]
	}
	ooxmlSig.Sequences = ooxmlSequences

	ooxml := ooxmlSig.ToDROID(false)
	ooxml.InternalSignatureCollection.InternalSignature[0].ID = "3"
	ooxml.FileFormatCollection.FileFormat[0].InternalSignatureID = "3"
	ooxml.FileFormatCollection.FileFormat[0].ID = "3"

	return ooxml.InternalSignatureCollection.InternalSignature[0], ooxml.FileFormatCollection.FileFormat[0]
}

// Create the entries needed for the OLE2 trigger in the standard
// signature file.
func getOLE2Trigger() (byteSeq, ffformat) {
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
	seqs := []string{"D0CF11E0A1B11AE1{20}FEFF"}
	rels := []string{BOF}
	ole2Sequences := make([]sequences, len(seqs))
	for idx := 0; idx < len(seqs); idx++ {
		ole2Sequences[idx].Sequence = seqs[idx]
		ole2Sequences[idx].Offset = 0
		ole2Sequences[idx].MaxOffset = 0
		ole2Sequences[idx].Relativity = rels[idx]
	}
	ole2Sig.Sequences = ole2Sequences

	ole2 := ole2Sig.ToDROID(false)
	ole2.InternalSignatureCollection.InternalSignature[0].ID = "4"
	ole2.FileFormatCollection.FileFormat[0].InternalSignatureID = "4"
	ole2.FileFormatCollection.FileFormat[0].ID = "4"

	return ole2.InternalSignatureCollection.InternalSignature[0], ole2.FileFormatCollection.FileFormat[0]
}
