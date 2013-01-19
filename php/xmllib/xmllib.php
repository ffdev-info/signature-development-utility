<?php

//Lightweight wrapper for PHP DOM Libraries to be added to as required.
//Currently used by web service handling scripts to generate xml.

//TODO: Could / should (?) optimise by removing $doc option from half these 
//		  functions and placing in a member variable. 

class XMLHelper
{
	/*********************************************************************
	*
	* Function		: newDOMDocument
	*
	* Description	: create and return a new DOM Object
	*
	* parameters	: n/a	
	*          
	* Returns		: $doc - DOM object / document
	*
	*********************************************************************/
	public function newDOMDocument()
	{
		$doc = new DOMDocument();
		$doc->formatOutput = true;
		$doc->encoding='UTF-8';
		return $doc;
	}

	/*********************************************************************
	*
	* Function		: xmlCreateRootElement
	*
	* Description	: create root element and append to given DOM document
	*
	* parameters	: $name - name of element
	*					  $namespace - optional namespace
	*					  $doc - DOM document
	*					  $parent - parent element to append new element
	*          
	* Returns		: $element - xml root element
	*
	*********************************************************************/
	public function xmlCreateRootElement($name, $namespace, $doc, $parent)
	{
		if($namespace === null)
			$element = $doc->createElement($name);
		else
			$element = $doc->createElementNS($namespace, $name);
		
		$parent->appendChild($element);	
		return $element;	
	}

	/*********************************************************************
	*
	* Function		: xmlCreateElement
	*
	* Description	: create a vanilla element and append to given parent node
	*
	* parameters	: $name - name of element
	*					  $doc - DOM document
	*					  $parent - parent element to append new element
	*          
	* Returns		: new xml node
	*
	*********************************************************************/
	public function xmlCreateElement($name, $doc, $parent)
	{
		$element = $doc->createElement($name); 
		$parent->appendChild($element);
		return $element;
	}
	
	/*********************************************************************
	*
	* Function		: xmlCreateTextElement
	*
	* Description	: create a new xml element containing text
	*
	* parameters	: $name - name of element
	*					  $value - value of text to add to element
	*					  $doc - DOM document
	*					  $parent - parent element to append new element
	*          
	* Returns		: new xml node
	*
	*********************************************************************/
	public function xmlCreateTextElement($name, $value, $doc, $parent)
	{
		$element = $doc->createElement($name); 
		$parent->appendChild($element);
		$this->xmlAddTextValue($value, $doc, $element);
		return $element;
	}	
	
	/*********************************************************************
	*
	* Function		: xmlAddAttribute
	*
	* Description	: create xml attribute and add to given node
	*
	* parameters	: $name - name of attribute
	*					  $value - value to give attribute
	*					  $doc - DOM document
	*					  $parent - parent element to append new attribute
	*          
	* Returns		: n/a
	*
	*********************************************************************/
	public function xmlAddAttribute($name, $value, $doc, $parent)
	{
		$attribute = $doc->createAttribute($name); 
		$this->xmlAddTextValue($value, $doc, $attribute);
		$parent->appendChild($attribute); 	
	}
	
	/*********************************************************************
	*
	* Function		: xmlAddTextValue
	*
	* Description	: public function to add text values to nodes or attributes
	*					  wraps the function to be used inside 'this' class or 
	*					  externally if required. 
	*
	* parameters	: $value - text value for element or attribute
	*					  $doc - DOM document
	*					  $parent - parent element to append value
	*          
	* Returns		: n/a
	*
	*********************************************************************/
	//add a text value to xml node or attribute. used inside this 
	//class for attributes or externally when working on the DOM...
	public function xmlAddTextValue($value, $doc, $parent)
	{
		$textValue = $doc->createTextNode($value);
		$parent->appendChild($textValue);
	}
	
	/*********************************************************************
	*
	* Function		: xmlOutput
	*
	* Description	: convert DOM to XML document and output to broweser.
	*					  set http header to allow browser to handle correctly.
	*
	* parameters	: $doc - DOM document
	*          
	* Returns		: outputs xml to browser
	*
	*********************************************************************/
	public function xmlOutput($doc)
	{
	  header ('Content-Type: text/xml');
	  print $doc->saveXML();		
	}

	/*********************************************************************
	*
	* Function		: combine_xml
	*
	* Description	: combine two DOM documents and return the new XML document
	*
	* parameters	: $xml1 - primary XML to concatenate with
	*					  $xml2 - secondary XML to be appended
	*					  $xml1_parent - primary node to afix secondary DOM to
	*					  $xml2_node - secondary node from which to afix DOM
	*          
	* Returns		: new XML document
	*
	*********************************************************************/
	public function combine_xml($xml1, $xml2, $xml1_parent, $xml2_node)
	{
		//TODO: Additional work needed here to solidify function a bit better
		//combining two DOM objects is a fairly precarious operation depending
		//a lot of calling class/function. Add better error handling. Fix variable
		//names.
	
		//TODO: assumption made here that we have a DOM object. Need better checking
		if(!is_object($xml1))
		{
			$doc1 = new DOMDocument();
			$doc1->loadXML($xml1);
			$doc1->formatOutput = true;
		}
		else
		{
			$doc1 = $xml1;
		}
		
		if(!is_object($xml2))
		{
			$doc2 = new DOMDocument();
			$doc2->loadXML($xml2);
			$doc2->formatOutput = true;
		}
		else
		{
			$doc2 = $xml2;
		}

		$node = $doc2->getElementsByTagName($xml2_node)->item(0);
		$new_node = $doc1->importNode($node, true);
		$doc1->getElementsByTagName($xml1_parent)->item(0)->appendChild($new_node);
	
		//TODO: return DOM or XML... function parameter (?)
		return $doc1->saveXML();
	}	
}

?>