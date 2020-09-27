<?php 
		
include_once("xmllib/xmllib.php");

class SignatureCollection
{
	public $sig_id;
	public $specificity = 'Specific';
	public $byteSequence = array();
}

class ByteSequence
{
	public $position;
	public $offset = 0;
	public $maxoffset = 0;
	public $byte_sequence;	
}

class SignatureGenerator
{
	// property declaration
	private $byte_len 	= 2;
	private $bofMax 		= 0;
	private $eofMax 		= 0;

	//maintain an array of values for subsequence
	private $suboffs = array();
	
	private $minFragLen 	= 0;
	private $eof	 		= false;
	private $varoff		= false;

	/*********************************************************************
	*
	* Function		:  output_xml($doc)
	*
	* Description	:  description...
	*
	* parameters	: 	$doc pointer to xml document
	*          
	* Returns		:  n/a
	*
	*********************************************************************/
	private function output_xml($doc)
	{
	  //output the xml...
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
	public function generateSignatureFromObject($signature_collection) 
	{
		$xml = new XMLHelper();
		unset($doc);

		$doc = $xml->newDOMDocument();
		$is = $xml->xmlCreateElement('InternalSignature', $doc, $doc);

		$xml->xmlAddAttribute('ID', $signature_collection->sig_id, $doc, $is);	
		$xml->xmlAddAttribute('Specificity', $signature_collection->specificity, $doc, $is);	

		for($x = 0; $x < sizeof($signature_collection->byteSequence); $x++)
		{
			unset($this->suboffs);
			$this->bofMax 		= 0;
			$this->eofMax 		= 0;
		
			$byte = $signature_collection->byteSequence[$x];

			//most basic thing we can do to ease validation...
			//always ensure we have an upper case string to work with...
			$string = strtoupper($byte->byte_sequence);

			$anchor = $byte->position;
			$offset = $byte->offset;
			$max = $byte->maxoffset;

			$anchor === 'Variable'  ? $this->varoff = true : $this->varoff = false;
			$anchor === 'EOFoffset' ? $this->eof = true : $this->eof = false;

			$subsequences = $this->splitIntoSubsequences($string);

			//compute subsequence stuff beforehand...
			for ($i = 0; $i < sizeof($subsequences); $i++)
			{
				($i == (sizeof($subsequences)-1)) ? $bLast = true : $bLast = false;
				($i == 0) ? $bFirst = true : $bFirst = false;

				//Once we've got our subsequences below we can output the xml nodes for
				//the values we've gathered...
				$subsequences[$i] = $this->stripCurlyWildcards($subsequences[$i], $bFirst, $bLast);		
			}

			$len_offset_array = sizeof($this->suboffs);	
			for($head=1; $head < $len_offset_array-1; $head+=2)		
			{
				//get the value from the head of the next subsequence
				//and store in a temporary variable to add to the tail of the current
				if($head != ($len_offset_array-2))
				{
					$tmp = $this->suboffs[$head+1];
					unset($this->suboffs[$head+1]);		//unset head from array to contract values
				}
		
				//new tail... add head value...
				if($head != 0 && $head != ($len_offset_array-1))
				{
					$this->suboffs[$head] = $this->suboffs[$head] += $tmp;
				}
			}			

			$this->suboffs = array_values($this->suboffs);
			
			$this->suboffs[0] = $this->suboffs[0] += $offset;
			$this->suboffs[sizeof($this->suboffs)-1] = $this->suboffs[sizeof($this->suboffs)-1] += $offset;

			$this->bofMax += $max;
			$this->eofMax += $max;
			
			if($this->eof) 
			{ 
				unset($this->suboffs[0]);
				$this->suboffs = array_values($this->suboffs);
				$this->eofMax += end($this->suboffs);
			}
			else
			{
				$this->bofMax += $this->suboffs[0];		
			}

			//need to do some work to bring some of this code together... member variables?
			$doc = $this->bringXMLTogether($doc, $is, $subsequences, $anchor, $offset, $max, $xml);

		}

		return $doc->saveXML();
	}

	/*********************************************************************
	*
	* Function		:  
	*
	* Description	:  
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	private function bringXMLTogether($doc, $is, $subsequences, $anchor, $offset, $max, $xml)
	{
		$byteSeq = $xml->xmlCreateElement('ByteSequence', $doc, $is);
		
		if($this->varoff === false)
		{
			$xml->xmlAddAttribute('Reference', $anchor, $doc, $byteSeq);	
		}

		for($i = 0; $i < sizeof($subsequences); $i++)
		{
			//Within subsequence tags is the breakdown of the rest of the signatue...
			//longest sequence, left right fragments etc. output using the functions below...
			$subSeq = $xml->xmlCreateElement('SubSequence', $doc, $byteSeq);
		
			$fragment_longest_pair = $this->longestUnambiguousSequence($subsequences[$i]);
			$this->handleStringFragments($fragment_longest_pair, $doc, $subSeq, $xml);

			//use outputFragPos() via handleStringFragments to increment
			//MinFragLength and output here...
			$xml->xmlAddAttribute('MinFragLength', $this->minFragLen, $doc, $subSeq);
			$this->Len = 0;
			//reset minfraglength to zero for next fragment...
			$this->minFragLen = 0;
		
			$xml->xmlAddAttribute('Position', $i+1, $doc, $subSeq);		

			if($this->varoff === false)
			{
				if ($this->eof === false && $i === 0)	//i.e. if $i == 0 (1)
				{
					$xml->xmlAddAttribute('SubSeqMaxOffset', $this->bofMax, $doc, $subSeq);
				}
				elseif ($this->eof === true && $i === sizeof($subsequences)-1)
				{
					$xml->xmlAddAttribute('SubSeqMaxOffset', $this->eofMax, $doc, $subSeq);
				}
			}
						
			$xml->xmlAddAttribute('SubSeqMinOffset', $this->suboffs[$i], $doc, $subSeq);

		}

		return $doc;
	}

	/*********************************************************************
	*
	* Function		:  
	*
	* Description	:  
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	private function splitIntoSubsequences($string)
	{
		$subsequences = array();
		$asterisk_count = substr_count($string, '*');
		
		if ($asterisk_count)
		{
			for($i = 0; $i < $asterisk_count; $i++)
			{
				$pos = strpos($string,'*');
				
				
				$bracket_test = substr($string, $pos, 2);

				//if asterisk appears within a bracket split but handle differently
				if(strpos($bracket_test, '}'))
				{
					$subsequences[] = substr($string, 0, $pos+2);
					$string = substr($string, $pos+2);		
				}
				else	//we can simply split the string as in normal circumstances
				{
					$subsequences[] = substr($string, 0, $pos);
					$string = substr($string, $pos+1);
				}
			}
		}

		//should have one fragment left from loop, add to array...
		//or put frag in here if no wildcards...
		$subsequences[] = $string;
		
		return $subsequences;
	}
	

	/*********************************************************************
	*
	* Function		:  
	*
	* Description	:  
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	private function stripCurlyWildcards($string, $bFirst, $bLast)
	{
		//will only ever have two values here bos / eos
		$start = 0;		// start of sequence wildcard
		$end = 0;		// end of sequence wildcard

		//array indexes. assign var names for ease of reading
		if(!defined("MIN_")) { define("MIN_", 0); }
		if(!defined("MAX_")) { define("MAX_", 1); }

		$s_pos = strpos($string, '{');
		if(!is_bool($s_pos))
		{
			if($s_pos == 0)
			{
				$len = strpos($string, '}') + 1;
				$start = substr($string, 0, $len);
				$string = substr($string, $len);
			}
		}

		//reverse string, look for first occurrence
		//of closing curly bracket '}'...
		$string = strrev($string);
		
		$s_pos = strpos($string, '}');
		if(is_bool($s_pos) != true)
		{		
			if($s_pos == 0)
			{
				$len = strpos($string, '{') + 1;
				$end = strrev(substr($string, 0, $len));
				$string = substr($string, $len);
			}
		}

		//reverse string finally to correct its orientation
		$string = strrev($string);

		//Strip wildcard vals of hypens and record value pair
		if ($start) 
		{ 
			$optarr1 = $this->getBracketedValues('-', $start);
						
			if(sizeof($optarr1) == 2)
			{
				if($optarr1[MAX_] == '*')
				{
					if($bFirst)		//for BOF sequences, if we have * it is variable...
					{
						$this->varoff = true;
					}
					$optarr1[MAX_] = 0;		//means little, only affects BOF or EOF if we have a MAXOFFSET attribute
				}
			}
			
			//add the difference between the two values to max offset, not the entire value again...
			if($bFirst && sizeof($optarr1) == 2) { $this->bofMax += ($optarr1[MAX_]-$optarr1[MIN_]); }
			if ($optarr1[MIN_]) { $this->suboffs[] = $optarr1[MIN_]; }
		}
		else
		{
			$this->suboffs[] = 0;
		}
		
		if ($end) 
		{ 
			$optarr2 = $this->getBracketedValues('-', $end);

			if(sizeof($optarr2) == 2)
			{			
				if($optarr2[MAX_] == '*')
				{
					if($bLast)		//for EOF sequences, if we have * it is variable...
					{
						$this->varoff = true;
					}
					$optarr2[MAX_] = 0;		//means little, only affects BOF or EOF if we have a MAXOFFSET attribute
				}
			}

			//add the difference between the two values to max offset, not the entire value again...
			if($bLast && sizeof($optarr2) == 2) { $this->eofMax += ($optarr2[MAX_]-$optarr2[MIN_]); }
			if ($optarr2[MIN_]) { $this->suboffs[] = $optarr2[MIN_]; }
		}
		else
		{
			$this->suboffs[] = 0;
		}

		return $string;
	}

	/*********************************************************************
	*
	* Function		:  
	*
	* Description	:  given a string find longest unambiguous sequence
	*						i.e. longest sequence that doesn't contain syntax
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	private function longestUnambiguousSequence($string)
	{
		$fragments = array();
		$all_fragments = array();
		$len = strlen($string);

		for($i = 0; $i < $len; $i++)
		{
			$frag_stored = false;

			$char = $string[$i];
		
			// ?? {n} {k-m} (a|b) [!a:b]	
			if ($char == '?' || $char == '{'
				|| $char == '(' || $char == '[')
			{				
				//if array is empty can take first substring as-is
				if (sizeof($fragments) == 0)
				{
					$fragments[] = substr($string, 0, $i);
					$all_fragments[] = substr($string, 0, $i);
					$string = substr($string, $i);
					$this->setStrLen($len, $i, $string);
				}
				else
				{
					if ($i != 0)
					{
						//should work, but need more data to test...
						$fragments[] = substr($string, 0, $i);
						$all_fragments[] = substr($string, 0, $i);						
					}

					switch($char)
					{
						case '?':
							//put wildcard into secondary array...
							$all_fragments[] = substr($string, $i, 2);
							$string = substr($string, $i+2);
							$this->setStrLen($len, $i, $string);
							break;

						case '[':
							//put wildcard into secondary array...
							$all_fragments[] = substr($string, $i, (strpos($string, ']') - $i) + 1);
							$string = substr($string, strpos($string, ']')+1);
							$this->setStrLen($len, $i, $string);
							break;

						case '{':
							$all_fragments[] = substr($string, $i, (strpos($string, '}') - $i) + 1);
							$string = substr($string, strpos($string, '}')+1);
							$this->setStrLen($len, $i, $string);
							break;				
						
						case '(':
							$all_fragments[] = substr($string, $i, (strpos($string, ')') - $i) + 1);
							$string = substr($string, strpos($string, ')')+1);
							$this->setStrLen($len, $i, $string);	
							break;			
					}
				}
			}
		}

		if ($string)
		{
			$fragments[] = $string;
			$all_fragments[] = $string;
		}
		
		$maxlen = max(array_map('strlen', $fragments));

		for($i = 0; $i < sizeof($fragments); $i++)
		{
			if(strlen($fragments[$i]) == $maxlen)
			{
				$longest_frag = $fragments[$i];
				$longest_index = $i;
				break;
			}
		}
		
		$frag_return = array();
		$frag_return[] = $longest_frag;
		$frag_return[] = $all_fragments;
		return $frag_return;
	}

	/*********************************************************************
	*
	* Function		:  
	*
	* Description	: 
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	//reset iterator and string count as required...
	private function setStrLen(&$len, &$i, $string)
	{
		$len = strlen($string);
		$i = -1;							//$i = -1 set to zero when enters loop
	}

	/*********************************************************************
	*
	* Function		:  
	*
	* Description	: 
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	//handle the different signature string fragments belonging to a single
	//subsequence. output xml as appropriate as we cycle through the string components.
	private function handleStringFragments($fragment_longest_pair, $doc, $parent, $xml)
	{
		$longest_frag = $fragment_longest_pair[0];
		$fragments = $fragment_longest_pair[1];

		$seq = $xml->xmlCreateElement('Sequence', $doc, $parent);
		$xml->xmlAddTextValue($longest_frag, $doc, $seq);
		
		$longestLen = strlen($longest_frag) / $this->byte_len;
		
		//default and value for BOF sequences...
		$shiftVal = $longestLen; 
		
		$seq = $xml->xmlCreateElement('DefaultShift', $doc, $parent);

		if($this->eof)
		{
			$shiftVal = -1;
			$xml->xmlAddTextValue(-($longestLen+1), $doc, $seq);	
		}
		else
		{
			$xml->xmlAddTextValue($longestLen+1, $doc, $seq);		
		}

		$uniqueByte = array();

		for ($i = 0; $i < strlen($longest_frag); $i+=$this->byte_len)
		//for ($i = strlen($longest_frag)-2; $i >= 0; $i-=$this->byte_len)
		{
			$bExists = false;
			$byte = substr($longest_frag, $i, $this->byte_len);
	
			for ($j = 0; $j < sizeof($uniqueByte); $j++)
			{
				if($byte === $uniqueByte[$j][0])
				{
					$bExists = true;
					if(!$this->eof) { $uniqueByte[$j][1] = $shiftVal; }	//update to get distance from end
					break;
				}
			}
			
			if ($bExists === false)
			{
				$uniqueByte[] = array($byte, $shiftVal);
			}

			$shiftVal--;
		}

		for ($j = 0; $j < sizeof($uniqueByte); $j++)
		{
			//loop to output xml elements
			$shift = $xml->xmlCreateElement('Shift', $doc, $parent);
			$xml->xmlAddTextValue($uniqueByte[$j][1], $doc, $shift);
			$xml->xmlAddAttribute('Byte', $uniqueByte[$j][0], $doc, $shift);	
		}

		//arrarys for left and right fragments...
		$left = array();
		$right = array();

		$first = &$left;

		for ($i = 0; $i < sizeof($fragments); $i++)
		{
			if ($fragments[$i] === $longest_frag)
			{
				$first = &$right;
			}
			else
			{
				$first[] = $fragments[$i];
			}
		}
	
		if ($left)  { $this->outputFragPos(true, $left, $doc, $parent, $xml); }
		if ($right) { $this->outputFragPos(false, $right, $doc, $parent, $xml); }

	}
	
	/*********************************************************************
	*
	* Function		:  
	*
	* Description	: 
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	private function recombineArray($fragments)
	{
		//array to bring together multiple parts of same fragment
		//such as those parts of array containing square brackets
		$recombine_arr = array();
	
		$string = '';
	
		for ($i = 0; $i < sizeof($fragments); $i++)
		{
			$char = $fragments[$i][0];

			if($char != '?' && $char != '{' && $char != '(')
			{
				$string = $string . $fragments[$i];
			}
			else
			{
				if($string != '')
					$recombine_arr[] = $string;	
				
				$string = '';
				$recombine_arr[] = $fragments[$i];
			}
		}
	
		//Add string remainder to array
		if($string != '')
			$recombine_arr[] = $string;	
	
		return $recombine_arr;
	}
	
	/*********************************************************************
	*
	* Function		:  
	*
	* Description	: 
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/	
	//output correct stuff for left or right fragments...
	private function outputFragPos($bLeft, $fragments, $doc, $parent, $xml)
	{	
		$fragments = $this->recombineArray($fragments);

		//print_r($fragments);

		$orientation = ($bLeft) ? "LeftFragment" : "RightFragment";

		//if bLeft need to work through array backwards...
		if($bLeft)
		{
			$fragments = array_reverse($fragments, false);
		}
		
		$fraglen = 0;
		//TODO: if bLeft and BOFSequence then min frag calculated from beginning
		//TODO: if !bLeft and EOFSequence then we calculate from the end minfraglength
	
		$string = '';	
		$fragPos = 0;
		
		$minoffset = 0;		//output and reset every time we output string.
		$maxoffset = 0;

		for ($i = 0; $i < sizeof($fragments); $i++)
		{
			//charAt(0) will always be syntactical character
			//check to see if this matches anything but '['
			$char = substr($fragments[$i], 0, 1);

			switch ($char)
			{
				case '?':
					//offset count + one byte... output string.
					
					//potential to have two ?? or more after one another.
					if(strlen($string) > 0)
					{
						$this->fragXMLOutput($orientation, $string, $fragPos+1, $minoffset, $maxoffset, $doc, $parent, $xml);
						$fraglen = $this->handleFragLength($fraglen, $minoffset, $string);
						$fragPos += 1;
						$minoffset = 0;
						$maxoffset = 0;
					}

					$minoffset += 1;
					$maxoffset += 1;
					$string = '';
					break; 
					
				case '{':
					//output $string var and...
					//values in curly brackets become offset values...
					
					if(strlen($string) > 0)
					{
						$this->fragXMLOutput($orientation, $string, $fragPos+1, $minoffset, $maxoffset, $doc, $parent, $xml);
						$fraglen = $this->handleFragLength($fraglen, $minoffset, $string);
						$fragPos += 1;
						$minoffset = 0;
						$maxoffset = 0;
					}
	
					$optarr = $this->getBracketedValues('-', $fragments[$i]);

					if(sizeof($optarr) == 2)
					{
						$minoffset = $minoffset + $optarr[0];
						$maxoffset = $maxoffset + $optarr[1];
					}
					elseif (sizeof($optarr) == 1)
					{
						$minoffset = $minoffset + $optarr[0];
						$maxoffset = $maxoffset + $optarr[0];
					}
					//else we've a bum array... not sure what to do.
					
					$string = '';
					break;
					
				case '(':
					//output $string var and...
					//output current array string into each of the potential options...
					if(strlen($string) > 0)
					{
						$this->fragXMLOutput($orientation, $string, $fragPos+1, $minoffset, $maxoffset, $doc, $parent, $xml);
						$fraglen = $this->handleFragLength($fraglen, $minoffset, $string);						
						$fragPos += 1;
						$minoffset = 0;
						$maxoffset = 0;
					}
					
					$optarr = $this->getBracketedValues('|', $fragments[$i]);

					//minimum length of stirng in options array to increment min $fraglength
					//handle differently as we don't want $fraglength to increase for each option
					$minlen = (min(array_map('strlen', $optarr))) / 2;

					//output identical nodes for both options, e.g. if pos == 6, pos == 6 for both of these
					for($j = 0; $j < sizeof($optarr); $j++)
					{
						$this->fragXMLOutput($orientation, $optarr[$j], $fragPos+1, $minoffset, $maxoffset, $doc, $parent, $xml);
					}
					$fraglen += $minlen;
					$fragPos += 1;	
					$minoffset = 0;
					$maxoffset = 0;
					
					$string = '';	
					break;
			
				default:
					$string = $string . $fragments[$i];
					break;
			}
		}
		
		if(strlen($string) > 0)
		{
			$this->fragXMLOutput($orientation, $string, $fragPos+1, $minoffset, $maxoffset, $doc, $parent, $xml);
			$fragPos += 1;
			$fraglen = $this->handleFragLength($fraglen, $minoffset, $string);
		}
		
		//output the fragment length count here as required...
		if (!$this->eof && $bLeft)
		{
			$this->minFragLen = $fraglen;
		}
		elseif ($this->eof && $bLeft == false)
		{
			$this->minFragLen = $fraglen;		
		}
	}

	/*********************************************************************
	*
	* Function		:  
	*
	* Description	: 
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	//count the number of characters output when we output $string
	private function handleFragLength($fraglen, $minoffset, $string)
	{
		$fraglen = $fraglen + $minoffset;
		
		//for each set of bracketed values remove len '[a4:a5]' minus two
		//e.g. [a4:a5] equals seven in length. remove five and we are left
		//with two. this divided by two equals one byte...
		//len to remove:
		$n_len = 6;	//negated length
		$p_len = 5; //plain length

		$complete_string_len = strlen($string);
		
		//This is very much a shortcut but one that does the trick remove 
		//occurrances of ':' to allow us to count the number of [AB:AC]
		//wildcards in a string to then calculate the minfraglength from...
		$strlen_card_removed = strlen(str_replace(':', '', $string));
		
		$no_bracketed_wildcards = $complete_string_len - $strlen_card_removed;
		
		$negation_no = 0;
		
		if ($no_bracketed_wildcards > 0)
		{
			$strlen_negation_removed = strlen(str_replace('!', '', $string));
			
			//number of wildcards with negation in...
			$negation_no = $complete_string_len - $strlen_negation_removed;
			
			//number of wildcards left without negation...
			$plain_count = $no_bracketed_wildcards - $negation_no;
			
			for ($i = 0; $i < $negation_no; $i++)
			{
				$complete_string_len = $complete_string_len - $n_len;
			}
			
			for ($i = 0; $i < $plain_count; $i++)
			{
				$complete_string_len = $complete_string_len - $p_len;
			}
		}
		
		$bytes = $complete_string_len / 2;
		
		//add value to fraglength count...
		$fraglen = $fraglen + $bytes;
		
		return $fraglen;
	}




	/*********************************************************************
	*
	* Function		:  
	*
	* Description	: 
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	//small function to ease the pain out outputting a new fragment value...
	private function fragXMLOutput($orientation, $string, $fragPos, $min, $max, $doc, $parent, $xml)
	{
		$frag = $xml->xmlCreateElement($orientation, $doc, $parent);
		$xml->xmlAddAttribute('MaxOffset', $max, $doc, $frag);
		$xml->xmlAddAttribute('MinOffset', $min, $doc, $frag);
		$xml->xmlAddAttribute('Position', $fragPos, $doc, $frag);
		$xml->xmlAddTextValue($string, $doc, $frag);	
	}




	/*********************************************************************
	*
	* Function		:  
	*
	* Description	: 
	*
	* parameters	: 	
	*          
	* Returns		:  
	*
	*********************************************************************/
	//split bracketed values into an array or single variable depending
	//on delimeter inbetween {'-'} ('|')
	private function getBracketedValues($delimeter, $fragments)
	{
		$len = strlen($fragments);
		$options =  '';
		
		//might not need array as we figure out how to output nodes
		//might do on the fly. array structures things nicely at the moment
		$optarr = array();
		
		//exit condition after open-bracket end before close-bracket
		for ($j = 1; $j < $len-1; $j++)
		{
			$optchar = $fragments[$j];
			
			if ($optchar == $delimeter)
			{
				$optarr[] = $options;
				$options = '';
			}
			else
			{
				$options = $options . $optchar;
			}
		}
		
		$optarr[] = $options;
		
		return $optarr;
	}
}
?>