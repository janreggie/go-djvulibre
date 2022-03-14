package iw44

// Pixmap represents a IW44-encoded color image.
// Contrast this with Bitmap which represents a grey-level image.
// This class provides functions for managing a color image
// represented as a collection of IW44 wavelet coefficients.
// The coefficients are stored in a memory efficient data structure.
//
// Method GetPixmap renders an arbitrary segment of the image into a Pixmap.
// Methods DecodeIff and EncodeIff read and write DjVu IW44 file.
//
// Images should not be copied once created.
type Pixmap struct {
	Image
	// Do things...
}
