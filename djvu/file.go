package djvu

import "sync"

// File plays the central role in decoding Images.
// First of all, it represents a DjVu file
// whether it's part of a multipage all-in-one-file DjVu document,
// or part of a multipage DjVu document where every page is in a separate file,
// or the whole single page document.
// File can read its contents from a file and store it back when necessary.
//
// Second, File does the greatest part of decoding work.
// In the past this was the responsibility of Image.
// Now, with the introduction of the multipage DjVu formats,
// the decoding routines have been extracted from Image
// and put into this separate class File.
//
// As Image before,
// File now contains public class variables corresponding to every component,
// that can ever be decoded from a DjVu file
// (such as INFO chunk, BG44 chunk, SJBZ chunk, etc.).
//
// Decoding can be started with the StartDecode method.
//
// Inclusion is also a new feature specifically designed for a multipage document.
// Indeed, inside a given document there can be a lot of things shared between its pages.
// Examples can be the document annotation Anno
// and other things like shared shapes and dictionary (to be implemented).
// To avoid putting these chunks into every page,
// we have invented new chunk called INCL
// which purpose is to make the decoder open the specified file and decode it.
//
// Source of Data
//
// File can be initialized in two ways:
//
// - With a Url and Port.
// In this case File will request its data
// thru the communication mechanism provided by Port in the constructor.
// If this file references (includes) any other file,
// data for them will also be requested in the same way.
//
// - With a ByteStream.
// In this case File will read its data directly from the passed stream.
// This constructor has been added to simplify creation of Files,
// which do not include anything else.
// In this case the ByteStream is enough for the #DjVuFile# to initialize.
//
// Progress Information
//
// File does not do decoding silently.
// Instead, it sends a whole set of notifications
// through the mechanism provided by Port and PortCaster.
// It tells the user of this struct about the progress of the decoding,
// about possible errors, chunk being decoded, etc.
// The data is requested using this mechanism too.
//
// Creating
//
// Depending on where you have data of the DjVu file,
// File can be initialized in two ways:
//
// - By providing Url and pointer to Port.
// In this case File will request data using communication mechanism provided by Port.
// This is useful when the data is on the web or when this file includes other files.
//
// - By providing a ByteStream with the data for the file.
// Use it only when the file doesn't include other files.
//
// There are also several methods provided
// for composing the desired Document and modifying File structure.
// These include InsertFile and UnlinkFile.
//
// Caching
//
// In the case of plugin it's important to do the caching of decoded images or files.
// File appears to be the best candidate for caching,
// and that's why it supports this procedure.
// Whenever a File is successfully decoded,
// it's added to the cache by Document.
// Next time somebody needs it,
// it will be extracted from the cache directly by Document
// and won't be decoded again.
//
// URLs
//
// Historically the biggest strain is put on making the decoder available for Netscape and IE plugins
// where the original files reside somewhere in the net.
// That is why File uses Urls to identify itself and other files.
// If you're working with files on the hard disk,
// you have to use the local URLs instead of file names.
//
// Sometimes it happens that a given file does not reside anywhere but the memory.
// No problem in this case either.
// There is a special port MemoryPort which can associate any URL with the corresponding data in the memory.
// All you need to do is to invent your own URL prefix for this case.
// `memory:` will do.
// The usage of absolute URLs has many advantages
// among which is the capability to cache files with their URL being the cache key.
//
// Please note, that File has been designed to work closely with Document.
// So please review the documentation on the latter too.
type File struct {
	Port

	mtx         sync.RWMutex
	initialized bool
	flags       uint64

	decodeDataPool  *DataPool
	decodeLifeSaver *File
}
