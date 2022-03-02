package djvu

/** System independent URL representation.

  This class is used in the library to store URLs in a system independent
  format. The idea to use a general class to hold URL arose after we
  realized, that DjVu had to be able to access files both from the WEB
  and from the local disk. While it is strange to talk about system
  independence of HTTP URLs, file names formats obviously differ from
  platform to platform. They may contain forward slashes, backward slashes,
  colons as separators, etc. There maybe more than one URL corresponding
  to the same file name. Compare #file:/dir/file.djvu# and
  #file://localhost/dir/file.djvu#.

  To simplify a developer's life we have created this class, which contains
  inside a canonical representation of URLs.

  File URLs are converted to internal format with the help of \Ref{GOS} class.

  All other URLs are modified to contain only forward slashes.
*/
type URL struct{}
