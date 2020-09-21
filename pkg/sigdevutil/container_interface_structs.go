// Structures that enable us to receive information from a web-form and
// process that into a container signature file.

package sigdevutil

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ContainerSignatureInterface contains a modernized the top-level
// mapping of the information needed to describe a DROID container
// signature.
type ContainerSignatureInterface struct {
	Description string  // Description of the container signature provides metadata for the container signature file.
	PUID        string  // PRONOM PUID.
	TriggerPUID string  // The PUID used to access container signature identification, e.g. office open XML, OOXML.
	Files       []files // The files that make up the container.
}

// files contains the paths and sequences that we need to make up a
// container signature.
type files struct {
	Path      string      // The path of the file within the container with '/' path separator.
	Sequences []sequences // The standard PRONOM sequence patterns used to match the file inside the container (optional)
}

// ProcessSignature converts a form input into a container structure
// which can then be converted into a container signature file.
func (container *ContainerSignatureInterface) ProcessSignature(form url.Values) {
	// Form metadata fields.
	const containerCount = "container-file-count"
	// Metadata fields.
	const containerDescription = "container-description"
	const connectingPUID = "connecting-puid"
	const containerType = "container-type"
	// File fields.
	const containerPathPrefix = "container-path-"
	// Sequence fields.
	const signatureFieldPrefix = "signature-"
	const offsetFieldPrefix = "max-offset-"
	const maxOffsetFieldPrefix = "offset-"
	const relativityFieldPrefix = "relativity-"
	// Form metadata to help with processing. Can also be done using
	// the form's slices alone.
	numberOfFiles, _ := strconv.Atoi(form[containerCount][0])
	container.Description = strings.TrimSpace(form[containerDescription][0])
	container.PUID = strings.TrimSpace(form[connectingPUID][0])
	container.TriggerPUID = strings.TrimSpace(form[containerType][0])
	var containerFiles = make([]files, numberOfFiles)
	for idx := 0; idx < numberOfFiles; idx++ {
		pathField := fmt.Sprintf("%s%d", containerPathPrefix, idx)
		containerFiles[idx].Path = form[pathField][0]
		signatureField := fmt.Sprintf("%s%d", signatureFieldPrefix, idx)
		if signatures, ok := form[signatureField]; ok {
			offsetField := fmt.Sprintf("%s%d", offsetFieldPrefix, idx)
			maxOffsetField := fmt.Sprintf("%s%d", maxOffsetFieldPrefix, idx)
			relativityField := fmt.Sprintf("%s%d", relativityFieldPrefix, idx)
			offsets := form[offsetField]
			maxOffsets := form[maxOffsetField]
			relativity := form[relativityField]
			seqs := make([]sequences, len(signatures))
			for idx := range seqs {
				seqs[idx].Sequence = strings.TrimSpace(signatures[idx])
				seqs[idx].Offset, _ = strconv.Atoi(offsets[idx])
				seqs[idx].MaxOffset, _ = strconv.Atoi(maxOffsets[idx])
				seqs[idx].Relativity = relativity[idx]
			}
			// Add sequences to file.
			containerFiles[idx].Sequences = seqs
		}
	}
	// Add files to container.
	container.Files = containerFiles
}

// ToDROIDContainer will convert a container interface structure to
// something compatible with the DROID format identification tool.
func (container *ContainerSignatureInterface) ToDROIDContainer() ContainerSignatureMapping {

	const internalSignatureID = "1"
	const defaultPosition = "1"

	var containerSignatureID string = "1"

	var fileFormatMapping FileFormatMapping
	fileFormatMapping.SignatureID = containerSignatureID
	fileFormatMapping.PUID = container.PUID

	var allFormatMappings FileFormatMappings
	allFormatMappings.FFMap = fileFormatMapping

	allFiles := len(container.Files)

	internalObjects := make([]fileContainer, allFiles)

	for idx, file := range container.Files {

		var binarySignatures = new(binarySigs)

		if len(file.Sequences) > 0 {
			byteSequences := len(file.Sequences)
			subSequences := make([]subSeq, byteSequences)
			for idx, sig := range file.Sequences {
				var signatureSequence sequence
				signatureSequence.Sequence = sig.Sequence
				signatureSequence.Position = defaultPosition
				signatureSequence.SubSeqMin = strconv.Itoa(sig.Offset)
				if sig.Relativity != VAR {
					signatureSequence.SubSeqMax = strconv.Itoa(sig.MaxOffset)
					subSequences[idx].Reference = sig.Relativity
				}
				subSequences[idx].SubSeq = signatureSequence
			}
			var byteSequence = make([]byteSeq, 1)
			byteSequence[0].ByteSequence = subSequences
			byteSequence[0].ID = internalSignatureID

			var internalSignature = new(intSig)
			internalSignature.InternalSignature = byteSequence

			binarySignatures.InternalSignatureCollection = internalSignature
		}
		internalObjects[idx].Path = file.Path
		// Conditionally output the binary signatures because if they
		// are otherwise empty DROID does not need to see these in the
		// signature file.
		if binarySignatures != nil &&
			binarySignatures.InternalSignatureCollection != nil {
			internalObjects[idx].BinarySignatures = binarySignatures
		}
	}

	var containerSignatureObj containerSignature
	containerSignatureObj.File = internalObjects

	var topLevelContainerSig contSig
	topLevelContainerSig.Description = container.Description
	topLevelContainerSig.Files = containerSignatureObj
	topLevelContainerSig.ID = containerSignatureID
	topLevelContainerSig.ContainerType = container.TriggerPUID

	var allTopLevelContainerSigs containerSignatures
	allTopLevelContainerSigs.ContainerSignature = topLevelContainerSig

	var containerSignatureMapping ContainerSignatureMapping
	containerSignatureMapping.ContainerSignatures = allTopLevelContainerSigs
	containerSignatureMapping.SchemaVersion = containerSchemaVersion
	containerSignatureMapping.SignatureVersion = containerSignatureVersion
	containerSignatureMapping.FileFormatMappings = allFormatMappings

	return containerSignatureMapping
}

// GetFileName is a small helper function that helps us make some
// useful metadata about our output.
func (container *ContainerSignatureInterface) GetFileName() string {
	const devSig = "dev-container-signature"
	nicePUID := strings.Replace(container.PUID, "/", "-", 1)
	return fmt.Sprintf("%s-%s-%s", devSig, nicePUID, generateDateNoSpaces())
}
