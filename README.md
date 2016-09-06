signature-development-utility
=============================

[PRONOM](http://www.nationalarchives.gov.uk/PRONOM/Default.aspx)/[DROID](http://www.nationalarchives.gov.uk/information-management/manage-information/preserving-digital-records/droid/) Signature Development Utility source code, first written in late 2011. 

This application is hosted by The National Archives: http://www.nationalarchives.gov.uk/pronom/sigdev/index.htm
and mirrored on my own site: http://exponentialdecay.co.uk/sd/index.htm

This code is hosted here to enable it to be built upon and improved by myself and others. The first point of call
will be to build unit tests to help unravel some of the complexity of the code and make it more modular. There are
also a handful of known issues to be dealt with, see the issues log for more information.

###Contribution

The form represents the limits of my JQuery knowledge at the time. We have three fields that we can create dynamically but it
would be nice if we could create forms dynamically also.

* Update index.htm to enable multiple forms to be created for a single submit
* Test process_signature_form.php to ensure that forms (number should be in 'Counter' variable are processed

All important data is submitted through process_signature_form.php where it is processed. Signature information for example
is processed through [generateSignatureCollection($count, $_POST)](https://github.com/exponential-decay/signature-development-utility/blob/master/php/process_signature_form.php#L85) 
and then once returned, InternalSignatureCollections and FileFormatCollections are pieced together as XML before download by the user. 

Signature information is output by [generatebytecode.php](https://github.com/exponential-decay/signature-development-utility/blob/master/php/generatebytecode/generatebytecode.php) where the argument is a SignatureCollection object (e.g. [generateSignatureFromObject($signature_collection) ](https://github.com/exponential-decay/signature-development-utility/blob/master/php/generatebytecode/generatebytecode.php#L63) the return is [PRONOM/DROID formatted XML](https://github.com/exponential-decay/signature-development-utility/blob/master/php/generatebytecode/generatebytecode.php#L148).

####Additional Contributions

* Unit tests would be wonderful. This is much harder for me now without access to PRONOM to write them against
the stored proecedure implmenetion. Any unit test would neeed to be written against this implementation and whatever can be
reverse engineered from the signature file. 

###License

Copyright (c) 2011 Ross Spencer

This software is provided 'as-is', without any express or implied warranty. In no event will the authors be held liable for any damages arising from the use of this software.

Permission is granted to anyone to use this software for any purpose, including commercial applications, and to alter it and redistribute it freely, subject to the following restrictions:

The origin of this software must not be misrepresented; you must not claim that you wrote the original software. If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
This notice may not be removed or altered from any source distribution.

