package djvu

// MultiDoc allows reading of multipage documents
//
// The "new" DjVu multipage documents can be of two types:
// "bundled" and "indirect".
// In the first case all pages are packed into one file,
// which is very like an archive internally.
// In the second case every page is stored in a separate file.
// Plus there can be other components,
// included into one or more pages,
// which also go into separate files.
// In addition to pages and components,
// in the case of the "indirect" format,
// there is one more top-level file with the document directory
// (see MultiDir), which is basically an index file containing the
// list of all files composing the document.
//
// This class can read documents of both formats
// and can save them under any format.
// It is therefore ideal for converting
// between "bundled" and "indirect" formats.
// It cannot be used however for reading obsolete formats.
// The best way to convert obsolete formats
// consists in reading them with Document
// and saving them using Write or Expand methods.
//
// This class can also be used
// to create and modify multipage documents at the low level
// without decoding every page or component
// (See InsertFile and DeleteFile methods).
//
type MultiDoc struct {
}
