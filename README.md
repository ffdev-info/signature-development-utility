# Signature Development Utility

[PRONOM][pronom-1]/[DROID][droid-1] Signature Development Utility source code,
first written in late 2011.

* **COPTR Entry:** [Signature development utility][coptr-1]

## Introduction

This code represents the pairing of standard signature development and
container signature development.

The new code is written in Golang. I wrote this replacement for the original
utility in 2020, but some of my assumptions about the structure of the new
signature files were incorrect, so for a while, I bootstrapped the code
onto my old PHP release to provide you with standard and container
signature capabilities.

I have since researched and corrected my assumptions, so I can now offer
you an entirely self-contained executable that allows you to write
PRONOM-compatible file format signatures in the comfort of your own server.

Now self-contained, the new version should be much easier to develop and
add features to. See below under [contributing](#contributing).

### Hosting

The utility is also hosted on [ffdev.info][ffdev-1].

### Installation

Download a suitable [release][release-1] and simply run the executable from
your command line. A port will be opened and you will be able to navigate
to it via localhost in your own browser. Follow the web-page prompts to
begin writing your signatures.

[release-1]: https://github.com/ffdev-info/signature-development-utility/releases

## Signature development

For more background on file format signature development,
David C., Francesca, Tyler, and I have created a resource:
[Searching for a signature: A file format identification workshop][search-1].

## Legacy version

My first iteration of this application was hosted by
[The National Archives][tna-1] and remains mirrored, for the time
being, on [my homepage][expo-1]. It was written in PHP and is no
longer maintained.

In the time it existed it enabled a greater
level of collaboration between external signature developers and
The National Archives, UK. Almost
immediately, Bill Underwood at Georgia Tech Research Institute began a
project submitting file format signatures developed using the utility,
a task that would have been much harder without democratizing tooling
such as this. All these years later, the previous utility also enabled
development of signature capabilities in new  PRONOM, i.e. PRONOM 2026.

The legacy implementation  provided a pure implementation of DROID 4.0
copmatible signatures, the only reference for developers outside of
the original PRONOM stored procedures used to achieve the same result.

A release of this utility is available in [this repository][gh-1].

## Contributing

Check out the [issues][issues-1] log for existing issues.

<!-- reference links -->

[pronom-1]: https://pronom.nationalarchives.gov.uk/
[droid-1]: https://www.nationalarchives.gov.uk/information-management/manage-information/preserving-digital-records/droid/
[tna-1]: https://web.archive.org/web/20210530124335/http://www.nationalarchives.gov.uk/pronom/sigdev/index.htm
[expo-1]: http://exponentialdecay.co.uk/sd/index.htm
[issues-1]: https://github.com/ffdev-info/signature-development-utility/issues
[coptr-1]: http://coptr.digipres.org/PRONOM_Signature_Development_Utility
[ffdev-1]: http://ffdev.info
[gh-1]: https://github.com/ffdev-info/signature-development-utility/releases/tag/1.0
[search-1]: https://ffdev-info.github.io/searching-for-a-signature/
