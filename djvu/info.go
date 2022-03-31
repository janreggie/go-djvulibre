package djvu

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/janreggie/go-djvulibre/djvu/bytestream"
	"github.com/pkg/errors"
)

// Info represents information about a DjVu File.
// This chunk should be the first chunk of a DjVu file.
type Info struct {
	// Width of the image (in pixels)
	Width uint16

	// Height of the image (in pixels)
	Height uint16

	// DjVu file version number.
	FileVersion FileVersion

	// Resolution of the DjVu image.
	// The resolution is given in "pixels per 2.54 centimeters"
	// (this unit is sometimes called "pixels per inch").
	// Display programs can use this information
	// to determine the natural magnification to use
	// for rendering a DjVu image.
	Dpi uint16

	// Gamma coefficient of the display for which the image was designed.
	// The rendering functions can use this information
	// in order to perform color correction for the intended display device.
	Gamma float64

	// How the File is rotated
	Orientation Orientation
}

func NewInfo() *Info { return &Info{} }

// Decode decodes the a DjVu's `INFO` chunk.
// This function reads binary data from a Reader stream
// and returns the DjVu file, and an error if it exists.
func (i *Info) Decode(r io.Reader) error {
	// Set to default values
	i.Width = 0
	i.Height = 0
	i.FileVersion = DJVUVERSION
	i.Dpi = 300
	i.Gamma = 2.2
	i.Orientation = ORIENTATION_0

	buffer := make([]byte, 10)
	size, err := io.ReadAtLeast(r, buffer, 10)
	if err != nil {
		return errors.Wrapf(err, "could not read INFO chunk")
	}
	if size == 0 {
		return io.ErrUnexpectedEOF
	}
	if size < 5 {
		return errors.New("corrupt file from INFO chunk")
	}
	i.Width = bytesToUint16(buffer[0], buffer[1])
	i.Height = bytesToUint16(buffer[2], buffer[3])

	i.FileVersion = FileVersion(buffer[4])
	if size > 5 && buffer[5] != 0xff {
		i.FileVersion = FileVersion(bytesToUint16(buffer[5], buffer[4]))
	}

	// Make sure that DPI is within the correct values
	if size > 7 && buffer[7] != 0xff {
		i.Dpi = bytesToUint16(buffer[7], buffer[6])
	}
	if i.Dpi < 25 || i.Dpi > 6000 {
		i.Dpi = 300
	}

	// Make sure that Gamma is within 0.3 to 5.0
	if size > 8 {
		i.Gamma = 0.1 * float64(buffer[8])
	}
	if i.Gamma < 0.3 {
		i.Gamma = 0.3
	}
	if i.Gamma > 5.0 {
		i.Gamma = 5.0
	}

	var flags uint8 // first 5 bits reserved for future implementation
	if size > 9 {
		flags = buffer[9]
	}
	switch Orientation(flags & 0b111) {
	case ORIENTATION_90:
		i.Orientation = ORIENTATION_90
	case ORIENTATION_180:
		i.Orientation = ORIENTATION_180
	case ORIENTATION_270:
		i.Orientation = ORIENTATION_270
	default:
		i.Orientation = ORIENTATION_0
	}

	return nil
}

// Encode encodes the `INFO` chunk.
// This function writes the fields of this Info object
// into a Wrtiter.
func (i *Info) Encode(w bytestream.Writer) error {
	var flags uint8
	switch i.Orientation {
	case ORIENTATION_90:
		flags = uint8(ORIENTATION_90)
	case ORIENTATION_180:
		flags = uint8(ORIENTATION_180)
	case ORIENTATION_270:
		flags = uint8(ORIENTATION_270)
	default:
		flags = uint8(ORIENTATION_0)
	}

	return w.Write16(i.Width).
		Write16(i.Height).
		Write8(uint8(i.FileVersion & 0xff)).
		Write8(uint8(i.FileVersion >> 8)).
		Write8(uint8(i.Dpi & 0xff)).
		Write8(uint8(i.Dpi >> 8)).
		Write8(uint8(10*i.Gamma + 0.5)).
		Write8(flags).
		Error()
}

// Returns the number of bytes used by this object
func (i *Info) GetMemoryUsage() uintptr {
	return unsafe.Sizeof(Info{})
}

// Obtain the flags for the default specifications.
func (i *Info) GetParamtags() string {
	var sb strings.Builder
	if i.Orientation != ORIENTATION_0 {
		fmt.Fprintf(&sb, `<PARAM name="ROTATE" value="%s" />`, i.Orientation)
		fmt.Println(&sb)
	}
	if i.Dpi != 0 {
		fmt.Fprintf(&sb, `<PARAM name="DPI" value="%v" />`, i.Dpi)
		fmt.Fprintln(&sb)
	}
	if i.Gamma != 0 {
		fmt.Fprintf(&sb, `<PARAM name="GAMMA" value="%v" />`, i.Dpi)
		fmt.Fprintln(&sb)
	}

	return sb.String()
}

// FileVersion determines the version of a specific File
// and whether the current version of DjVuLibre can support decoding it.
type FileVersion uint16

const (
	// Current DjVu format version.
	DJVUVERSION FileVersion = 26

	// This is the value used in files produced with DjVuLibre.
	// This is smaller than DJVUVERSION
	// because version number inflation causes problems with older software.
	DJVUVERSION_FOR_OUTPUT FileVersion = 24

	// This is the version which introduced orientations.
	DJVUVERSION_ORIENTATION FileVersion = 22

	// Oldest DjVu format version supported by this library.
	// This release of the library cannot completely decode DjVu files
	// whose version field is less than or equal to this number.
	DJVUVERSION_TOO_OLD FileVersion = 15

	// Newest DjVu format partially supported by this library.
	// This release of the library will attempt to decode files
	// whose version field is smaller than this macro.
	// If the version field is greater than or equal to this number,
	// the decoder will return an error.
	DJVUVERSION_TOO_NEW FileVersion = 50
)

// Orientation describes how a File is rotated.
// The constant identifiers are expressed in degrees counterclockwise.
type Orientation uint8

const (
	ORIENTATION_0   Orientation = 0b001
	ORIENTATION_90  Orientation = 0b110
	ORIENTATION_180 Orientation = 0b010
	ORIENTATION_270 Orientation = 0b101
)

func (o Orientation) String() string {
	switch o {
	case ORIENTATION_0:
		return "0"
	case ORIENTATION_90:
		return "90"
	case ORIENTATION_180:
		return "180"
	case ORIENTATION_270:
		return "270"
	default:
		return "unknown"
	}
}

// b0<<8 + b1. For decoding
func bytesToUint16(b0, b1 byte) uint16 {
	return uint16(b0)<<8 + uint16(b1)
}
