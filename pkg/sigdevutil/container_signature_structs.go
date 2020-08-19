// Structures associated with the output of a container signature file
// compatible with DROID.

package sigdevutil

import (
	"encoding/xml"
	"fmt"
)

// ContainerSignatureMapping ...
type ContainerSignatureMapping struct {
	ContainerSignatures containerSignatures `xml:"ContainerSignatures"`
	SchemaVersion       string              `xml:"SchemaVersion,attr"`
	SignatureVersion    string              `xml:"SignatureVersion,attr"`
	FileFormatMappings  FileFormatMappings  `xml:"FileFormatMappings"`
	TriggerPuids        TriggerPuids        `xml:"TriggerPuids"`
}

// containerSignatures ...
type containerSignatures struct {
	ContainerSignature contSig `xml:"ContainerSignature"`
}

// contSig ...
type contSig struct {
	Description   string             `xml:"Description"`
	Files         containerSignature `xml:"Files"`
	ID            string             `xml:"Id,attr"`
	ContainerType string             `xml:"ContainerType,attr"`
}

// containerSignature ...
type containerSignature struct {
	File []fileContainer `xml:"File"`
}

// fileContainer ...
type fileContainer struct {
	Path             string      `xml:"Path"`
	BinarySignatures *binarySigs `xml:"BinarySignatures,omitempty"`
}

type binarySigs struct {
	InternalSignatureCollection *intSig `xml:"InternalSignatureCollection,omitempty"`
}

// FileFormatMappings ...
type FileFormatMappings struct {
	FFMap FileFormatMapping `xml:"FileFormatMapping"`
}

// FileFormatMapping ...
type FileFormatMapping struct {
	SignatureID string `xml:"signatureId,attr"`
	PUID        string `xml:"Puid,attr"`
}

// TriggerPuids ...
type TriggerPuids struct {
	TriggerPUID []TriggerPuidContainer `xml:"TriggerPuid"`
}

// TriggerPuidContainer ...
type TriggerPuidContainer struct {
	ContainerType string `xml:"ContainerType,attr"`
	Puid          string `xml:"Puid,attr"`
}

// CreateTriggerPUIDS ...
func (signatureFile *ContainerSignatureMapping) CreateTriggerPUIDS() {
	triggerPUIDContainer := make([]TriggerPuidContainer, numberOfTriggers)
	triggerPUIDContainer1 := TriggerPuidContainer{}
	triggerPUIDContainer2 := TriggerPuidContainer{}
	triggerPUIDContainer3 := TriggerPuidContainer{}

	triggerPUIDContainer1.ContainerType = triggerOLE2
	triggerPUIDContainer1.Puid = triggerOLE2PUID

	triggerPUIDContainer2.ContainerType = triggerZIP1
	triggerPUIDContainer2.Puid = triggerZIP1PUID

	triggerPUIDContainer3.ContainerType = triggerZIP2
	triggerPUIDContainer3.Puid = triggerZIP2PUID

	triggerPUIDContainer[0] = triggerPUIDContainer1
	triggerPUIDContainer[1] = triggerPUIDContainer2
	triggerPUIDContainer[2] = triggerPUIDContainer3

	signatureFile.TriggerPuids.TriggerPUID = triggerPUIDContainer
}

// String returns the signature file as a string.
func (signatureFile ContainerSignatureMapping) String() string {
	signatureFile.CreateTriggerPUIDS()
	out, _ := xml.MarshalIndent(signatureFile, "", " ")
	return fmt.Sprintf("%s%s", xml.Header, out)
}
