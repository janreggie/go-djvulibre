package djvu

/** #DjVuDocument# provides convenient interface for opening, decoding
    and saving back DjVu documents in single page and multi page formats.

    {\bf Input formats}
    It can read multi page DjVu documents in either of the 4 formats: 2
    obsolete ({\em old bundled} and {\em old indexed}) and two new
    ({\em new bundled} and {\em new indirect}).

    {\bf Output formats}
    To encourage users to switch to the new formats, the #DjVuDocument# can
    save documents back only in the new formats: {\em bundled} and
    {\em indirect}.

    {\bf Conversion.} Since #DjVuDocument# can open DjVu documents in
    an obsolete format and save it in any of the two new formats
    ({\em new bundled} and {\em new indirect}), this class can be used for
    conversion from obsolete formats to the new ones. Although it can also
    do conversion between the new two formats, it's not the best way to
    do it. Please refer to \Ref{DjVmDoc} for details.

    {\bf Decoding.} #DjVuDocument# provides convenient interface for obtaining
    \Ref{DjVuImage} corresponding to any page of the document. It uses
    \Ref{DjVuFileCache} to do caching thus avoiding unnecessary multiple decoding of
    the same page. The real decoding though is accomplished by \Ref{DjVuFile}.

    {\bf Messenging.} Being derived from \Ref{DjVuPort}, #DjVuDocument#
    takes an active part in exchanging messages (requests and notifications)
    between different parties involved in decoding. It reports (relays)
    errors, progress information and even handles some requests for data (when
    these requests deal with local files).

    Typical usage of #DjVuDocument# class in a threadless command line
    program would be the following:
    \begin{verbatim}
    static const char file_name[]="/tmp/document.djvu";
    GP<DjVuDocument> doc=DjVuDocument::create_wait(file_name);
    const int pages=doc->get_pages_num();
    for(int page=0;page<pages;page++)
    {
       GP<DjVuImage> dimg=doc->get_page(page);
       // Do something
    };
    \end{verbatim}

    {\bf Comments for the code above}
    \begin{enumerate}
       \item Since the document is assumed to be stored on the hard drive,
             we don't have to cope with \Ref{DjVuPort}s and can pass
	     #ZERO# pointer to the \Ref{init}() function. #DjVuDocument#
	     can access local data itself. In the case of a plugin though,
	     one would have to implement his own \Ref{DjVuPort}, which
	     would handle requests for data arising when the document
	     is being decoded.
       \item In a threaded program instead of calling the \Ref{init}()
             function one can call \Ref{start_init}() and \Ref{stop_init}()
	     to initiate and interrupt initialization carried out in
	     another thread. This possibility of initializing the document
	     in another thread has been added specially for the plugin
	     because the initialization itself requires data, which is
	     not immediately available in the plugin. Thus, to prevent the
	     main thread from blocking, we perform initialization in a
	     separate thread. To check if the class is completely and
	     successfully initialized, use \Ref{is_init_ok}(). To see if
	     there was an error, use \Ref{is_init_failed}(). To
	     know when initialization is over (whether successfully or not),
	     use \Ref{is_init_complete}(). To wait for this to happen use
	     \Ref{wait_for_complete_init}(). Once again, all these things are
	     not required for single-threaded program.

	     Another difference between single-threaded and multi-threaded
	     environments is that in a single-threaded program, the image is
	     fully decoded before it's returned. In a multi-threaded
	     application decoding starts in a separate thread, and the pointer
	     to the \Ref{DjVuImage} being decoded is returned immediately.
	     This has been done to enable progressive redisplay
	     in the DjVu plugin. Use communication mechanism provided by
	     \Ref{DjVuPort} and \Ref{DjVuPortcaster} to learn about progress
	     of decoding.  Or try #dimg->wait_for_complete_decode()# to wait
	     until the decoding ends.
       \item See Also: \Ref{DjVuFile}, \Ref{DjVuImage}, \Ref{GOS}.
    \end{enumerate}

    {\bf Initialization}
    As mentioned above, the #DjVuDocument# can go through several stages
    of initialization. The functionality is gradually added while it passes
    one stage after another:
    \begin{enumerate}
       \item First of all, immediately after the object is created \Ref{init}()
             or \Ref{start_init}() functions must be called. {\bf Nothing}
	     will work until this is done. \Ref{init}() function will not
	     return until the initialization is complete. You need to make
	     sure, that enough data is available. {\bf Do not call \Ref{init}()
	     in the plugin}. \Ref{start_init}() will start initialization
	     in another thread. Use \Ref{stop_init}() to interrupt it.
	     Use \Ref{is_init_complete}() to check the initialization progress.
	     Use \Ref{wait_for_complete_init}() to wait for init to finish.
       \item The first thing the initializing code learns about the document
	     is its type (#BUNDLED#, #INDIRECT#, #OLD_BUNDLED# or #OLD_INDEXED#).
	     As soon as it happens, document flags are changed and
	     #notify_doc_flags_changed()# request is sent through the
	     communication mechanism provided by \Ref{DjVuPortcaster}.
       \item After the document type becomes known, the initializing code
             proceeds with learning the document structure. Gradually the
	     flags are updated with values:
	     \begin{itemize}
	        \item #DOC_DIR_KNOWN#: Contents of the document became known.
		      This is meaningful for #BUNDLED#, #OLD_BUNDLED# and
		      #INDIRECT# documents only.
		\item #DOC_NDIR_KNOWN#: Contents of the document navigation
		      directory became known. This is meaningful for old-style
		      documents (#OLD_BUNDLED# and #OLD_INDEXED#) only
		\item #DOC_INIT_OK# or #DOC_INIT_FAILED#:
		      The initializating code finished.
	     \end{itemize}
    \end{enumerate} */
type Document struct {
	Port
}
