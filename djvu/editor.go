package djvu

/** #DjVuDocEditor# is an extension of \Ref{DjVuDocument} class with
  additional capabilities for editing the document contents.

  It can be used to:
  \begin{enumerate}
     \item Create (compose) new multipage DjVu documents using single
           page DjVu documents. The class does {\bf not} do compression.
     \item Insert and remove different pages of multipage DjVu documents.
     \item Change attributes ({\em names}, {\em IDs} and {\em titles})
           of files composing the DjVu document.
     \item Generate thumbnail images and integrate them into the document.
  \end{enumerate}
*/
type Editor struct {
	Document
}
