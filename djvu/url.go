package djvu

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"
	"unicode"
)

// TODO: You might want to make Url an *interface*
// and `BaseUrl` a concrete struct
// for `Filename` and some other child interfaces to "work".
// But for now, let's stick to this.

// Url is a System independent URL representation.
//
// This class is used in the library to store URLs in a system independent format.
// The idea to use a general class to hold URLs arose after we realized
// that DjVu had to be able to access files both from the WEB and from the local disk.
// While it is strange to talk about system independence of HTTP URLs,
// file names formats obviously differ from platform to platform.
// They may contain forward slashes, backward slashes, colons as separators, etc.
// There maybe more than one URL corresponding to the same file name.
// Compare `file:/dir/file.djvu` and `file://localhost/dir/file.djvu`.
//
// To simplify a developer's life we have created this class,
// which contains inside a canonical representation of URLs.
//
// File URLs are converted to internal format with the help of the os stdlib.
//
// All other URLs are modified to contain only forward slashes.
type Url struct {
	mtx sync.RWMutex
	url string

	// Hint: for cgi, use net/url.Values...
	// Except you'd be using CgiName(int) so index does matter...
	// More info on CGI: <https://en.wikipedia.org/wiki/Common_Gateway_Interface>
	cgiNameArr  []string
	cgiValueArr []string

	validUrl bool
}

func NewUrl(urlString string) *Url {
	// TODO: Implement
	return nil
}

// Copy copies a Url
func (url *Url) Copy() *Url {
	panic("unimplemented")
}

// IsValid checks if the URL is valid.
// If invalid, reinitialize
func (url *Url) IsValid() bool {
	panic("unimplemented")
}

// Extracts the Protocol part from the URL and returns it
func (url *Url) Protocol() string {
	panic("unimplemented")
}

// Returns string after the first `#`
// with decoded escape sequences.
func (url *Url) HashArgument() string {
	panic("unimplemented")
}

// Inserts `arg` after a separating hash into the URL.
func (url *Url) SetHashArgument(arg string) {
	panic("unimplemented")
}

// Returns the total number of CGI arguments in the URL.
func (url *Url) CgiArguments() int {
	panic("unimplemented")
}

// Returs the total number of DjVu-related CGI arguments
// (arguments following `DJVUOPTS` in the URL)
func (url *Url) DjvuCgiArguments() int {
	panic("unimplemented")
}

// Returns that part of CGI argument `num`,
// which is before the equal sign
func (url *Url) CgiName(num int) string {
	panic("unimplemented")
}

// Returns that part of DjVu-related CGI argument number `num`,
// which is before the equal sign
func (url *Url) DjvuCgiName(num int) string {
	panic("unimplemented")
}

// Returns that part of CGI argument number `num`,
// which is after the equal sign
func (url *Url) CgiValue(num int) string {
	panic("unimplemented")
}

// Returns that part of DjVu-related CGI argument number `num`,
// which is after the equal sign
func (url *Url) DjvuCgiValue(num int) string {
	panic("unimplemented")
}

// Returns array of all known CGI names
// (part of CGI argument before the equal sign)
func (url *Url) CgiNames() []string {
	url.mtx.RLock()
	defer url.mtx.RUnlock()

	retval := make([]string, len(url.cgiNameArr))
	copy(retval, url.cgiNameArr)
	return retval
}

// Returns array of all known DjVu-related CGI arguments
// (arguments following `DJVUOPTS` option)
func (url *Url) DjvuCgiNames() []string {
	panic("unimplemented")
}

// Returns array of all known CGI names
// (part of CGI argument before the equal sign)
// TODO: Documentation above may be misleading!!! Did you mean *after* equal sign?
func (url *Url) CgiValues() []string {
	url.mtx.RLock()
	defer url.mtx.RUnlock()

	retval := make([]string, len(url.cgiValueArr))
	copy(retval, url.cgiValueArr)
	return retval
}

// Returns array of values of DjVu-related CGI arguments
// (arguments following `DJVUOPTS` option)
func (url *Url) DjvuCgiValues() []string {
	panic("unimplemented")
}

// Erases everything after the first `#` or `?`
func (url *Url) ClearAllArguments() {
	panic("unimplemented")
}

// Erases everything after the first `#`
func (url *Url) ClearHashArguments() {
	panic("unimplemented")
}

// Erases DjVu CGI arguments (following `DJVUOPTS`)
func (url *Url) ClearDjvuCgiArguments() {
	panic("unimplemented")
}

// Erases all CGI arguments (following the first `?`)
func (url *Url) ClearCgiArguments() {
	panic("unimplemented")
}

// Appends the specified CGI argument.
// Will insert `DJVUOPTS` if necessary.
// TODO: Pointer necessary?
func (url *Url) AddDjvuCgiArgument(name string, value *string) {
	panic("unimplemented")
}

// Returns the URL corresponding to the dictionary
// containing the document with this URL.
// The function basically takes the URL
// and clears everything after the last slash.
func (url *Url) Base() *Url

// Returns the absolute URL without the host part.
func (url *Url) Pathname() string

// Returns the name part of this URL.
// For example, if the URL is `http://www.lizardtech.com/file%201.djvu`,
// then this function will return `file%201.djvu`.
// Contrast with Fname which returns `file 1.djvu`.
func (url *Url) Name() string

// Returns the name part of this URL with escape sequences expanded.
// For example, if the URL is `http://www.lizardtech.com/file%201.djvu`,
// then this function will return `file 1.djvu`.
// Contrast with Name which returns `file%201.djvu`.
func (url *Url) Fname() string {
	panic("unimplemented")
}

// Returns the extension part of name of document in this URL.
func (url *Url) Extension() string {
	panic("unimplemented")
}

// Checks if this is an empty URL
func (url *Url) IsEmpty() bool {
	url.mtx.RLock()
	defer url.mtx.RUnlock()
	return len(url.url) == 0
}

// Checks whether the URL is local (i.e., starts with `file:/`)
func (url *Url) IsLocalFileUrl() bool

// Checks whether two URLs are the same
func (url *Url) Equal(rhs *Url) bool

// Returns internal URL representation.
func (url *Url) Raw() string {
	url.mtx.RLock()
	defer url.mtx.RUnlock()

	return url.url
}

// Applies heuristic rules to convert a URl into a valid file name.
// Returns a simple basename in case of failure.
//
// TODO: Export logic to some helper function, then export this method, effectively calling that helper.
func (url *Url) utf8Filename() string {
	if url.url == "" {
		return ""
	}

	retval := ""
	uu := decodeReserved(url.url)

	// Expect file URL to start with `file:` (filespec)
	if !strings.HasPrefix(uu, filespec) {
		return path.Base(uu)
	}
	uu = uu[len(filespec):]

	if runtime.GOOS == osMac {
		// Remove leading slashes
		uu = strings.TrimLeft(uu, "/")
		uu = strings.TrimPrefix(uu, localhost)
		uu = strings.TrimLeft(uu, "/")
	} else {
		if strings.HasPrefix(uu, localhostspec1) {
			uu = strings.TrimPrefix(uu, localhostspec1) // RFC 1738 local host form
		} else if strings.HasPrefix(uu, localhostspec2) {
			uu = strings.TrimPrefix(uu, localhostspec2) // RFC 1738 local host form
		} else if len(uu) > 4 && //  "file://<letter>:/<path>"
			uu[:2] == "//" && // "file://<letter>|/<path>"
			unicode.IsLetter(rune(uu[2])) &&
			(uu[3] == colon || uu[3] == vertical) && uu[4] == slash {
			uu = uu[2:]
		} else if len(uu) > 2 && // "file:/<path>"
			uu[0] == slash && uu[1] != slash {
			uu = uu[1:]
		}
	}

	// Check if we are finished
	if runtime.GOOS == osMac {
		// TODO: Implement...
		panic("unimplemented")
	} else {
		retval = expandName(uu, root)
	}

	if runtime.GOOS == osWindows || runtime.GOOS == osMac {
		if unicode.IsLetter(rune(uu[0])) && uu[1] == vertical && uu[2] == slash {
			drive := fmt.Sprintf("%v%v%v", uu[0], colon, backslash)
			retval = expandName(uu[3:], drive)
		}
	}

	return retval
}

// Returns a string representation of the URL.
// This function normally returns a standard file URL as described in RFC 1738.
// Some versions of MSIE do not support this standard syntax.
// A brain damaged MSIE compatible syntax is generated
// when the optional argument `useragent` contains string `MSIE` or `Microsoft`.
func (url *Url) GetStringWithUseragent(useragent string) string

// TODO: Could we integrate this with GetString above?
// TODO: Remove the url.validUrl and add validation into url.init
func (url *Url) GetStringWithErr(nothrow bool) string {
	url.mtx.Lock()
	defer url.mtx.Unlock()
	if !url.validUrl {
		url.init()
	}
	return url.url
}

// Return whether this URL is an existing file, directory, or device.
func (url *Url) IsLocalPath() bool {
	panic("unimplemented")
}

// Return whether this URL is an existing file
func (url *Url) IsFile() bool {
	panic("unimplemented")
}

// Return whether this URL is an existing directory
func (url *Url) IsDir() bool {
	panic("unimplemented")
}

// Follows symbolic links
func (url *Url) FollowSymlinks() *Url {
	panic("unimplemented")
}

// Creates the specified directory
//
// TODO: Returns what?
func (url *Url) Mkdir() int {
	panic("unimplemented")
}

// Deletes file or directory.
// Directories are not deleted unless the directory is empty.
// Returns a negative number if an error occurs.
//
// TODO: Return an error instead.
//
// TODO: Create a "service" which does operating system manipulation
// instead of turning them as methods for URL
func (url *Url) DeleteFile() int {
	panic("unimplemented")
}

// Recursively erases contents of directory.
// The directory itself will not be removed.
func (url *Url) ClearDir(timeout int) int {
	panic("unimplemented")
}

// Rename a file or directory
func (url *Url) RenameTo(newUrl *Url) int {
	panic("unimplemented")
}

// List the contents of a directory
func (url *Url) ListDir() []*Url {
	panic("unimplemented")
}

// Returns a filename for a URL.
// Argument must be a legal file URL.
// This function applies heuristic rules to convert the URL into a valid file name.
// It is guaranteed that this function can properly parse all URLs
// generated by `filename_to_url`
// The heuristics also work better when the file actually exists.
// An error is returned when this function cannot parse the URL
// or when the URL is not a file URL.
//
// URL formats are as described in RFC 1738
// plus the following alternative formats for files on the local host:
//
//     file://<letter>:/<path>
//     file://<letter>|/<path>
//     file:/<path>
//
// which are accepted because various browsers recognize them.
//
// TODO: Can we use os for this?
func (url *Url) Filename() (string, error) {
	panic("unimplemented")
}

// Hashing function
func (url *Url) Hash() uint32 {
	panic("unimplemented")
}

func (url *Url) init() error {
	url.validUrl = true
	if url.url == "" {
		return nil
	}

	protocol := protocol(url.url)
	if len(protocol) < 2 {
		return errors.New("GURL.no_protocol " + url.url)
	}

	// For the `localhost` protocol
	if strings.HasPrefix(url.url, localhost) {
		// Take the arguments first
		argsStr := strings.TrimLeftFunc(url.url, func(r rune) bool { return !isArgumentInit(r) })
		url.url = url.url[0 : len(url.url)-len(argsStr)] // Before the start of the arguments

		// Do things here...

		// Append the arguments back
		url.url += argsStr
	}

	panic("unimplemented") // TODO
}

func (url *Url) convertSlashes() {
	if runtime.GOOS == "windows" {
		xurl := url.GetStringWithErr(false)
		protocol := protocol(xurl)
		remaining := xurl[len(protocol):]
		remaining = strings.ReplaceAll(remaining, "/", "\\")
		url.url = protocol + remaining
	}
}

func (url *Url) beautifyPath() {
	panic("unimplemented")
}

func (url *Url) parseCgiArgs() {
	panic("unimplemented")
}

func (url *Url) storeCgiArgs() {
	panic("unimplemented")
}

// Escape special characters
func encodeReserved(gs string) string {
	panic("unimplemented")
}

// Decodes reserved characters from the URL
func decodeReserved(urlString string) string {
	panic("unimplemented")
}

func beautifyPath(urlString string) string {
	panic("unimplemented")
}

// Returns the full path name of filename interpreted relative to fromDir.
// Use current working dir when fromDir is empty.
func expandName(filename string, fromDir string) string {
	panic("unimplemented")
}

// protocol extracts the protocol part of the url string.
// Example: `protocol("https://www.google.com")` returns the string `https`.
// If a protocol cannot be found, return an empty string.
func protocol(urlString string) string {
	remaining := strings.TrimLeft(urlString, alphanum+"+-.")
	if len(remaining) >= 3 && remaining[:3] == "://" {
		protocolLength := len(urlString) - len(remaining)
		return urlString[:protocolLength]
	}
	return ""
}

// Returns whether r forms the start of the arguments whether in hash or CGI
func isArgumentInit(r rune) bool {
	return r == '#' || r == '?'
}

// Returns whether r is an argument separator whether in hash or CGI
func isArgumentSeparator(r rune) bool {
	return r == '&' || r == ';'
}

// Returns the hex value of a character, -1 if it isn't a hex
func hexVal(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	if r >= 'A' && r <= 'F' {
		return int(r - 'A' + 10)
	}
	if r >= 'a' && r <= 'f' {
		return int(r - 'a' + 10)
	}

	return -1
}

// Filename represents a File URL.
// Idk why that should be independent but okay.
// For now we're just directly copying from the original source code
// and changing things along the way...
type Filename struct {
	Url
}
