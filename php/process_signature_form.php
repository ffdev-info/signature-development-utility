<?php 

	include_once("generatebytecode/generatebytecode.php");
	
	//uuid lib from: http://jkingweb.ca/code/php/lib.uuid/
	include_once("uuid/lib.uuid.php");

	if(isset($_POST['rdf']) == true)
	{
		generateRDF();
	}
	else
	{
		$counter = $_POST['counter'];
		$byte_sequences = generateSignatureCollection($counter);
		$file_formats = generateFormatCollection(1);
		generateSignatureFile($byte_sequences, $file_formats);
	}
	
	
	function generateRDF()
	{
		$xml = new XMLHelper();
		
		$doc = $xml->newDOMDocument();
		
		$rdf = $xml->xmlCreateRootElement('rdf:RDF', null, $doc, $doc);
		$xml->xmlAddAttribute('xmlns:xsd', 'http://www.w3.org/2001/XMLSchema#', $doc, $rdf);
		$xml->xmlAddAttribute('xmlns:sigdev', 'http://nationalarchives.gov.uk/preservation/sigdev/signature/', $doc, $rdf);
		$xml->xmlAddAttribute('xmlns:bytes', 'http://nationalarchives.gov.uk/preservation/sigdev/signature/byteSequence/', $doc, $rdf);
		$xml->xmlAddAttribute('xmlns:rdf', 'http://www.w3.org/1999/02/22-rdf-syntax-ns#', $doc, $rdf);
		$xml->xmlAddAttribute('xmlns:rdfs', 'http://www.w3.org/2000/01/rdf-schema#', $doc, $rdf);

		$sigdev = $xml->xmlCreateElement('sigdev:DevelopmentSignature', $doc, $rdf);
		$xml->xmlAddAttribute('rdf:about', 'http://nationalarchives.gov.uk/preservation/sigdev/signature/' . UUID::mint( 4 ), $doc, $sigdev);
		
		$xml->xmlCreateTextElement('rdfs:label', $_POST['name1'], $doc, $sigdev);
		$xml->xmlCreateTextElement('sigdev:version', $_POST['version1'], $doc, $sigdev);
		$xml->xmlCreateTextElement('sigdev:extension', $_POST['extension1'], $doc, $sigdev);
		$xml->xmlCreateTextElement('sigdev:internetMediaType', $_POST['mimetype1'], $doc, $sigdev);
		$xml->xmlCreateTextElement('sigdev:puid', $_POST['puid1'], $doc, $sigdev);
		
		$count = $_POST['counter'];
		
		$uuid_var = UUID::mint( 4 );
		
		for($x = 1; $x <= $count; $x++)
		{
			$sequence_url = 'http://nationalarchives.gov.uk/preservation/sigdev/signature/byteSequence/' . $uuid_var . '/' . $x;
			
			$byteseq = $xml->xmlCreateElement('sigdev:byteSequence', $doc, $sigdev);
			$seq = $xml->xmlCreateElement('rdf:Description', $doc, $byteseq);
			$xml->xmlAddAttribute('rdf:about', $sequence_url, $doc, $seq);
			
			$str = $xml->xmlCreateTextElement('bytes:string', $_POST['signature'.$x], $doc, $seq);
			$xml->xmlAddAttribute('rdf:datatype', 'http://nationalarchives.gov.uk/preservation/sigdev/signature/droidRegularExpression', $doc, $str);
			
			$xml->xmlCreateTextElement('bytes:anchor', $_POST['anchor'.$x], $doc, $seq);
		
			$offset = $xml->xmlCreateTextElement('bytes:offset', $_POST['offset'.$x], $doc, $seq);
			$xml->xmlAddAttribute('rdf:datatype', 'http://www.w3.org/2001/XMLSchema#integer', $doc, $offset);

			$maxOffset = $xml->xmlCreateTextElement('bytes:maxOffset', $_POST['maxoffset'.$x], $doc, $seq);
			$xml->xmlAddAttribute('rdf:datatype', 'http://www.w3.org/2001/XMLSchema#integer', $doc, $maxOffset);
		}
		
		$filename = 'Content-disposition: attachment; filename=' . str_replace(' ', '-', $_POST[name1]) . '-v' . str_replace(' ', '-', $_POST[version1]) . '.rdf';
		
		header($filename);
		header ('Content-Type: text/xml');		
		print $doc->saveXML();	
	}
	
	
	/*********************************************************************
	*
	* Function		:  
	*
	* Description	:  description...
	*
	* parameters	: 	
	*          
	* Returns		:  n/a
	*
	*********************************************************************/
	function generateSignatureCollection($count)
	{
		$gen = new SignatureGenerator();	
		$signature_collection = new SignatureCollection();

		for($i = 1; $i < $count + 1; $i++)	
		{
			$new_sig = 'signature' . $i;
			$new_anc = 'anchor' . $i;
			$new_off = 'offset' . $i;
			$new_max = 'maxoffset' . $i;
			
			$byte_sequence = new ByteSequence();
			
			$signature_collection->sig_id = 1;
			$signature_collection->specificity = 'Specific';
			
			$byte_sequence->position = $_POST[$new_anc];
			$byte_sequence->offset = $_POST[$new_off];
			$byte_sequence->maxoffset = $_POST[$new_max];
			$byte_sequence->byte_sequence = $_POST[$new_sig];
			
			$signature_collection->byteSequence[] = $byte_sequence;
		}
	
		return $gen->generateSignatureFromObject($signature_collection);
	}
	
	/*********************************************************************
	*
	* Function		:  
	*
	* Description	:  description...
	*
	* parameters	: 	
	*          
	* Returns		:  n/a
	*
	*********************************************************************/
	function generateFormatCollection($i)
	{
		$new_name 		= 'name' . $i;
		$new_version 	= 'version' . $i;
		$new_puid 		= 'puid' . $i;
		$new_mime 		= 'mimetype' . $i;
		$new_ext 		= 'extension' . $i;
		
		$xml = new XMLHelper();
		
		$collection = $xml->newDOMDocument();
		
		$ffc = $xml->xmlCreateElement('FileFormatCollection', $collection, $collection);
		$ff =  $xml->xmlCreateElement('FileFormat', $collection, $ffc);
		
		$xml->xmlAddAttribute('ID', $i, $collection, $ff);
		$xml->xmlAddAttribute('Name', $_POST[$new_name], $collection, $ff);
				
		if(isset($_POST[$new_puid]) == true)
		{
			$xml->xmlAddAttribute('PUID', $_POST[$new_puid], $collection, $ff);
		}
		else
		{
			$test_puid = substr($_POST[$new_name], 0, 3) . '/' . $i;
			$xml->xmlAddAttribute('PUID', $test_puid, $collection, $ff);		
		}
		
		if(isset($_POST[$new_version]) == true)
		{
			$xml->xmlAddAttribute('Version', $_POST[$new_version], $collection, $ff);
		}
		
		if(isset($_POST[$new_mime]) == true)
		{
			$xml->xmlAddAttribute('MIMEType', $_POST[$new_mime], $collection, $ff);
		}	
		
		$sig =  $xml->xmlCreateElement('InternalSignatureID', $collection, $ff);		
		$xml->xmlAddTextValue($i, $collection, $sig);
		
		if(isset($_POST[$new_ext]) == true)
		{
			$sig =  $xml->xmlCreateElement('Extension', $collection, $ff);		
			$xml->xmlAddTextValue($_POST[$new_ext], $collection, $sig);
		}
		
		//header ('Content-Type: text/xml');	
		return $collection->saveXML();
	}
	
	/*********************************************************************
	*
	* Function		:  
	*
	* Description	:  description...
	*
	* parameters	: 	
	*          
	* Returns		:  n/a
	*
	*********************************************************************/
	function generateSignatureFile($byte_sequences, $file_formats)
	{
		$xml = new XMLHelper();
		$collection = $xml->newDOMDocument();

		$xml->xmlCreateElement('InternalSignatureCollection', $collection, $collection);

		$collection = $collection->saveXML();

		$collection = $xml->combine_xml($collection, $byte_sequences, 'InternalSignatureCollection', 'InternalSignature');

		//Create the FFSignatureFile Portion of the document
		$xmlns = 'http://www.nationalarchives.gov.uk/pronom/SignatureFile';
		
		// Current date/time in your computer's time zone.
		$date = new DateTime();
		$xml = new XMLHelper();
		$doc1 = $xml->newDOMDocument();

		$signatureFile = $xml->xmlCreateElement('FFSignatureFile', $doc1, $doc1);
		$xml->xmlAddAttribute('Version', '1', $doc1, $signatureFile);
		$xml->xmlAddAttribute('xmlns', $xmlns, $doc1, $signatureFile);

		$xml->xmlAddAttribute('DateCreated', $date->format(DateTime::W3C), $doc1, $signatureFile);	

		$doc1 = $doc1->saveXML();		

		$doc1 = $xml->combine_xml($doc1, $collection, 'FFSignatureFile', 'InternalSignatureCollection');
		$doc1 = $xml->combine_xml($doc1, $file_formats, 'FFSignatureFile', 'FileFormatCollection');	

		$filename = 'Content-disposition: attachment; filename=' . str_replace(' ', '-', $_POST[name1]) . '-v' . str_replace(' ', '-', $_POST[version1]) . '-signature-file.xml';
				
		header($filename);
		header ('Content-Type: text/xml');		
		print $doc1;			
	}
?>
