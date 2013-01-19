$(document).ready(function() 
{
	$('#btnAdd').click(function()
	{
		var num		= $('.clonedInput').length;											// how many "duplicatable" input fields we currently have
		var newNum	= new Number(num + 1);													// the numeric ID of the new input field being added
		var newElem = $('#input' + num).clone().attr('id', 'input' + newNum); 	// create the new element via clone(), and manipulate it's ID using newNum value

		// manipulate the name/id values of the input inside the new element
		newElem.children(':eq(1)').attr('id', 'signature' + newNum)
										  .attr('name', 'signature' + newNum)
										  .attr('value', '2525454F46(0A|0D|0D0A)');

		// manipulate the name/id values of the input inside the new element
		newElem.children(':eq(4)').attr('id', 'anchor' + newNum)
										  .attr('name', 'anchor' + newNum)
										  .attr('value', 'EOFoffset');

		// manipulate the name/id values of the input inside the new element
		newElem.children(':eq(6)').attr('id', 'offset' + newNum)
										  .attr('name', 'offset' + newNum)
										  .attr('value', '0');
										  
		// manipulate the name/id values of the input inside the new element
		newElem.children(':eq(8)').attr('id', 'maxoffset' + newNum)
										  .attr('name', 'maxoffset' + newNum)
										  .attr('value', '0');

		// insert the new element after the last "duplicatable" input field
		$('#input' + num).after(newElem);

		// enable the "remove" button
		$('#btnDel').attr('disabled','');

		// business rule: you can only add 5 names
		if (newNum == 3)
		{
			$('#btnAdd').attr('disabled','disabled');
		}
		
		$('.counter').attr('value', num+1);
	});

	$('#btnDel').click(function() 
	{
		var num	= $('.clonedInput').length;	// how many "duplicatable" input fields we currently have
		$('#input' + num).remove();				// remove the last element

		// enable the "add" button
		$('#btnAdd').attr('disabled','');

		// if only one element remains, disable the "remove" button
		if (num-1 == 1)
		{
			$('#btnDel').attr('disabled','disabled');
		}
		
		$('.counter').attr('value', num-1);
	});

	$('#btnDel').attr('disabled','disabled');
	
	$(".help-text").hide();
	  //toggle the componenet with class msg_body
	  jQuery(".help").click(function()
	  {
		 jQuery(this).next(".help-text").slideToggle("Slow");
	  });
});