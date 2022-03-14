package iw44

import (
	"github.com/janreggie/go-djvulibre/djvu/bytestream"
	"github.com/janreggie/go-djvulibre/djvu/image"
)

// Image represents Image-encoded gray-level and color images.
// This struct acts as a base for images
// represented as a collection of Image wavelet coefficients.
// The coefficients are stored in a memory efficient data structure.
//
// Method GetBitmap renders an arbitrary segment of the image into a Bitmap.
// Methods DecodeIFF and EncodeIFF read and write DjVu Image files.
//
// Images should not be copied once created.
type Image interface {
	//////////////
	//  Access  //
	//////////////

	// Returns the width of the image
	GetWidth() uint16

	// Returns the height of the image
	GetHeight() uint16

	// Reconstructs the complete image as a pixmap.
	GetBitmap() *image.Bitmap

	// Reconstructs a segment of the image at a given scale.
	// The subsampling ratio `subsample` must be a power of two
	// between `1` and `32`.
	// Argument `rect` specifies
	// which segment of the subsampled image should be reconstructed.
	// The reconstructed image is returned as a Bitmap
	// whose size is equal to the size of the rectangle.
	GetBitmapSample(subsample uint8, rect *image.Rect) *image.Bitmap

	// Reconstructs the complete image as a pixmap.
	GetPixmap() *image.Pixmap

	// Reconstructs a segment of the image at a given scale.
	// The subsampling ratio `subsample` must be a power of two
	// between `1` and `32`.
	// Argument `rect` specifies
	// which segment of the subsampled image should be reconstructed.
	// The reconstructed image is returned as a Pixmap
	// whose size is equal to the size of the rectangle.
	GetPixmapSample(subsample uint8, rect *image.Rect) *image.Pixmap

	// Returns the amount of memory used by the wavelet coefficients.
	// This amount of memory is expressed in bytes.
	GetMemoryUsage() uintptr

	// Returns the filling ratio of the internal data structure.
	// Wavelet coefficients are stored in a sparse array.
	// This function tells what percentage of bins
	// have been effectively allocated.
	GetPercentMemory() uint64

	/////////////////////////////
	//  Encoding and Decoding  //
	/////////////////////////////

	// Encodes one data chunk into a ByteStream.
	// Parameter controls how much data is generated.
	// The chunk data is written with no IFF header.
	// Successive calls to EncodeChunk encode successive chunks.
	//
	// You must call CloseCodec after encoding the last chunk of a file.
	EncodeChunk(bs bytestream.ByteStream) int

	// Writes a gray level image into DjVu IW44 file.
	// This function creates a composite chunk
	// (identifier `FORM:BM44` or `FORM:PM44`)
	// composed of `chunks` chunks
	// (identifier `BM44` or `PM44`).
	// Data for each chunk is generated with EncodeChunk
	// using the corresponding parameters.
	EncodeIFF(iff *bytestream.IFF, chunks uint16, params *EncoderParams)

	// Decodes one data chunk from a ByteStream.
	// Successive calls to DecodeChunk decode successive chunks.
	//
	// You must call CloseCodec after decoding the last chunk of a file.
	DecodeChunk(bs bytestream.ByteStream) int

	// This function enters a composite chunk
	// (identifier `FORM:BM44`, or `FORM:PM44`),
	// and decodes a maximum of `maxChunks` data chunks
	// (identifier `BM44#`).
	// Data for each chunk is processed using the function DecodeChunk.
	DecodeIff(iff *bytestream.IFF, maxChunks uint16)

	/////////////////////
	//  Miscellaneous  //
	/////////////////////

	// Resets the encoder/decoder state.
	// The first call to DecodeChunk or EncodeChunk
	// initializes the coder for encoding/decoding.
	// This method must be called after processing the last chunk
	// in order to reset the coder and release the associated memory.
	CloseCodec()

	// Returns the chunk serial number.
	// This function returns the serial number of the last chunk
	// encoded with EncodeChunk or decoded with DecodeChunk.
	// The first chunk always has serial number `1`.
	// Successive chunks have increasing serial numbers.
	// Value `0` is returned
	// if no chunks have been encoded/decoded since.
	GetSerial() uint16

	// Set the chrominance delay parameter.
	// This function can be called before encoding the first color IW44 data chunk.
	// Parameter `delay` is an encoding delay
	// which reduces the bitrate associated with the chrominance information.
	// The default chrominance encoding delay is 10.
	//
	// This returns a value.
	// TODO: Look at implementations and determine what it returns.
	SetCrcbDelay(delay uint16) uint16

	// Sets the `dbfrac` parameter.
	// This function can be called before encoding the first IW44 data chunk.
	// Parameter `frac` modifies the decibel estimation algorithm
	// in such a way that the decibel target
	// only pertains to the average error of the fraction `frac`
	// of the most misrepresented 32x32 pixel blocks.
	// Setting arguments `frac` to `1.0` restores the normal behavior.
	SetDbFrac(frac float64)
}

func NewDecode(imageType ImageType) Image {
	panic("unimplemented")
}

func NewEncode(ImageType ImageType) Image {
	panic("unimplemented")
}

func NewEncodeWithBitmap(bitmap *image.Bitmap, mask *image.Bitmap) Image {
	panic("unimplemented")
}

func NewEncodeWithPixmap(pixmap *image.Pixmap, mask *Bitmap, crcbMode CRCBMode) Image {
	panic("unimplemented")
}

// TODO: How can we turn these subclasses into idiomatic Go?
type Transform struct{}
type Block struct{}
type Alloc struct{}
type PrimaryHeader struct{}
type SecondaryHeader struct{}
type TertiaryHeader struct{}

// Determines the type of image
type ImageType bool

const (
	IMAGE_GREY  ImageType = false
	IMAGE_COLOR ImageType = true
)
