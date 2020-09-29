// Structures that enable us to receive information from a web-form and
// process that into a standard signature file.

package sigdevutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// SignatureInterface provides a modern mapping for the DROID signature
// file, and provides opportunities to map DROID like patterns into
// other formats in future.
type SignatureInterface struct {
	PUID          string      // PRONOM PUID.
	FormatName    string      // File format name.
	VersionNumber string      // Version number of the format.
	MimeType      string      // File format MIMEtype.
	Extension     string      // File format extension.
	Sequences     []sequences // File format signature sequences.
}

// sequences ...
type sequences struct {
	Sequence   string // The magic-number pattern to match in the file composed of PRONOM regex.
	Offset     int    // The offset within the file we will find the sequence.
	MaxOffset  int    // Maximum offset within the file we will find the sequence.
	Relativity string // Relativity within the file we will find the sequence from, e.g. from End of file.
}

// ProcessSignature will process a signature development form input and
// convert it into a standard signature structure which can then be
// output as a standard DROID signature.
func (signature *SignatureInterface) ProcessSignature(form url.Values) {
	// Form constants. These need to be turned into template values at
	// some point.
	const name = "format-name"
	const ext = "format-extension"
	const version = "format-version"
	const puid = "puid"
	const mimetype = "mimetype"
	const triggers = "triggers"
	// Signature metadata.
	signature.PUID = strings.TrimSpace(form[puid][0])
	signature.FormatName = strings.TrimSpace(form[name][0])
	signature.VersionNumber = strings.TrimSpace(form[version][0])
	signature.MimeType = strings.TrimSpace(form[mimetype][0])
	signature.Extension = strings.TrimSpace(form[ext][0])
	// Signature information.
	const signatureField = "signature-input-0"
	const offsetField = "offset-0"
	const maxOffsetField = "max-offset-0"
	const signatureRelativityField = "signature-relativity-0"
	if _, ok := form[triggers]; ok {
		// We're not going to add this sequence to the signature file,
		// we'll just add the metadata.
		return
	}
	if signatures, ok := form[signatureField]; ok {
		seqs := make([]sequences, len(signatures))
		// Form fields.
		offset := form[offsetField]
		maxOffset := form[maxOffsetField]
		relativity := form[signatureRelativityField]
		for idx := range signatures {
			seqs[idx].Sequence = strings.TrimSpace(signatures[idx])
			seqs[idx].Offset, _ = strconv.Atoi(offset[idx])
			rel := relativity[idx]
			if rel != VAR {
				seqs[idx].MaxOffset, _ = strconv.Atoi(maxOffset[idx])
				seqs[idx].Relativity = relativity[idx]
			}
		}
		signature.Sequences = seqs
	}
}

// ToDROID will convert our signature file to a DROID compatible format.
func (signature *SignatureInterface) ToDROID(triggers bool) FFSignatureFile {
	// Notes on working with the DROID signature file structs: Work
	// backwards if possible. From the lower 'leafs' of the structure
	// to the top level where we specify the file format information.
	// Largely because the DROID signature file is designed as XML and
	// is fairly counter-intuitive to how we might otherwise like to
	// work nowadays.
	var zipInt, ooxmlInt, ole2Int byteSeq
	var zipFF, ooxmlFF, ole2FF ffformat
	if triggers {
		zipInt, zipFF = getZIPTrigger()
		ooxmlInt, ooxmlFF = getOOXMLTrigger()
		ole2Int, ole2FF = getOLE2Trigger()
	}

	// This will be automatically incremented once we are working with
	// arrays of signatures.
	const internalSignatureID = "1"
	const specificity = "Specific"

	const defaultMinFragLen = "0"
	const defaultPosition = "1"

	byteSequences := len(signature.Sequences)

	subSequences := make([]subSeq, byteSequences)

	for idx, sig := range signature.Sequences {
		var signatureSequence sequence
		signatureSequence.Sequence = sig.Sequence
		signatureSequence.MinFragLen = defaultMinFragLen
		signatureSequence.Position = defaultPosition
		signatureSequence.SubSeqMin = strconv.Itoa(sig.Offset)
		if sig.Relativity != VAR && sig.Relativity != "" {
			signatureSequence.SubSeqMax = strconv.Itoa(sig.MaxOffset)
			subSequences[idx].Reference = sig.Relativity
		}
		subSequences[idx].SubSeq = signatureSequence
	}

	var byteSequence = make([]byteSeq, 1)
	byteSequence[0].ByteSequence = subSequences

	byteSequence[0].ID = internalSignatureID
	byteSequence[0].Specificity = specificity

	var internalSignature intSig

	if triggers {
		byteSequence = append(byteSequence, zipInt)
		byteSequence = append(byteSequence, ooxmlInt)
		byteSequence = append(byteSequence, ole2Int)
		_, byteSequence = byteSequence[0], byteSequence[1:]
	}

	internalSignature.InternalSignature = byteSequence

	var format = make([]ffformat, 1)

	format[0].Name = signature.FormatName
	format[0].PUID = signature.PUID
	format[0].Version = signature.VersionNumber
	format[0].MIMEType = signature.MimeType
	format[0].Extension = signature.Extension
	format[0].ID = internalSignatureID
	if !triggers {
		format[0].InternalSignatureID = internalSignatureID
	}

	var formatCollection ffColl

	if triggers {
		format = append(format, zipFF)
		format = append(format, ooxmlFF)
		format = append(format, ole2FF)
	}
	formatCollection.FileFormat = format

	var droidSignature FFSignatureFile
	droidSignature.Xmlns = signatureFileNamespace
	droidSignature.Version = signatureFileVersion
	droidSignature.DateCreated = generateDate()

	droidSignature.InternalSignatureCollection = internalSignature
	droidSignature.FileFormatCollection = formatCollection

	return droidSignature
}

// GetFileName is a small helper function that helps us make some
// useful metadata about our output.
func (signature *SignatureInterface) GetFileName() string {
	const sigFile = "signature-file"
	niceName := formatFilenameString(signature.FormatName)
	niceVersion := formatFilenameString(signature.VersionNumber)
	return fmt.Sprintf("%s-%s-%s", niceName, niceVersion, sigFile)
}

// Bootstrap signature development utility 2.0 to Signature development
// utility 1.0.
func (signature *SignatureInterface) ToPHP(port string) string {

	counter := strconv.Itoa(len(signature.Sequences))

	const ORIGINALURL = "http://localhost:%s/php/process_signature_form.php"

	const count = "counter"
	const name = "name1"
	const version = "version1"
	const ext = "extension1"
	const mime = "mimetype1"
	const puid = "puid1"

	data := url.Values{
		count:   {counter},
		name:    {signature.FormatName},
		version: {signature.VersionNumber},
		ext:     {signature.Extension},
		mime:    {signature.MimeType},
		puid:    {signature.PUID},
	}

	const sig = "signature"
	const anchor = "anchor"
	const offset = "offset"
	const maxoffset = "maxoffset"

	const variableAnchor = "Variable"

	for idx := 1; idx <= len(signature.Sequences); idx++ {
		sigField := fmt.Sprintf("%s%d", sig, idx)
		anchorField := fmt.Sprintf("%s%d", anchor, idx)
		offsetField := fmt.Sprintf("%s%d", offset, idx)
		maxOffsetField := fmt.Sprintf("%s%d", maxoffset, idx)

		sequence := signature.Sequences[idx-1]

		data[sigField] = []string{sequence.Sequence}

		// Variable sequence handling for 1.0.
		data[anchorField] = []string{sequence.Relativity}
		// Variable offsets are no longer set by the time it reaches
		// this point in the code, so identify it by negating BOF and EOF.
		if sequence.Relativity != BOF && sequence.Relativity != EOF {
			data[anchorField] = []string{variableAnchor}
		}

		data[offsetField] = []string{strconv.Itoa(sequence.Offset)}
		data[maxOffsetField] = []string{strconv.Itoa(sequence.MaxOffset)}
	}

	originalURL := fmt.Sprintf(ORIGINALURL, port)
	fmt.Fprintf(os.Stderr, "Bootstrap URL: %s", originalURL)

	resp, err := http.PostForm(originalURL, data)
	if err != nil {
		return fmt.Sprintf("Error sending request: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error sending request: %s", err)
	}
	return string(body)
}
