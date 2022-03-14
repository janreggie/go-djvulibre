package iw44

// Bitmap represents a IW44-encoded grey-level image.
// Contrast this with Bitmap which represents a color image.
// This class provides functions for managing a grey level image
// represented as a collection of IW44 wavelet coefficients.
// The coefficients are stored in a memory efficient data structure.
//
// Method GetBitmap renders an arbitrary segment of the image into a Bitmap.
// Methods DecodeIff and EncodeIff read and write DjVu IW44 file.
//
// Images should not be copied once created.
type Bitmap struct {
	Image
}
