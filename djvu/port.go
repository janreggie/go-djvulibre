package djvu

// Port is an interface for sending and receiving messages generated during decoding process.
// In order for clases to send or receive requests,
// they must embed a *SimplePort, which implements Port.
type Port interface {
	// Copy constructor.
	// When Ports are copied,
	// the portcaster copies all incoming and outgoing routes of the original.
	Copy() Port

	// Should return true if the called class "inherits" class `className`.
	// When a destination receives a request,
	// it can retrieve the pointer to the source Port.
	// This virtual function should be able to help to identify the source of the request.
	// For example, `File` is also derived from `Port`.
	// In order for the receiver to recognize the sender,
	// `File` should override this function to return true
	// when `className` is either `Port` or `File`.
	Inherits(className string) bool

	// This request is issued to request translation of the ID,
	// used in an DjVu INCL chunk to a URL,
	// which may be used to request data associated with included file.
	// A Document usually intercepts all such requests,
	// and the user doesn't have to worry about the translation.
	IdToUrl(source Port, id string) Url

	// This request is used to get a file corresponding to the given ID.
	// A Document is supposed to intercept it
	// and either create a new instance of File or reuse an existing one from the cache.
	IdToFile(source Port, id string) *File

	// This request is issued when decoder needs additional data for decoding.
	// Both File and Document are initialized with a URL, not the document data.
	// As soon as they need the data, they call this function,
	// whose responsibility is to locate the source of the data basing on the Url passed
	// and return it back in the form of the DataPool.
	// If this particular receiver is unable to fullfil the request, it should return false.
	RequestData(source Port, url *Url) *DataPool

	// This notification is sent when an error occurs
	// and the error message should be shown to the user.
	// Returns whether the receiver is able to process the request.
	NotifyError(source Port, msg string) bool

	// This notification is sent to update the decoding status.
	// Returns whether the receiver is able to process the request.
	NotifyStatus(source Port, msg string) bool

	// This notification is sent by an Image when it should be redrawn.
	// It may be used to implement progressive redisplay.
	NotifyRedisplay(source *Image)

	// This notification is sent by an Image
	// when its geometry has been changed as a result of decoding.
	// It may be used to implement progressive redisplay.
	NotifyRelayout(source *Image)

	// This notification is sent when a new chunk has been decoded.
	NotifyChunkDone(source Port, name string)

	// This notification is sent after the File flags have been changed.
	// This happens, for example, when:
	//
	// - Decoding succeeded, failed or just stopped
	// - All data has been received
	// - All included files have been created
	NotifyFileFlagsChanged(source *File, setMask int64, clearMask int64)

	// This notification is sent after the File flags have been changed.
	// This happens, for example, after it receives enough data
	// and can determine its structure (BUNDLED, OLD_INDEXED, etc.)
	NotifyDocFlagsChanged(source *Document, setMask int64, clearMask int64)

	// This notification is sent from time to time while decoding is in progress.
	// The purpose is obvious:
	// to provide a way to know how much is done
	// and how long the decoding will continue.
	// Argument `done` is a number from 0 to 1 reflecting the progress.
	NotifyDecodeProgress(source Port, done float64)
}

// This is the standard types for defining what to do in case of errors.
// This is only used by some of the subclasses,
// but it needs to be defined here to guarantee all subclasses use the same enum types.
// In general, many errors are non recoverable.
// Using a setting other than ABORT may just result in even more errors.
type ErrorRecoveryAction int64

const (
	ABORT ErrorRecoveryAction = iota
	SKIP_PAGES
	SKIP_CHUNKS
	KEEP_ALL
)
