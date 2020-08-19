// Fixtures for testing the signature output capabilities of this
// package.

package sigdevutil

// A basic fixture representing a standard signature file that works
// with DROID. See also: Hammer and nail. It might be more elegant to
// not have an entire chunk of XML, but also, this works for what we
// need for now.
const standardSignatureOne = `<?xml version="1.0" encoding="UTF-8"?>
<FFSignatureFile xmlns="http://www.nationalarchives.gov.uk/pronom/SignatureFile" Version="1" DateCreated="1970-01-01T00:00:00">
 <InternalSignatureCollection>
  <InternalSignature ID="1" Specificity="Specific">
   <ByteSequence Reference="BOFoffset">
    <SubSequence Position="1" MinFragLength="0" SubSeqMinOffset="0" SubSeqMaxOffset="0">
     <Sequence>1234</Sequence>
    </SubSequence>
   </ByteSequence>
   <ByteSequence Reference="EOFoffset">
    <SubSequence Position="1" MinFragLength="0" SubSeqMinOffset="0" SubSeqMaxOffset="0">
     <Sequence>3456</Sequence>
    </SubSequence>
   </ByteSequence>
   <ByteSequence>
    <SubSequence Position="1" MinFragLength="0" SubSeqMinOffset="0">
     <Sequence>7890</Sequence>
    </SubSequence>
   </ByteSequence>
  </InternalSignature>
 </InternalSignatureCollection>
 <FileFormatCollection>
  <FileFormat ID="1" Name="test format" PUID="fmt/123" Version="1.0" MIMEType="application/octet-stream">
   <InternalSignatureID>1</InternalSignatureID>
   <Extension>ext</Extension>
  </FileFormat>
 </FileFormatCollection>
</FFSignatureFile>`

// A basic container signature for fmt/39 which should work with DROID
// and has a single byte sequence and two paths, so some okay
// complexity.
const containerSignatureOne = `<?xml version="1.0" encoding="UTF-8"?>
<ContainerSignatureMapping SchemaVersion="1.0" SignatureVersion="1">
 <ContainerSignatures>
  <ContainerSignature Id="1" ContainerType="OLE2">
   <Description>test format</Description>
   <Files>
    <File>
     <Path>WordDocument</Path>
    </File>
    <File>
     <Path>CompObj</Path>
     <BinarySignatures>
      <InternalSignatureCollection>
       <InternalSignature ID="1">
        <ByteSequence Reference="BOFoffset">
         <SubSequence Position="1" SubSeqMinOffset="40" SubSeqMaxOffset="1024">
          <Sequence>10000000576F72642E446F63756D656E742E</Sequence>
         </SubSequence>
        </ByteSequence>
       </InternalSignature>
      </InternalSignatureCollection>
     </BinarySignatures>
    </File>
   </Files>
  </ContainerSignature>
 </ContainerSignatures>
 <FileFormatMappings>
  <FileFormatMapping signatureId="1" Puid="fmt/39"></FileFormatMapping>
 </FileFormatMappings>
 <TriggerPuids>
  <TriggerPuid ContainerType="OLE2" Puid="fmt/111"></TriggerPuid>
  <TriggerPuid ContainerType="ZIP" Puid="fmt/189"></TriggerPuid>
  <TriggerPuid ContainerType="ZIP" Puid="x-fmt/263"></TriggerPuid>
 </TriggerPuids>
</ContainerSignatureMapping>`
