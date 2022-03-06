package djvu

import "runtime"

const (
	djvuopts        = "DJVUOPTS"
	localhost       = "file://localhost/"
	localhostspec1  = "//localhost/"
	localhostspec2  = "///"
	filespecslashes = "file://"
	filespec        = "file:"

	backslash = '\\'
	colon     = ':'
	dot       = '.'
	percent   = '%'
	slash     = '/'
	tilde     = '~'

	maxpathlen = 1024
)

const (
	alphanum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var root = func() string {
	switch runtime.GOOS {
	case "darwin":
		return ""
	case "linux":
		return "/"
	case "windows":
		return "\\"
	}
	panic("define something here for your operating system")
}()
