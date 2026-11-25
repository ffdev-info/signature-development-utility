// Tests for the different interfaces created for creating DROID
// compatible signature files.

package sigdevutil

import (
	"testing"
	"time"
)

// setupSignature will setup a signature to use in testing...
func setupSignatureInterface() SignatureInterface {
	var sig SignatureInterface
	sig.PUID = "fmt/123"
	sig.FormatName = "test format"
	sig.VersionNumber = "1.0"
	sig.MimeType = "application/octet-stream"
	sig.Extension = "ext"
	seqs := []string{"1234", "3456", "7890"}
	rels := []string{BOF, EOF, VAR}
	sequences := make([]sequences, len(seqs))
	for idx := 0; idx < len(seqs); idx++ {
		sequences[idx].Sequence = seqs[idx]
		sequences[idx].Offset = 0
		sequences[idx].MaxOffset = 0
		sequences[idx].Relativity = rels[idx]
	}
	sig.Sequences = sequences
	return sig
}

// TestStandardSingatureGeneration will test the standard signature
// generation to make sure we're making something we know that works
// with DROID.
func TestStandardSingatureGeneration(t *testing.T) {
	// Replace time.Now() with a placeholder we can guarantee the result
	// of when we call it.
	now = func() time.Time { return time.Date(1970, 1, 1, 0, 0, 0, 651387237, time.UTC) }
	signatureInterface := setupSignatureInterface()
	standardSignatureFile := signatureInterface.ToDROID(false)
	// Compare the output with our fixture.
	if standardSignatureFile.String() != standardSignatureOne {
		t.Errorf(
			"Test XML doesn't match expected XML, outputting test XML: %s",
			standardSignatureFile.String(),
		)
	}
}

// TestStandardSignatureFilename will simply test that we can get a
// nice filename from the type.
func TestStandardSignatureFilename(t *testing.T) {
	const expected = "development-signature-fmt-123"
	sig := setupSignatureInterface()
	res := sig.GetFileName()
	if res != expected {
		t.Errorf("We didn't get the anticipated file name '%s', got '%s", expected, res)
	}
}

// setupSignature will setup a signature to use in testing...
func setupContainerSignatureInterface() ContainerSignatureInterface {
	var sig ContainerSignatureInterface
	sig.Description = "test format"
	sig.PUID = "fmt/39"
	sig.TriggerPUID = "OLE2"
	internalObjects := []string{"WordDocument", "CompObj"}
	internalFiles := make([]files, len(internalObjects))
	// Create a single sequence for the compObj file in this OLE2.
	compObjSeq := "10000000576F72642E446F63756D656E742E"
	seqs := make([]sequences, 1)
	seqs[0].Sequence = compObjSeq
	seqs[0].Offset = 40
	seqs[0].MaxOffset = 1024
	seqs[0].Relativity = BOF
	internalFiles[0].Path = internalObjects[0]
	internalFiles[1].Path = internalObjects[1]
	internalFiles[1].Sequences = seqs
	sig.Files = internalFiles
	return sig
}

// TestContainerSignatureGeneration will test the generation of
// container signature files compatible with DROID.
func TestContainerSignatureGeneration(t *testing.T) {
	signatureInterface := setupContainerSignatureInterface()
	containerSingatureFile := signatureInterface.ToDROIDContainer()
	// Compare the output with our fixture.
	if containerSingatureFile.String() != containerSignatureOne {
		t.Errorf(
			"Test XML doesn't match expected XML, outputting test XML: %s",
			containerSingatureFile.String(),
		)
	}
}

// TestStandardSignatureFilename will simply test that we can get a
// nice filename from the type.
func TestContainerSignatureFilename(t *testing.T) {
	const expected = "development-container-signature-fmt-39"
	sig := setupContainerSignatureInterface()
	res := sig.GetFileName()
	if res != expected {
		t.Errorf("We didn't get the anticipated file name '%s', got '%s", expected, res)
	}
}
