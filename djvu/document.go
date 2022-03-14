package djvu

// Document allows for opening, decoding and saving back DjVu documents
// in single page and multi page formats.
//
// Input formats
//
// It can read multi page DjVu documents in either of the 4 formats:
// 2 obsolete ("old bundled" and "old indexed")
// and two new ("new bundled" and "new indirect").
//
// Output formats
//
// To encourage users to switch to the new formats,\
// Document can only save documents back only in the new formats:
// "bundled" and "indirect".
//
// Conversion
//
// Since Document can open DjVu documents in an obsolete format
// and save it in any of the two new formats,
// this class can be used for conversion
// from obsolete formats to the new ones.
// Although it can also do conversion between the new two formats,
// it's not the best way to do it.
// Please refer to Multidoc for details.
//
// Decoding
//
// Document provides convenient interface for obtaining Images
// corresponding to any page of the document.
// It uses FileCache to do caching
// thus avoiding unnecessary multiple decoding of the same page.
// The real decoding though is accomplished by File.
//
// Messaging
//
// Containing an instance of Port,
// Document takes an active part in exchanging messages
// (requests and notifications)
// between different parties involved in decoding.
// It reports (relays) errors,
// progress information
// and even handles some requests for data
// (when these requests deal with local files).
//
// Typical use of Document would be the following:
//
//     // TODO: Change this code when done
//     filename := "/tmp/document.djvu"
//     doc := CreateWaitDocument(filename)
//     pages := doc.GetPagesNum()
//     for p := 0; p < pages; p++ {
//         page := doc.GetPage(p)
//         // Do something with img
//     }
//
// TODO: Change the following comments
// - Since the document is assumed to be stored on the hard drive,
// we don't have to cope with \Ref{DjVuPort}s and can pass
// #ZERO# pointer to the \Ref{init}() function. #DjVuDocument#
// can access local data itself. In the case of a plugin though,
// one would have to implement his own \Ref{DjVuPort}, which
// would handle requests for data arising when the document
// is being decoded.
// \item In a threaded program instead of calling the \Ref{init}()
// function one can call \Ref{start_init}() and \Ref{stop_init}()
// to initiate and interrupt initialization carried out in
// another thread. This possibility of initializing the document
// in another thread has been added specially for the plugin
// because the initialization itself requires data, which is
// not immediately available in the plugin. Thus, to prevent the
// main thread from blocking, we perform initialization in a
// separate thread. To check if the class is completely and
// successfully initialized, use \Ref{is_init_ok}(). To see if
// there was an error, use \Ref{is_init_failed}(). To
// know when initialization is over (whether successfully or not),
// use \Ref{is_init_complete}(). To wait for this to happen use
// \Ref{wait_for_complete_init}(). Once again, all these things are
// not required for single-threaded program.
//
// Another difference between single-threaded and multi-threaded
// environments is that in a single-threaded program, the image is
// fully decoded before it's returned. In a multi-threaded
// application decoding starts in a separate thread, and the pointer
// to the \Ref{DjVuImage} being decoded is returned immediately.
// This has been done to enable progressive redisplay
// in the DjVu plugin. Use communication mechanism provided by
// \Ref{DjVuPort} and \Ref{DjVuPortcaster} to learn about progress
// of decoding.  Or try #dimg->wait_for_complete_decode()# to wait
// until the decoding ends.
// \item See Also: \Ref{DjVuFile}, \Ref{DjVuImage}, \Ref{GOS}.
// \end{enumerate}
//
// Initialization
//
// As mentioned above,
// the Document can go through several stages of initialization.
// The functionality is gradually added
// while it passes one stage after another:
//
// TODO: Update docs...
type Document struct {
	Port
}
