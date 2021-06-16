A library for reading, ~~writing, and modifying~~ epub3 files. The other parts will come later.

While working on Quill, we discovered that there are no good libraries for reading and writing epub files in go. So this library aspires to let you read and write epubs using go. The first thing we need is reading metadata, so that's where we are right now. Expect more functionality as Quill's needs expand.

A lot of the xml-related code in this package is working around the broken stdlib xml implementation. For example, the golang/xml library does not handle writing namespaces correctly, which is critical for epubs (and most any other practical xml). When we can get around to writing our own xml parser/writer it should get a lot simpler.

For more information on epubs, see
* [W3EPUB3 Overview](https://www.w3.org/publishing/epub3/epub-overview.html)
* [W3 EPUB3 Spec](https://www.w3.org/publishing/epub3/epub-spec.html)