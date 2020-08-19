# Signature Development Utility

[PRONOM][pronom-1]/[DROID][droid-1] Signature Development Utility source code,
first written in late 2011.

* **COPTR Entry:** [Signature development utility][coptr-1]

## Current version: 2.0

The new development is written in Golang and hosted on
[ffdev.info][ffdev-1]. It has been written to better support the container
signature workflow. It has also been written so that new features can be
developed easier as I tend not to write in PHP anymore. There is more
information at [ffdev.info][ffdev-1].

## Legacy version...

The first iteration of this application is hosted by
[The National Archives][tna-1] and mirrored on [my own site][expo-1]. It is
written in PHP and, well, it's a bit harder to maintain, but still it provides
a pretty pure implementation of what signature files used to be like in DROID 4
and 5 and largely 6, although the patterns are not de-constructed and compiled
differently in DROID 6 at runtime. 1.0 can be found in the releases section of
[this repository][gh-1].

## Contributing

Check out the [issues][issues-1] log for ideas for contributing and things I
hope to be working on.

[pronom-1]: http://www.nationalarchives.gov.uk/PRONOM/Default.aspx
[droid-1]: http://www.nationalarchives.gov.uk/information-management/manage-information/preserving-digital-records/droid/
[tna-1]: http://www.nationalarchives.gov.uk/pronom/sigdev/index.htm
[expo-1]: http://exponentialdecay.co.uk/sd/index.htm
[issues-1]: https://github.com/exponential-decay/signature-development-utility/issues
[coptr-1]: http://coptr.digipres.org/PRONOM_Signature_Development_Utility
[ffdev-1]: http://ffdev.info
[gh-1]: https://github.com/exponential-decay/signature-development-utility/releases/tag/1.0
